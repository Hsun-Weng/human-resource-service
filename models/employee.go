package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID          uint      `gorm:"primary_key; column:id"`
	Name        string    `gorm:"column:name"`
	Password    string    `gorm:"column:password"`
	Status      string    `gorm:"column:status"`
	CreatedTime time.Time `gorm:"column:created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time"`
}
