package database

import "github.com/tupizz/go-foundation/41-api-project/internal/entity"

type UserDBInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
