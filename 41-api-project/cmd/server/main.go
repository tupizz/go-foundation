package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/tupizz/go-foundation/41-api-project/configs"
	"github.com/tupizz/go-foundation/41-api-project/internal/dto"
	"github.com/tupizz/go-foundation/41-api-project/internal/entity"
	"github.com/tupizz/go-foundation/41-api-project/internal/infra/database"
	"github.com/tupizz/go-foundation/41-api-project/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := configs.LoadConfig(".")

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.User{}, &entity.Product{})

	// Http server
	productHandler := handlers.NewProductHandler(database.NewProductRepository(db))
	userHandler := handlers.NewUserHandler(database.NewUserRepository(db))
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/ping"))
	r.Use(InjectDbMiddleware(db))
	r.Use(middleware.WithValue("JWT", cfg.TokenAuth))                            // inject the jwt token authenticator in the context
	r.Use(middleware.WithValue("JWTExpiresIn", time.Duration(cfg.JWTExpiresIn))) // inject the jwt expires in config value in the context

	// Routes
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Use(InjectUserMiddleware)
		r.Get("/", productHandler.GetProducts)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)
		r.Post("/login", userHandler.GetJwt)
	})

	// Start server
	fmt.Println("Server is running on port", cfg.WebServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.WebServerPort), r)
}

func InjectDbMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, "DB", db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// middleware recebe um handler e retorna um handler basicamente
func InjectUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// getting data from the jwt token
		_, userData, _ := jwtauth.FromContext(r.Context())
		userEmail := userData["email"].(string)

		// getting db from the context
		db := r.Context().Value("DB").(*gorm.DB)

		// querying the user to populate for each request
		user, err := database.NewUserRepository(db).FindByEmail(userEmail)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "User", user)

		log.Println("User injected in the context")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
