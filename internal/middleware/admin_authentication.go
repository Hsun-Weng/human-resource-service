package middleware

import (
	"context"
	"errors"
	"github.com/Hsun-Weng/human-resource-service/internal/constants/job_role"
	"github.com/Hsun-Weng/human-resource-service/internal/data"
	"github.com/Hsun-Weng/human-resource-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminAuthenticationMiddleware interface {
	AdminAuthHandler() gin.HandlerFunc
}

type adminAuthenticationMiddleware struct {
	cacheService services.CacheService
}

func NewAdminAuthenticationMiddleware(cacheService services.CacheService) AdminAuthenticationMiddleware {
	return &adminAuthenticationMiddleware{cacheService: cacheService}
}

func (adminAuth *adminAuthenticationMiddleware) AdminAuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		employeeId := c.GetUint("employee_id")
		cacheEmployee, err := adminAuth.getCacheEmployee(c.Request.Context(), uint(employeeId))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		if cacheEmployee.Role != string(job_role.Manager) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func (adminAuth *adminAuthenticationMiddleware) getCacheEmployee(ctx context.Context, employeeId uint) (*data.EmployeeCache, error) {
	cacheEmployee, err := adminAuth.cacheService.GetCacheEmployee(ctx, employeeId)
	if err != nil {
		return nil, errors.New("can not get employee from cache")
	}
	return cacheEmployee, nil
}
