// internal/repository/user_repo.go

package repository

import (
	"gorm.io/gorm"
	"zcloud-bg/internal/model"
)

type OrganizationRepository struct {
	DB *gorm.DB
}

func (repo *OrganizationRepository) FindByID(organizationID string) (*model.Organization, error) { // 确保这个打印语句执行
	var organization model.Organization
	// 使用 GORM 正确执行查询
	result := repo.DB.Where("organization_id = ?", organizationID).First(&organization)
	if result.Error != nil {
		return nil, result.Error
	}

	// 加载相关的组织信息
	// 注意：根据你的实际需求调整加载逻辑
	// err := repo.DB.Model(&organization).Association("Organizations").Find(&organization.Organizations)
	// if err != nil {
	// 	return nil, err
	// }

	// err = repo.DB.Model(&organization).Association("MemberOrganizations").Find(&organization.MemberOrganizations)
	// if err != nil {
	// 	return nil, err
	// }

	return &organization, nil
}

func (repo *OrganizationRepository) Update(organization *model.Organization) error {
	// 使用 GORM 更新记录
	result := repo.DB.Save(organization)
	if result.Error != nil {
		return result.Error
	}

	return nil
}