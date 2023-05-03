package repositories

import (
	"echo-test-1/models/entities"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MemberRepository interface {
	GetAllMembers() ([]entities.Member, error)
	CreateMember(member entities.Member) (*entities.Member, error)
	UpdateMember(id uint, member entities.Member) (*entities.Member, error)
	DeleteMember(id uint) error
}

type MemberRepositoryImpl struct {
	db *gorm.DB
}

func NewMemberRepository(DB *gorm.DB) MemberRepository {
	return &MemberRepositoryImpl{db: DB}
}

func (r *MemberRepositoryImpl) GetAllMembers() ([]entities.Member, error) {
	var members []entities.Member
	err := r.db.Preload(clause.Associations).Find(&members).Error
	return members, err
}

func (r *MemberRepositoryImpl) CreateMember(member entities.Member) (*entities.Member, error) {
	err := r.db.Create(&member).Error
	return &member, err
}

func (r *MemberRepositoryImpl) UpdateMember(id uint, member entities.Member) (*entities.Member, error) {
	err := r.db.Model(&member).Where("id = ?", id).Updates(&member).First(&member).Error
	return &member, err
}

func (r *MemberRepositoryImpl) DeleteMember(id uint) error {
	err := r.db.Delete(&entities.Member{}, id).Error

	return err
}
