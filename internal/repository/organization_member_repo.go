// internal/repository/organization_member_repo.go

package repository

import (
	"gorm.io/gorm"
	"zcloud-bg/internal/model"
)

type OrganizationMemberRepository struct {
	DB *gorm.DB
}

// CreateOrganizationMember creates a new organization member record.
func (repo *OrganizationMemberRepository) CreateOrganizationMember(member *model.OrganizationMember) error {
	result := repo.DB.Create(member)
	return result.Error
}

// GetOrganizationMemberByID retrieves an organization member by their organization ID and user ID.
func (repo *OrganizationMemberRepository) GetOrganizationMemberByID(organizationID, userID string) (*model.OrganizationMember, error) {
	var member model.OrganizationMember
	result := repo.DB.Where("organization_id = ? AND user_id = ?", organizationID, userID).First(&member)
	if result.Error != nil {
		return nil, result.Error
	}
	return &member, nil
}

// UpdateOrganizationMember updates a given organization member.
func (repo *OrganizationMemberRepository) UpdateOrganizationMember(member *model.OrganizationMember) error {
	result := repo.DB.Model(&model.OrganizationMember{}).Where("organization_id = ? AND user_id = ?", member.OrganizationID, member.UserID).Updates(member)
	return result.Error
}

// DeleteOrganizationMember deletes an organization member by their organization ID and user ID.
func (repo *OrganizationMemberRepository) DeleteOrganizationMember(organizationID, userID string) error {
	result := repo.DB.Where("organization_id = ? AND user_id = ?", organizationID, userID).Delete(&model.OrganizationMember{})
	return result.Error
}

// ListOrganizationMembers lists all organization members, with optional filters for organization and user IDs.
func (repo *OrganizationMemberRepository) ListOrganizationMembers(organizationID, userID string) ([]model.OrganizationMember, error) {
	var members []model.OrganizationMember
	query := repo.DB.Order("joined_at DESC") // 默认按加入时间降序排序

	// 根据提供的筛选条件构建查询
	if organizationID != "" {
		query = query.Where("organization_id = ?", organizationID)
	}
	// if userID != "" {
	// 	query = query.Where("user_id = ?", userID)
	// }

	// 执行查询
	result := query.Find(&members)
	if result.Error != nil {
		return nil, result.Error
	}

	return members, nil
}

func (repo *OrganizationMemberRepository) ListOrganizationManagers(organizationID, userID string) ([]model.OrganizationMember, error) {
	var managers []model.OrganizationMember
	result := repo.DB.
        Joins("INNER JOIN roles ON organization_members.role_id = roles.role_id").
        Where("organization_members.organization_id = ?", organizationID).
        Where("roles.role_name = ?", "管理员").
        Order("organization_members.joined_at DESC").
        Find(&managers)

    if result.Error != nil {
        return nil, result.Error
    }

    return managers, nil

}
