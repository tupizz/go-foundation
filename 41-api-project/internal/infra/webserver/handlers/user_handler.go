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
	Repository   database.UserDBInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int // in seconds
}

func NewUserHandler(db database.UserDBInterface, jwt *jwtauth.JWTAuth, jwtExpiresIn int) *UserHandler {
	return &UserHandler{
		Repository:   db,
		Jwt:          jwt,
		JwtExpiresIn: jwtExpiresIn,
	}
}

func (h *UserHandler) GetJwt(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	_, token, _ := h.Jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(h.JwtExpiresIn)).Unix(),
	})

	accessToken := struct {
		AccessToken string `json:"access_token"`
	}{
		AccessToken: token,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

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
