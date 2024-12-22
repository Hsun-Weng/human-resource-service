package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	ID            uint      `gorm:"primaryKey;autoIncrement"`
	CompanyEmail  string    `gorm:"type:varchar(200);unique;"`
	Name          string    `gorm:"type:varchar(100);"`
	PersonalEmail string    `gorm:"type:varchar(200);"`
	Password      string    `gorm:"type:varchar(100);"`
	Phone         string    `gorm:"type:varchar(100);"`
	LivingAddress string    `gorm:"type:varchar(500);"`
	Salary        float32   `gorm:"type:decimal(10,2)"`
	Role          string    `gorm:"type:varchar(10);"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	DeletedAt     *time.Time
}
