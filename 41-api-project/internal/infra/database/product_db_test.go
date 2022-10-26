package database

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tupizz/go-foundation/41-api-project/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

func TestProductRepository_Create(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Playstation 5", 7640.0)
	assert.NoError(t, err)
	productRepository := NewProductRepository(db)
	err = productRepository.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestProductRepository_FindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	// create some products
	productRepository := NewProductRepository(db)
	for i := 1; i <= 25; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*1000)
		assert.NoError(t, err)
		err = productRepository.Create(product)
		assert.NoError(t, err)
		assert.NotEmpty(t, product.ID)
	}

	// find all products page 1 with 10 items
	products, err := productRepository.FindAll(1, 10, "desc")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 25", products[0].Name)
	assert.Equal(t, "Product 16", products[9].Name)

	// find all products page 2 with 10 items
	products, err = productRepository.FindAll(2, 10, "desc")
	assert.NoError(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 15", products[0].Name)
	assert.Equal(t, "Product 6", products[9].Name)

	// find all products page 3 with 10 items (last page)
	products, err = productRepository.FindAll(3, 10, "desc")
	assert.NoError(t, err)
	assert.Equal(t, 5, len(products))
	assert.Equal(t, "Product 5", products[0].Name)
	assert.Equal(t, "Product 1", products[4].Name)
}

func TestProductRepository_FindByID(t *testing.T) {
	// test configuration part
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	// CreateUser a product to be tested
	product, _ := entity.NewProduct("Playstation 5", 7640.0)
	productRepository := NewProductRepository(db)
	productRepository.Create(product)

	// assert that the product was created
	foundProduct, err := productRepository.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, foundProduct.ID)
	assert.Equal(t, product.Name, foundProduct.Name)
	assert.Equal(t, product.Price, foundProduct.Price)
	assert.Equal(t, product.CreatedAt.Unix(), foundProduct.CreatedAt.Unix())
}

func TestProductRepository_Delete(t *testing.T) {
	// test configuration part
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	// CreateUser a product to be tested
	product, _ := entity.NewProduct("Playstation 5", 7640.0)
	productRepository := NewProductRepository(db)
	productRepository.Create(product)

	// assert that the product was created
	foundProduct, err := productRepository.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, foundProduct.ID)

	// delete the product
	err = productRepository.Delete(product.ID.String())
	assert.NoError(t, err)

	// assert that the product was deleted
	_, err = productRepository.FindByID(product.ID.String())
	assert.Error(t, err)
}

func TestProductRepository_Update(t *testing.T) {
	// test configuration part
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	// CreateUser a product to be tested
	product, _ := entity.NewProduct("Playstation 5", 7640.0)
	productRepository := NewProductRepository(db)
	productRepository.Create(product)

	// assert that the product was created
	foundProduct, err := productRepository.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, foundProduct.ID)

	// update the product
	product.Name = "Playstation 5 Digital Edition"
	product.Price = 4999.99
	err = productRepository.Update(product)
	assert.NoError(t, err)

	// assert that the product was updated
	updatedProduct, err := productRepository.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, updatedProduct.ID)
	assert.Equal(t, product.Name, updatedProduct.Name)
	assert.Equal(t, product.Price, updatedProduct.Price)
}
