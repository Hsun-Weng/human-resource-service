package services

import (
	"context"
	"errors"
	"github.com/Hsun-Weng/human-resource-service/internal/repository"
	"github.com/Hsun-Weng/human-resource-service/pkg/util"
	"golang.org/x/crypto/bcrypt"
)

type Authentication struct {
	CompanyEmail string
	Password     string
}

type LoginService interface {
	Login(ctx context.Context, authentication Authentication) (*string, error)
}

type loginService struct {
	employeeRepository repository.EmployeeRepository
	cacheService       CacheService
}

func NewLoginService(employeeRepository repository.EmployeeRepository, cacheService CacheService) LoginService {
	return &loginService{employeeRepository, cacheService}
}

func (service *loginService) Login(ctx context.Context, authentication Authentication) (*string, error) {
	employee, err := service.employeeRepository.FindByCompanyEmail(authentication.CompanyEmail)
	if err != nil {
		return nil, err
	}
	if employee == nil {
		return nil, errors.New("invalid company email")
	}
	err = bcrypt.CompareHashAndPassword([]byte(employee.Password), []byte(authentication.Password))
	if err != nil {
		return nil, errors.New("invalid password")
	}
	token, err := util.GenerateJWT(employee.ID, employee.Role)
	if err != nil {
		return nil, errors.New("signed error")
	}
	// Cache the login information
	err = service.cacheService.CacheEmployee(ctx, employee)
	if err != nil {
		return nil, err
	}
	return &token, nil
}
