package repositories

import (
	"echo-test-1/models/entities"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts() ([]entities.Product, error)
	CreateProduct(product entities.Product) (*entities.Product, error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{db: DB}
}

func (r *ProductRepositoryImpl) GetAllProducts() ([]entities.Product, error) {
	products := []entities.Product{}
	err := r.db.Model(&products).Find(&products).Error
	return products, err
}

func (r *ProductRepositoryImpl) CreateProduct(product entities.Product) (*entities.Product, error) {
	err := r.db.Create(&product).Error

	return &product, err
}
