// internal/repository/user_repo.go

package repository

import (
	"gorm.io/gorm"
	"zcloud-bg/internal/model"
)

type UserRepository struct {
	DB *gorm.DB
}

func (repo *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	result := repo.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
