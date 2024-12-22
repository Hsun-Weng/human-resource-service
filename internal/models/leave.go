package models

import (
	"time"

	"gorm.io/gorm"
)

type Leave struct {
	gorm.Model
	ID         uint   `gorm:"primaryKey;autoIncrement;"`
	EmployeeId uint   `gorm:"index:idx_employee_id_created_at"`
	LeaveType  string `gorm:"type:varchar(10);"`
	StartTime  time.Time
	EndTime    time.Time
	Status     string    `gorm:"type:varchar(10);"`
	Reason     string    `gorm:"type:varchar(200);"`
	CreatedAt  time.Time `gorm:"index:idx_employee_id_created_at"`
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}
