package handlers

import (
	"encoding/json"
	"github.com/go-chi/jwtauth"
	"github.com/tupizz/go-foundation/41-api-project/internal/dto"
	"github.com/tupizz/go-foundation/41-api-project/internal/entity"
	"github.com/tupizz/go-foundation/41-api-project/internal/infra/database"
	"net/http"
	"time"
)

type UserHandler struct {
	Repository database.UserDBInterface
}

func NewUserHandler(db database.UserDBInterface) *UserHandler {
	return &UserHandler{
		Repository: db,
	}
}

// GetJwt godoc
// @Summary Get JWT
// @Description Get JWT
// @Tags user
// @Accept  json
// @Produce  json
// @Param request body dto.GetJwtDTO true "login request"
// @Success 200 {object} dto.GetJwtResponse
// @Failure 400 {object} dto.ErrorMessage
// @Failure 401 {object} dto.ErrorMessage
// @Failure 500 {object} dto.ErrorMessage
// @Router /users/login [post]
func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	jwtTokenAuth := r.Context().Value("JWT").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JWTExpiresIn").(time.Duration)

	var loginDto dto.GetJwtDTO
	err := json.NewDecoder(r.Body).Decode(&loginDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	user, err := h.Repository.FindByEmail(loginDto.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	if !user.ValidatePassword(loginDto.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: "invalid credentials"})
		return
	}

	_, token, _ := jwtTokenAuth.Encode(map[string]interface{}{
		"sub":   user.ID.String(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Second * jwtExpiresIn).Unix(),
	})

	accessToken := &dto.GetJwtResponse{
		AccessToken: token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// CreateUser godoc
// @Summary Create user
// @Description Create user
// @Tags user
// @Accept  json
// @Produce  json
// @Param request body dto.CreateUserInput true "user request"
// @Success 201 {object} entity.User
// @Failure 400 {object} dto.ErrorMessage
// @Failure 500 {object} dto.ErrorMessage
// @Router /users [post]
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var userDto dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	user, err := entity.NewUser(userDto.Name, userDto.Email, userDto.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	err = h.Repository.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dto.ErrorMessage{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	return
}
