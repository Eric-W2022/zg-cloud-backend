// internal/repository/user_repo.go

package repository

import (
	"gorm.io/gorm"
	"zcloud-bg/internal/model"
)

type RoleRepository struct {
	DB *gorm.DB
}
func NewRoleRepository(db *gorm.DB) *RoleRepository {
    return &RoleRepository{DB: db}
}

func (repo *RoleRepository) FindByID(roleID string) (*model.Role, error) { // 确保这个打印语句执行
	var role model.Role
	// 使用 GORM 正确执行查询
	result := repo.DB.Where("role_id = ?", roleID).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}