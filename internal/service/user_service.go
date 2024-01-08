// internal/service/user_service.go

package service

import (
	"fmt"
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
		fmt.Printf("Error in FindByID: %v\n", err)
		return nil, err
	}
	// 这里可以添加更多的业务逻辑，如加载用户的组织信息等

	return user, nil
}

func (s *UserService) UpdateUserName(userID string, newName string) error {
	// Fetch the user from the repository
	user, err := s.UserRepo.FindByID(userID)
	if err != nil {
		fmt.Printf("Error in FindByID: %v\n", err)
		return err
	}

	// Update the user's name
	user.Username = newName

	// Save the updated user back to the repository
	err = s.UserRepo.Update(user)
	if err != nil {
		fmt.Printf("Error in Update: %v\n", err)
		return err
	}

	return nil
}
