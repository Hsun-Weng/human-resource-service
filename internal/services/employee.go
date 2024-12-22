package services

import (
	"fmt"
	"github.com/Hsun-Weng/human-resource-service/internal/models"
	"github.com/Hsun-Weng/human-resource-service/internal/repository"
)

type EmployeeService interface {
	GetEmployee(employeeId uint) (*models.Employee, error)
	GetEmployees(page int, size int) ([]*models.Employee, error)
	GetEmployeeTotalCount() (int64, error)
}

type employeeService struct {
	employeeRepository repository.EmployeeRepository
}

func NewEmployeeService(employeeRepository repository.EmployeeRepository) EmployeeService {
	return &employeeService{employeeRepository}
}

func (service *employeeService) GetEmployee(employeeId uint) (*models.Employee, error) {
	employee, err := service.employeeRepository.FindById(employeeId)
	if err != nil {
		return nil, err
	}
	if employee == nil {
		return nil, fmt.Errorf("employee with id %d not found", employeeId)
	}
	return employee, nil
}

func (service *employeeService) GetEmployees(page int, size int) ([]*models.Employee, error) {
	employees, err := service.employeeRepository.FindAll(page, size)
	if err != nil {
		return nil, err
	}
	if employees == nil {
		return nil, fmt.Errorf("employees not found")
	}
	return employees, nil
}

func (service *employeeService) GetEmployeeTotalCount() (int64, error) {
	count, err := service.employeeRepository.CountAll()
	if err != nil {
		return 0, err
	}
	return count, nil
}
