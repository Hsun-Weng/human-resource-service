package services

import (
	"context"
	"errors"
	"github.com/Hsun-Weng/human-resource-service/internal/data"
	"github.com/Hsun-Weng/human-resource-service/internal/models"
	"github.com/Hsun-Weng/human-resource-service/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEmployeeRepository struct {
	mock.Mock
}
type MockCacheService struct {
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

func (m *MockCacheService) CacheEmployee(ctx context.Context, employee *models.Employee) error {
	args := m.Called(ctx, employee)
	return args.Error(0)
}

func (m *MockCacheService) GetCacheEmployee(ctx context.Context, employeeId uint) (*data.EmployeeCache, error) {
	args := m.Called(ctx, employeeId)
	return args.Get(0).(*data.EmployeeCache), args.Error(1)
}

// TestLogin tests the Login method of the LoginService.
func TestLogin_Success(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	mockCache := new(MockCacheService)
	service := NewLoginService(mockRepo, mockCache)
	ctx := context.Background()

	mockPassword, _ := util.HashPassword("correct-password")
	mockToken, _ := util.GenerateJWT(1, "MANAGER")

	mockEmployee := &models.Employee{
		ID:           1,
		CompanyEmail: "test@company.com",
		Password:     mockPassword,
		Role:         "MANAGER",
	}

	mockRepo.On("FindByCompanyEmail", "test@company.com").Return(mockEmployee, nil)
	mockCache.On("CacheEmployee", ctx, mockEmployee).Return(nil)

	auth := Authentication{
		CompanyEmail: "test@company.com",
		Password:     "correct-password",
	}
	token, err := service.Login(ctx, auth)

	assert.NoError(t, err)
	assert.Equal(t, mockToken, *token)
}

func TestLogin_InvalidPassword(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	mockCache := new(MockCacheService)
	service := NewLoginService(mockRepo, mockCache)
	ctx := context.Background()

	mockPassword, _ := util.HashPassword("correct_password")

	mockEmployee := &models.Employee{
		ID:           1,
		CompanyEmail: "test@company.com",
		Password:     mockPassword,
		Role:         "MANAGER",
	}
	mockRepo.On("FindByCompanyEmail", "test@company.com").Return(mockEmployee, nil)
	mockCache.On("CacheEmployee", ctx, mockEmployee).Return(nil)

	auth := Authentication{
		CompanyEmail: "test@company.com",
		Password:     "wrong-password",
	}

	token, err := service.Login(context.Background(), auth)

	assert.Error(t, err)
	assert.Nil(t, token)
	assert.Equal(t, "invalid password", err.Error())
}

func TestLogin_UserNotFound(t *testing.T) {
	mockRepo := new(MockEmployeeRepository)
	mockCache := new(MockCacheService)
	service := NewLoginService(mockRepo, mockCache)

	mockRepo.On("FindByCompanyEmail", "nonexistent@company.com").Return((*models.Employee)(nil), errors.New("user not found"))

	auth := Authentication{
		CompanyEmail: "nonexistent@company.com",
		Password:     "any-password",
	}

	token, err := service.Login(context.Background(), auth)

	assert.Error(t, err)
	assert.Nil(t, token)
	assert.Equal(t, "user not found", err.Error())
}
