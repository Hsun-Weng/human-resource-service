package services

import (
	"github.com/Hsun-Weng/human-resource-service/internal/constants/leave_status"
	"github.com/Hsun-Weng/human-resource-service/internal/models"
	"github.com/Hsun-Weng/human-resource-service/internal/requests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockLeaveRepository struct {
	mock.Mock
}

func (m *MockLeaveRepository) CreateLeave(leave *models.Leave) error {
	args := m.Called(leave)
	return args.Error(0)
}

func (m *MockLeaveRepository) FindByEmployeeIdAndStatusInAndDateBetween(employeeId uint, statuses []string, startDate, endDate *time.Time) (int64, error) {
	args := m.Called(employeeId, statuses, startDate, endDate)
	return int64(args.Int(0)), args.Error(1)
}

func (m *MockLeaveRepository) FindAll(status string, page int, size int) ([]*models.Leave, error) {
	args := m.Called(status, page, size)
	return args.Get(0).([]*models.Leave), args.Error(1)
}

func (m *MockLeaveRepository) CountAll(status string) (int64, error) {
	args := m.Called(status)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockLeaveRepository) FindById(id uint) (*models.Leave, error) {
	args := m.Called(id)
	return args.Get(0).(*models.Leave), args.Error(1)
}

func (m *MockLeaveRepository) UpdateStatusById(id uint, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

func TestCreateLeave(t *testing.T) {
	mockRepo := new(MockLeaveRepository)
	service := NewLeaveService(mockRepo)

	mockRepo.On("FindByEmployeeIdAndStatusInAndDateBetween", uint(1), []string{leave_status.Pending, leave_status.Approved}, mock.Anything, mock.Anything).Return(0, nil)

	createLeaveRequest := &requests.CreateLeaveRequest{
		LeaveType: "Sick",
		StartDate: requests.CustomDate(time.Now()),
		EndDate:   requests.CustomDate(time.Now().Add(24 * time.Hour)),
		Reason:    "Feeling sick",
	}

	mockRepo.On("CreateLeave", mock.Anything).Return(nil)

	err := service.CreateLeave(1, createLeaveRequest)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetLeaves(t *testing.T) {
	mockRepo := new(MockLeaveRepository)
	service := NewLeaveService(mockRepo)

	mockRepo.On("FindAll", "approved", 1, 10).Return([]*models.Leave{
		&models.Leave{EmployeeId: 1, Status: leave_status.Approved, LeaveType: "Sick", Reason: "Feeling better"},
	}, nil)

	leaves, err := service.GetLeaves("approved", 1, 10)

	assert.NoError(t, err)
	assert.Len(t, leaves, 1)
	assert.Equal(t, leave_status.Approved, leaves[0].Status)
}

func TestReviewLeave(t *testing.T) {
	mockRepo := new(MockLeaveRepository)
	service := NewLeaveService(mockRepo)

	leave := &models.Leave{
		ID:         1,
		Status:     leave_status.Pending,
		EmployeeId: 1,
	}

	mockRepo.On("FindById", uint(1)).Return(leave, nil)
	mockRepo.On("UpdateStatusById", uint(1), "approved").Return(nil)

	err := service.ReviewLeave(1, "approved")

	assert.NoError(t, err)
}

func TestReviewLeave_LeaveNotFound(t *testing.T) {
	mockRepo := new(MockLeaveRepository)
	service := NewLeaveService(mockRepo)

	mockRepo.On("FindById", uint(1)).Return((*models.Leave)(nil), nil)

	err := service.ReviewLeave(1, "approved")

	assert.Error(t, err)
	assert.Equal(t, "leave not found", err.Error())
}

func TestReviewLeave_AlreadyReviewed(t *testing.T) {
	mockRepo := new(MockLeaveRepository)
	service := NewLeaveService(mockRepo)

	leave := &models.Leave{
		ID:         1,
		Status:     leave_status.Approved,
		EmployeeId: 1,
	}

	mockRepo.On("FindById", uint(1)).Return(leave, nil)

	err := service.ReviewLeave(1, "approved")

	assert.Error(t, err)
	assert.Equal(t, "leave is not pending", err.Error())
}
