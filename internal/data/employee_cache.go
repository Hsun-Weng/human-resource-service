package data

import (
	"time"
)

type EmployeeCache struct {
	ID            uint      `json:"id"`
	CompanyEmail  string    `json:"company_email"`
	Name          string    `json:"name"`
	PersonalEmail string    `json:"personal_email"`
	Phone         string    `json:"phone"`
	LivingAddress string    `json:"living_address"`
	Salary        float32   `json:"salary"`
	Role          string    `json:"role"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
