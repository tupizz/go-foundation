package database

import (
	"github.com/stretchr/testify/assert"
	"github.com/tupizz/go-foundation/41-api-project/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		t.Error(err)
	}

	user, _ := entity.NewUser("John Doe", "john.doe@email.com", "123456")
	userDb := NewUser(db)
	err = userDb.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestUser_FindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		return
	}

	user, _ := entity.NewUser("John Doe", "john.doe@email.com", "123456")
	userDb := NewUser(db)
	err = userDb.Create(user)
	assert.Nil(t, err)

	userFound, err := userDb.FindByEmail("john.doe@email.com")
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
