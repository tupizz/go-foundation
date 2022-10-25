package database

import (
	"github.com/tupizz/go-foundation/41-api-project/internal/entity"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (p *ProductRepository) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *ProductRepository) FindByID(id string) (*entity.Product, error) {
	product := &entity.Product{}
	err := p.DB.Where("id = ?", id).First(product).Error
	return product, err
}

func (p *ProductRepository) Update(product *entity.Product) error {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return err
	}
	return p.DB.Save(product).Error
}

func (p *ProductRepository) Delete(id string) error {
	productToBeDeleted, err := p.FindByID(id)
	if err != nil {
		return err
	}

	return p.DB.Delete(productToBeDeleted).Error
}

func (p *ProductRepository) FindAll(page, limit int, sort string) ([]*entity.Product, error) {
	var products []*entity.Product
	var err error

	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}

	if page != 0 && limit != 0 {
		err = p.DB.Order("created_at " + sort).Offset((page - 1) * limit).Limit(limit).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}

	return products, err
}
