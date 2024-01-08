package service

import (
	"zcloud-bg/internal/model"
	"zcloud-bg/internal/repository"
)

type OrganizationMemberService struct {
	repo *repository.OrganizationMemberRepository
}

func NewOrganizationMemberService(repo *repository.OrganizationMemberRepository) *OrganizationMemberService {
	return &OrganizationMemberService{repo: repo}
}

func (s *OrganizationMemberService) AddMember(member *model.OrganizationMember) error {
	return s.repo.CreateOrganizationMember(member)
}

func (s *OrganizationMemberService) GetMember(organizationID, userID string) (*model.OrganizationMember, error) {
	return s.repo.GetOrganizationMemberByID(organizationID, userID)
}

func (s *OrganizationMemberService) UpdateMember(member *model.OrganizationMember) error {
	return s.repo.UpdateOrganizationMember(member)
}

func (s *OrganizationMemberService) RemoveMember(organizationID, userID string) error {
	return s.repo.DeleteOrganizationMember(organizationID, userID)
}

func (s *OrganizationMemberService) ListMembers(organizationID, userID string) ([]model.OrganizationMember, error) {
	return s.repo.ListOrganizationMembers(organizationID, userID)
}

func (s *OrganizationMemberService) GetUserOrganizations(userID string) ([]model.OrganizationMember, error) {
	return s.repo.ListOrganizationMembers("", userID)
}

func (s *OrganizationMemberService) ListManagers(organizationID, userID string) ([]model.OrganizationMember, error) {
	return s.repo.ListOrganizationManagers(organizationID, userID)
}
