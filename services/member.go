package services

import (
	"echo-test-1/models/entities"
	"echo-test-1/models/webs"
	"echo-test-1/repositories"
)

type MemberService interface {
	GetAllMembers() ([]webs.MemberResponse, error)
	CreateMember(memberDTO webs.MemberDTO) (webs.MemberResponse, error)
	UpdateMember(id uint, memberDTO webs.MemberDTO) (webs.MemberResponse, error)
	DeleteMember(id uint) error
}

type MemberServiceImpl struct {
	memberRepository repositories.MemberRepository
}

func NewMemberService(MemberRepo repositories.MemberRepository) MemberService {
	return &MemberServiceImpl{memberRepository: MemberRepo}
}

func convertBodyMemberResp(member entities.Member) webs.MemberResponse {
	return webs.MemberResponse{
		ID:        member.ID,
		Username:  member.Username,
		Gender:    member.Gender,
		SkinType:  member.SkinType,
		SkinColor: member.SkinColor,
	}
}

func (s *MemberServiceImpl) GetAllMembers() ([]webs.MemberResponse, error) {
	members, err := s.memberRepository.GetAllMembers()
	memberResp := []webs.MemberResponse{}

	for i := 0; i < len(members); i++ {
		memberBody := convertBodyMemberResp(members[i])
		memberResp = append(memberResp, memberBody)
	}
	return memberResp, err
}

func (s *MemberServiceImpl) CreateMember(memberDTO webs.MemberDTO) (webs.MemberResponse, error) {
	member := entities.Member{
		Username:  memberDTO.Username,
		Gender:    memberDTO.Gender,
		SkinType:  memberDTO.SkinType,
		SkinColor: memberDTO.SkinColor,
	}
	newMember, err := s.memberRepository.CreateMember(member)
	respNewMember := convertBodyMemberResp(*newMember)
	return respNewMember, err
}

func (s *MemberServiceImpl) UpdateMember(id uint, memberDTO webs.MemberDTO) (webs.MemberResponse, error) {
	member := entities.Member{
		Username:  memberDTO.Username,
		Gender:    memberDTO.Gender,
		SkinType:  memberDTO.SkinType,
		SkinColor: memberDTO.SkinColor,
	}

	updateMember, err := s.memberRepository.UpdateMember(id, member)
	respUpdateMember := convertBodyMemberResp(*updateMember)
	return respUpdateMember, err
}

func (s *MemberServiceImpl) DeleteMember(id uint) error {
	err := s.memberRepository.DeleteMember(id)
	return err
}
