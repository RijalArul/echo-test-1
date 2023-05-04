package services

import (
	"echo-test-1/models/entities"
	"echo-test-1/models/webs"
	"echo-test-1/repositories"
)

type LikeReviewService interface {
	LikeReview(reviewID uint, memberID uint) (webs.LikeReviewResp, error)
	DislikeReview(reviewID uint) error
}

type LikeReviewServiceImpl struct {
	likeReviewRepository repositories.LikeReviewRepository
}

func NewLikeReviewService(LikeReviewRepository repositories.LikeReviewRepository) LikeReviewService {
	return &LikeReviewServiceImpl{likeReviewRepository: LikeReviewRepository}
}

func convertBodyLikeResponse(likeReview *entities.LikeReview) webs.LikeReviewResp {
	return webs.LikeReviewResp{
		ID:       likeReview.ID,
		ReviewID: likeReview.ReviewProductID,
		MemberID: likeReview.MemberID,
	}
}

func (s *LikeReviewServiceImpl) LikeReview(reviewID uint, memberID uint) (webs.LikeReviewResp, error) {
	likeReview := entities.LikeReview{
		ReviewProductID: reviewID,
		MemberID:        memberID,
	}
	newLikeReview, err := s.likeReviewRepository.CreateLike(likeReview)
	respNewLike := convertBodyLikeResponse(newLikeReview)
	return respNewLike, err

}

func (s *LikeReviewServiceImpl) DislikeReview(reviewID uint) error {
	err := s.likeReviewRepository.DeleteLike(reviewID)
	return err
}
