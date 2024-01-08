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

func (repo *UserRepository) FindByID(userID string) (*model.User, error) { // 确保这个打印语句执行
	var user model.User
	// 使用 GORM 正确执行查询
	result := repo.DB.Where("user_id = ?", userID).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	// 加载相关的组织信息
	// 注意：根据你的实际需求调整加载逻辑
	err := repo.DB.Model(&user).Association("Organizations").Find(&user.Organizations)
	if err != nil {
		return nil, err
	}

	err = repo.DB.Model(&user).Association("MemberOrganizations").Find(&user.MemberOrganizations)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) Update(user *model.User) error {
	// 使用 GORM 更新记录
	result := repo.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}