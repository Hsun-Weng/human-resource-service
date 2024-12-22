package services_test

import (
	"errors"
	"github.com/Hsun-Weng/human-resource-service/internal/models"
	"github.com/Hsun-Weng/human-resource-service/internal/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockEmployeeRepository struct {
	mock.Mock
}

func (m *MockEmployeeRepository) FindById(employeeId uint) (*models.Employee, error) {
	args := m.Called(employeeId)
	return args.Get(0).(*models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) FindAll(page int, size int) ([]*models.Employee, error) {
	args := m.Called(page, size)
	return args.Get(0).([]*models.Employee), args.Error(1)
}

func (m *MockEmployeeRepository) CountAll() (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockEmployeeRepository) FindByCompanyEmail(companyEmail string) (*models.Employee, error) {
	args := m.Called(companyEmail)
	return args.Get(0).(*models.Employee), args.Error(1)
}

func TestGetEmployee(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	service := services.NewEmployeeService(mockRepo)

	expectedEmployee := &models.Employee{
		ID:   1,
		Name: "John Doe",
	}
	mockRepo.On("FindById", uint(1)).Return(expectedEmployee, nil)

	employee, err := service.GetEmployee(1)

	assert.NoError(t, err)
	assert.Equal(t, expectedEmployee, employee)
}

func TestGetEmployee_NotFound(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	service := services.NewEmployeeService(mockRepo)

	mockRepo.On("FindById", uint(1)).Return((*models.Employee)(nil), errors.New("employee not found"))

	employee, err := service.GetEmployee(1)

	assert.Error(t, err)
	assert.Nil(t, employee)

}

func TestGetEmployees(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	service := services.NewEmployeeService(mockRepo)

	employees := []*models.Employee{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Doe"},
	}
	mockRepo.On("FindAll", 1, 10).Return(employees, nil)

	result, err := service.GetEmployees(1, 10)

	assert.NoError(t, err)
	assert.Equal(t, employees, result)

}

func TestGetEmployeeTotalCount(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	service := services.NewEmployeeService(mockRepo)

	mockRepo.On("CountAll").Return(int64(100), nil)

	count, err := service.GetEmployeeTotalCount()

	assert.NoError(t, err)
	assert.Equal(t, int64(100), count)
}

func TestGetEmployeeTotalCount_Error(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	service := services.NewEmployeeService(mockRepo)

	mockRepo.On("CountAll").Return(int64(0), errors.New("database error"))

	count, err := service.GetEmployeeTotalCount()

	assert.Error(t, err)
	assert.Equal(t, int64(0), count)

	mockRepo.AssertExpectations(t)
}
