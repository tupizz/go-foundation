package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
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
	db.AutoMigrate(&entity.Product{}, &entity.Product{})

	// Http server
	productHandler := handlers.NewProductHandler(database.NewProductRepository(db))
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/ping"))
	r.Get("/products", productHandler.GetProducts)
	r.Post("/products", productHandler.CreateProduct)
	r.Get("/products/{id}", productHandler.GetProduct)
	r.Put("/products/{id}", productHandler.UpdateProduct)
	r.Delete("/products/{id}", productHandler.DeleteProduct)

	fmt.Println("Server is running on port", cfg.WebServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.WebServerPort), r)
}
