package dto

type CreateProductDTO struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetJwtResponse struct {
	AccessToken string `json:"access_token"`
}
