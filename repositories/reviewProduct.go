package repositories

import (
	"echo-test-1/models/entities"

	"gorm.io/gorm"
)

type ReviewProductRepository interface {
	CreateReview(reviewProduct entities.ReviewProduct) (*entities.ReviewProduct, error)
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
