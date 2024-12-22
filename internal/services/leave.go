package services

import (
	"errors"
	"fmt"
	"github.com/Hsun-Weng/human-resource-service/internal/constants/leave_status"
	"github.com/Hsun-Weng/human-resource-service/internal/models"
	"github.com/Hsun-Weng/human-resource-service/internal/repository"
	"github.com/Hsun-Weng/human-resource-service/internal/requests"
	"time"
)

type LeaveService interface {
	CreateLeave(employeeId uint, createLeaveRequest *requests.CreateLeaveRequest) error
	ReviewLeave(id uint, status string) error
	GetLeaves(status string, page int, size int) ([]*models.Leave, error)
	GetLeaveTotalCount(status string) (int64, error)
}

type leaveService struct {
	repository repository.LeaveRepository
}

func NewLeaveService(leaveRepository repository.LeaveRepository) LeaveService {
	return &leaveService{repository: leaveRepository}
}

func (service *leaveService) CreateLeave(employeeId uint, createLeaveRequest *requests.CreateLeaveRequest) error {
	startDate := time.Time(createLeaveRequest.StartDate)
	endDate := time.Time(createLeaveRequest.EndDate)
	err := service.checkDatesOverlap(employeeId, &startDate, &endDate)
	if err != nil {
		return err
	}
	leave := models.Leave{
		EmployeeId: employeeId,
		LeaveType:  createLeaveRequest.LeaveType,
		StartTime:  startDate,
		EndTime:    endDate,
		Status:     string(leave_status.Pending),
		Reason:     createLeaveRequest.Reason,
	}
	err = service.repository.CreateLeave(&leave)
	if err != nil {
		return err
	}
	return nil
}

func (service *leaveService) checkDatesOverlap(employeeId uint, startDate *time.Time, endDate *time.Time) error {
	checkStatuses := []string{leave_status.Pending, leave_status.Approved}
	count, err := service.repository.FindByEmployeeIdAndStatusInAndDateBetween(employeeId, checkStatuses, startDate, endDate)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("leave is overlap")
	}
	return nil
}

func (service *leaveService) GetLeaves(status string, page int, size int) ([]*models.Leave, error) {
	leaves, err := service.repository.FindAll(status, page, size)
	if err != nil {
		return nil, err
	}
	return leaves, nil
}

func (service *leaveService) GetLeaveTotalCount(status string) (int64, error) {
	count, err := service.repository.CountAll(status)
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (service *leaveService) ReviewLeave(id uint, status string) error {
	leave, err := service.repository.FindById(id)
	if err != nil {
		return err
	}
	if leave == nil {
		return fmt.Errorf("leave not found")
	}
	if leave.Status != leave_status.Pending {
		return fmt.Errorf("leave is not pending")
	}
	err = service.repository.UpdateStatusById(id, status)
	if err != nil {
		return err
	}
	return nil
}
