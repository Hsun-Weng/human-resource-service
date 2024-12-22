package middleware

import (
	"github.com/Hsun-Weng/human-resource-service/pkg/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
		if bearerToken == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		claims, err := util.ParseJWT(bearerToken)
		if err != nil {
			log.Println("Invalid token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		c.Set("employee_id", claims.EmployeeId)
		c.Next()
	}
}
