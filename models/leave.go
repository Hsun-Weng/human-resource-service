package models

import (
	"time"

	"gorm.io/gorm"
)

type Leave struct {
	gorm.Model
	ID          uint      `gorm:"primary_key; column:id"`
	EmployeeId  uint      `gorm:"column:employee_id"`
	StartTime   time.Time `gorm:"column:start_time"`
	EndTime     time.Time `gorm:"column:end_time"`
	Status      int       `gorm:"column:status"`
	CreatedTime time.Time `gorm:"column:created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time"`
}
