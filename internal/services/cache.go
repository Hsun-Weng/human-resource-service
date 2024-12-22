package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Hsun-Weng/human-resource-service/internal/constants/redis_keys"
	"github.com/Hsun-Weng/human-resource-service/internal/data"
	"github.com/Hsun-Weng/human-resource-service/internal/models"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type CacheService interface {
	CacheEmployee(ctx context.Context, employee *models.Employee) error
	GetCacheEmployee(ctx context.Context, employeeId uint) (*data.EmployeeCache, error)
}

type cacheService struct {
	client *redis.Client
}

func NewCacheService(client *redis.Client) CacheService {
	return &cacheService{client: client}
}

func (service *cacheService) CacheEmployee(ctx context.Context, employee *models.Employee) error {
	key := redis_keys.GetLoginEmployeeKey(employee.ID)
	cacheEmployee := data.EmployeeCache{
		ID:            employee.ID,
		Name:          employee.Name,
		Role:          employee.Role,
		CompanyEmail:  employee.CompanyEmail,
		LivingAddress: employee.LivingAddress,
		PersonalEmail: employee.PersonalEmail,
		Phone:         employee.Phone,
		Salary:        employee.Salary,
		CreatedAt:     employee.CreatedAt,
		UpdatedAt:     employee.UpdatedAt,
	}
	cacheString, err := json.Marshal(cacheEmployee)
	if err != nil {
		return err
	}
	// cache for ten hours
	err = service.client.Set(ctx, key, cacheString, 60*10*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (service *cacheService) GetCacheEmployee(ctx context.Context, employeeId uint) (*data.EmployeeCache, error) {
	key := redis_keys.GetLoginEmployeeKey(employeeId)
	val, err := service.client.Get(ctx, key).Result()
	if errors.Is(redis.Nil, err) {
		fmt.Println("key does not exist")
	} else if err != nil {
		log.Fatalf("could not get key: %v", err)
	} else {
		fmt.Printf("key value: %s\n", val)
	}
	var cacheEmployee data.EmployeeCache
	err = json.Unmarshal([]byte(val), &cacheEmployee)
	if err != nil {
		return nil, err
	}
	return &cacheEmployee, nil
}
