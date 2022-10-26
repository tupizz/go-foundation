package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/tupizz/go-foundation/41-api-project/configs"
	"github.com/tupizz/go-foundation/41-api-project/internal/entity"
	"github.com/tupizz/go-foundation/41-api-project/internal/infra/database"
	"github.com/tupizz/go-foundation/41-api-project/internal/infra/webserver/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
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
	userHandler := handlers.NewUserHandler(database.NewUserRepository(db), cfg.TokenAuth, cfg.JWTExpiresIn)
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Heartbeat("/ping"))

	// Routes
	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(cfg.TokenAuth))
		r.Use(jwtauth.Authenticator)
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
