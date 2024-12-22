package util

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func QueryParamInt(c *gin.Context, key string, defaultValue int) int {
	param := c.Query(key)
	result, err := strconv.Atoi(param)
	if err != nil {
		return defaultValue
	}
	return result
}
