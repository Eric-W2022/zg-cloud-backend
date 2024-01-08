
package service

import (
	"fmt"
	"zcloud-bg/internal/model"
	"zcloud-bg/internal/repository"
)

type RoleService struct {
	RoleRepo *repository.RoleRepository
}

func NewRoleService(roleRepo *repository.RoleRepository) *RoleService {
	return &RoleService{
		RoleRepo: roleRepo,
	}
}

func (s *RoleService) GetRoleByID(roleID string) (*model.Role, error) {
	// 从 RoleRepository 获取用户
	role, err := s.RoleRepo.FindByID(roleID)

	if err != nil {
		fmt.Printf("Error in FindByID: %v\n", err)
		return nil, err
	}
	// 这里可以添加更多的业务逻辑，如加载用户的组织信息等

	return role, nil
}
