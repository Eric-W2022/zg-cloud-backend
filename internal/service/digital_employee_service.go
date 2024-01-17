package service

import (
	"zcloud-bg/internal/model"
	"zcloud-bg/internal/repository"
)

type EmployeeService struct {
	EmployeeRepo *repository.EmployeeRepository
}

func NewEmployeeService(employeeRepo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		EmployeeRepo: employeeRepo,
	}
}

// func (s *EmployeeService) ListEmployees(organizationID string) (*model.Digital_Employee, error) {
// 	// 从 EmployeeRepository 获取用户
// 	employee, err := s.EmployeeRepo.FindByID(organizationID)

// 	if err != nil {
// 		fmt.Printf("Error in FindByID: %v\n", err)
// 		return nil, err
// 	}

// 	return employee, nil
// }

func (s *EmployeeService) ListEmployees(organizationID, userID string) ([]model.Digital_Employee, error) {
	return s.EmployeeRepo.ListDigitalEmployees(organizationID, userID)
}