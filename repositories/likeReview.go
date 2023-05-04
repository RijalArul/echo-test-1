package repositories

import (
	"echo-test-1/models/entities"

	"gorm.io/gorm"
)

type LikeReviewRepository interface {
	CreateLike(likeReview entities.LikeReview) (*entities.LikeReview, error)
	DeleteLike(id uint) error
}

type LikeReviewRepositoryImpl struct {
	db *gorm.DB
}

func NewLikeReviewRepository(DB *gorm.DB) LikeReviewRepository {
	return &LikeReviewRepositoryImpl{db: DB}
}

func (r *LikeReviewRepositoryImpl) CreateLike(likeReview entities.LikeReview) (*entities.LikeReview, error) {
	findLikeReview := entities.LikeReview{}
	review := r.db.Where(map[string]interface{}{"review_product_id": likeReview.ReviewProductID, "member_id": likeReview.MemberID}).First(&findLikeReview).Error

	if review != nil {
		err := r.db.Create(&likeReview).Error
		return &likeReview, err
	}
	return &findLikeReview, nil
}

func (r *LikeReviewRepositoryImpl) DeleteLike(id uint) error {
	err := r.db.Delete(&entities.LikeReview{}, id).Error
	return err
}
