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

func (repo *EmployeeRepository) ListDigitalEmployees(organizationID, userID string) ([]model.Digital_Employee, error) {
	var employees []model.Digital_Employee
	query := repo.DB.Order("created_at DESC") // 默认按加入时间降序排序

	// 根据提供的筛选条件构建查询
	if organizationID != "" {
		query = query.Where("organization_id = ?", organizationID)
	}
	// if userID != "" {
	// 	query = query.Where("user_id = ?", userID)
	// }

	// 执行查询
	result := query.Find(&employees)
	if result.Error != nil {
		return nil, result.Error
	}

	return employees, nil
}
