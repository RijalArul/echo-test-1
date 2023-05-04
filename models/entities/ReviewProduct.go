package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type ReviewProduct struct {
	GormModel
	MemberID   uint
	ProductID  uint
	DescReview string       `gorm:"type:text" valid:"required~Your Desc Review is required"`
	LikeReview []LikeReview `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Member     *Member
	Product    *Product
}

func (rp *ReviewProduct) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(rp)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
