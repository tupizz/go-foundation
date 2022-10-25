package dto

type CreateProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}
