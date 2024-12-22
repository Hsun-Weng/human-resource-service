package repository

import (
	"errors"
	"github.com/Hsun-Weng/human-resource-service/internal/models"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	FindById(id uint) (*models.Employee, error)
	FindByCompanyEmail(companyEmail string) (*models.Employee, error)
	FindAll(page int, size int) ([]*models.Employee, error)
	CountAll() (int64, error)
}

type employeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepository{db}
}

func (repository *employeeRepository) FindById(id uint) (*models.Employee, error) {
	var employee models.Employee
	result := repository.db.Find(&employee, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("employee not found")
	}
	return &employee, nil
}

func (repository *employeeRepository) FindByCompanyEmail(companyEmail string) (*models.Employee, error) {
	var employee models.Employee
	result := repository.db.First(&employee, "company_email = ?", companyEmail)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("employee not found")
	}
	return &employee, nil
}

func (repository *employeeRepository) FindAll(page int, size int) ([]*models.Employee, error) {
	var employees []*models.Employee
	query := repository.db.Model(&employees)
	offset := (page - 1) * size
	query = query.Order("id desc").Offset(offset).Limit(size)
	if err := query.Find(&employees).Error; err != nil {
		return nil, err
	}
	return employees, nil
}

func (repository *employeeRepository) CountAll() (int64, error) {
	var count int64
	result := repository.db.Model(&models.Employee{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}
