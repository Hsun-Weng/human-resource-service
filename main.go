package main

import "github.com/gin-gonic/gin"

func main() {

	app := gin.Default()

	app.GET("/common/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "health check"})
	})
	err := app.Run(":8080")
	if err != nil {
		panic(err)
	}
}
