package main

import (
	"encoding/json"
	"fmt"
	"github.com/tupizz/go-foundation/41-api-project/configs"
	"github.com/tupizz/go-foundation/41-api-project/internal/dto"
	"github.com/tupizz/go-foundation/41-api-project/internal/entity"
	"github.com/tupizz/go-foundation/41-api-project/internal/infra/database"
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

	productHandler := NewProductHandler(database.NewProductRepository(db))
	http.HandleFunc("/products", productHandler.CreateProduct)

	fmt.Println("Server is running on port", cfg.WebServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.WebServerPort), nil)
}

type ProductHandler struct {
	Repository database.ProductDBInterface
}

func NewProductHandler(db database.ProductDBInterface) *ProductHandler {
	return &ProductHandler{
		Repository: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var productDto dto.CreateProductDTO
	err := json.NewDecoder(r.Body).Decode(&productDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	p, err := entity.NewProduct(productDto.Name, productDto.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}
	err = h.Repository.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
	return
}
