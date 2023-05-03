package entities

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	GormModel
	NameProduct string `gorm:"not null;unique" valid:"required~Your Name Product is required,minstringlength(6)~Name Product has to have a minimum length of 6 characters"`
	Price       int    `gorm:"not null" valid:"required~Your Price is required"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
