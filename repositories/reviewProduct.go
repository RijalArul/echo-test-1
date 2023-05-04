package repositories

import (
	"echo-test-1/models/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ReviewProductRepository interface {
	CreateReview(reviewProduct entities.ReviewProduct) (*entities.ReviewProduct, error)
	ReviewsProduct(productID uint) ([]entities.ReviewProduct, error)
}

type ReviewProductRepositoryImpl struct {
	db *gorm.DB
}

func NewReviewProductRepository(DB *gorm.DB) ReviewProductRepository {
	return &ReviewProductRepositoryImpl{db: DB}
}

func (r *ReviewProductRepositoryImpl) CreateReview(reviewProduct entities.ReviewProduct) (*entities.ReviewProduct, error) {
	err := r.db.Create(&reviewProduct).Error

	return &reviewProduct, err
}

func (r *ReviewProductRepositoryImpl) ReviewsProduct(productID uint) ([]entities.ReviewProduct, error) {
	reviews := []entities.ReviewProduct{}
	err := r.db.Preload(clause.Associations).Where("product_id = ?", productID).Find(&reviews).Error
	return reviews, err
}
