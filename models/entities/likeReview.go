package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type LikeReview struct {
	GormModel
	ReviewProductID uint
	MemberID        uint
	ReviewProduct   *ReviewProduct
	Member          *Member
}

func (lr *LikeReview) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(lr)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
