package database

import (
	"github.com/tupizz/go-foundation/41-api-project/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (u *UserRepository) CreateUser(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}
	err := u.DB.Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
