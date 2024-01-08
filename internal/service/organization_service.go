// internal/service/organization_service.go

package service

import (
	"fmt"
	"zcloud-bg/internal/model"
	"zcloud-bg/internal/repository"
)

type OrganizationService struct {
	OrganizationRepo *repository.OrganizationRepository
}

func NewOrganizationService(organizationRepo *repository.OrganizationRepository) *OrganizationService {
	return &OrganizationService{
		OrganizationRepo: organizationRepo,
	}
}

func (s *OrganizationService) GetOrganizationByID(organizationID string) (*model.Organization, error) {
	// 从 OrganizationRepository 获取用户
	organization, err := s.OrganizationRepo.FindByID(organizationID)

	if err != nil {
		fmt.Printf("Error in FindByID: %v\n", err)
		return nil, err
	}
	// 这里可以添加更多的业务逻辑，如加载用户的组织信息等

	return organization, nil
}

func (s *OrganizationService) UpdateOrganizationName(organizationID string, newName string) error {
	// Fetch the organization from the repository
	organization, err := s.OrganizationRepo.FindByID(organizationID)
	if err != nil {
		fmt.Printf("Error in FindByID: %v\n", err)
		return err
	}

	// Update the organization's name
	organization.Name = newName

	// Save the updated organization back to the repository
	err = s.OrganizationRepo.Update(organization)
	if err != nil {
		fmt.Printf("Error in Update: %v\n", err)
		return err
	}

	return nil
}

