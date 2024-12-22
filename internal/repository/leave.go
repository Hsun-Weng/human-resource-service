package repository

import (
	"errors"
	"github.com/Hsun-Weng/human-resource-service/internal/models"
	"gorm.io/gorm"
	"time"
)

type LeaveRepository interface {
	CreateLeave(Leave *models.Leave) error
	FindByEmployeeIdAndStatusInAndDateBetween(employeeId uint, statuses []string, startDate *time.Time, endDate *time.Time) (int64, error)
	CountAll(status string) (int64, error)
	FindAll(status string, page int, size int) ([]*models.Leave, error)
	FindById(id uint) (*models.Leave, error)
	UpdateStatusById(id uint, status string) error
}

type leaveRepository struct {
	db *gorm.DB
}

func NewLeaveRepository(db *gorm.DB) LeaveRepository {
	return &leaveRepository{db}
}

func (r *leaveRepository) CreateLeave(Leave *models.Leave) error {

	result := r.db.Save(Leave)
	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("failed to save leave")
	}
	return nil
}

func (r *leaveRepository) FindByEmployeeIdAndStatusInAndDateBetween(employeeId uint, statuses []string, startDate *time.Time, endDate *time.Time) (int64, error) {
	var count int64
	result := r.db.Model(&models.Leave{}).Where("employee_id = ? "+
		"AND status in (?) "+
		"AND start_time <= ? "+
		"AND end_time >= ?", employeeId, statuses, endDate, startDate).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (repository *leaveRepository) CountAll(status string) (int64, error) {
	query := repository.db.Model(&models.Leave{})
	if status != "" {
		query = query.Where("status = ?", status)
	}
	var count int64
	result := query.Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

func (repository *leaveRepository) FindAll(status string, page int, size int) ([]*models.Leave, error) {
	var leaves []*models.Leave
	query := repository.db.Model(&leaves)
	if status != "" {
		query.Where("status = ?", status)
	}
	offset := (page - 1) * size
	query = query.Order("created_at desc").Offset(offset).Limit(size).Find(&leaves)
	if err := query.Find(&leaves).Error; err != nil {
		return nil, err
	}
	return leaves, nil
}

func (r *leaveRepository) FindById(id uint) (*models.Leave, error) {
	leave := &models.Leave{}
	result := r.db.Where("id = ?", id).First(leave)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("leave not found")
	}
	return leave, nil
}

func (r *leaveRepository) UpdateStatusById(id uint, status string) error {
	result := r.db.Model(&models.Leave{}).Where("id = ?", id).UpdateColumn("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to update status")
	}
	return nil
}
