package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Member struct {
	GormModel
	Username       string          `gorm:"not null;unique" valid:"required~Your username is required,minstringlength(6)~Username has to have a minimum length of 6 characters"`
	Gender         string          `gorm:"not null" valid:"required~Your Gender is required"`
	SkinType       string          `gorm:"not null" valid:"required~Your Skin Type is required"`
	SkinColor      string          `gorm:"not null" valid:"required~Your Skin Color is required"`
	ReviewProducts []ReviewProduct `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (m *Member) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(m)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
