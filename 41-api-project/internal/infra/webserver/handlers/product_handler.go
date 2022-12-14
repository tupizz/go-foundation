package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/tupizz/go-foundation/41-api-project/internal/dto"
	"github.com/tupizz/go-foundation/41-api-project/internal/entity"
	"github.com/tupizz/go-foundation/41-api-project/internal/infra/database"
	entityPackage "github.com/tupizz/go-foundation/41-api-project/pkg/entity"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	Repository database.ProductDBInterface
}

func NewProductHandler(db database.ProductDBInterface) *ProductHandler {
	return &ProductHandler{
		Repository: db,
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product
// @Tags products
// @Accept  json
// @Produce  json
// @Param request body dto.CreateProductDTO true "Product to be created"
// @Success 201 {object} entity.Product
// @Failure 400 {object} dto.ErrorMessage
// @Failure 500 {object} dto.ErrorMessage
// @Router /products [post]
// @Security ApiKeyAuth
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

// GetProduct godoc
// @Summary Get a product
// @Description Get a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Success 200 {object} entity.Product
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Router /products/{id} [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: "you should provide the product id"})
		return
	}

	product, err := h.Repository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID"
// @Param request body entity.Product true "Product to be updated"
// @Success 200 {object} entity.Product
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Router /products/{id} [put]
// @Security ApiKeyAuth
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: "you should provide the product id"})
		return
	}

	_, err := h.Repository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	var productToBeUpdated entity.Product
	err = json.NewDecoder(r.Body).Decode(&productToBeUpdated)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	productToBeUpdated.ID, err = entityPackage.ParseID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	err = h.Repository.Update(&productToBeUpdated)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	// return the product that was just updated
	productJustUpdated, err := h.Repository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productJustUpdated)
	return
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product
// @Tags products
// @Accept  json
// @Produce  json
// @Param id path string true "Product ID" Format(uuid)
// @Success 204
// @Failure 400 {object} dto.ErrorMessage
// @Failure 404 {object} dto.ErrorMessage
// @Router /products/{id} [delete]
// @Security ApiKeyAuth
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: "you should provide the product id"})
		return
	}

	_, err := h.Repository.FindByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	err = h.Repository.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
	return
}

// GetProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Accept  json
// @Produce  json
// @Param page query string false "Page number"
// @Param limit query string false "limit"
// @Success 200 {array} entity.Product
// @Failure 400 {object} dto.ErrorMessage
// @Router /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		pageInt = 0
	}

	products, err := h.Repository.FindAll(pageInt, limitInt, sort)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
