package services

import (
	"echo-test-1/models/entities"
	"echo-test-1/models/webs"
	"echo-test-1/repositories"
)

type ReviewProductService interface {
	CreateReview(reviewDTO webs.ReviewProductDTO, memberID uint, productID uint) (webs.ReviewProductResp, error)
	ReviewsProduct(productID uint) ([]webs.ReviewProductResp, error)
}

type ReviewProductServiceImpl struct {
	reviewProductRepository repositories.ReviewProductRepository
}

func NewReviewProductService(ReviewProductRepo repositories.ReviewProductRepository) ReviewProductService {
	return &ReviewProductServiceImpl{reviewProductRepository: ReviewProductRepo}
}

func convertBodyReviewResp(reviewProduct entities.ReviewProduct) webs.ReviewProductResp {
	return webs.ReviewProductResp{
		MemberID:   reviewProduct.MemberID,
		ProductID:  reviewProduct.ProductID,
		DescReview: reviewProduct.DescReview,
	}
}

func (s *ReviewProductServiceImpl) CreateReview(reviewDTO webs.ReviewProductDTO, memberID uint, productID uint) (webs.ReviewProductResp, error) {
	reviewProduct := entities.ReviewProduct{
		MemberID:   memberID,
		ProductID:  productID,
		DescReview: reviewDTO.DescReview,
	}

	newReview, err := s.reviewProductRepository.CreateReview(reviewProduct)
	newReviewResp := convertBodyReviewResp(*newReview)

	return newReviewResp, err
}

func (s *ReviewProductServiceImpl) ReviewsProduct(productID uint) ([]webs.ReviewProductResp, error) {
	reviews, err := s.reviewProductRepository.ReviewsProduct(productID)

	reviewsResp := []webs.ReviewProductResp{}

	for i := 0; i < len(reviews); i++ {
		reviewsBody := convertBodyReviewResp(reviews[i])
		reviewsResp = append(reviewsResp, reviewsBody)
	}

	return reviewsResp, err
}
