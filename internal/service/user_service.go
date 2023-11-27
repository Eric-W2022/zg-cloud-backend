// internal/service/user_service.go

package service

import (
	"zcloud-bg/internal/model"
	"zcloud-bg/internal/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		UserRepo: userRepo,
	}
}

func (s *UserService) GetUserByID(userID string) (*model.User, error) {
	// 从 UserRepository 获取用户
	user, err := s.UserRepo.FindByID(userID)
	if err != nil {
		return nil, err
	}

	// 这里可以添加更多的业务逻辑，如加载用户的组织信息等

	return user, nil
}
