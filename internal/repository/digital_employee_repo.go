// internal/repository/Employee_repo.go

package repository

import (
	"gorm.io/gorm"
	"zcloud-bg/internal/model"
)

type EmployeeRepository struct {
	DB *gorm.DB
}

func (repo *EmployeeRepository) FindByID(organizationID string) (*model.Digital_Employee, error) { // 确保这个打印语句执行
	var employee model.Digital_Employee

	result := repo.DB.Where("organization_id = ?", organizationID).First(&employee)
	if result.Error != nil {
		return nil, result.Error
	}

	return &employee, nil
}
