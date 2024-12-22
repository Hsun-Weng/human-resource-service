package routers

import (
	v1 "github.com/Hsun-Weng/human-resource-service/internal/controllers/v1"
	middleware2 "github.com/Hsun-Weng/human-resource-service/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Router struct {
	LoginController     *v1.LoginController
	EmployeeController  *v1.EmployeeController
	LeaveController     *v1.LeaveController
	AdminAuthMiddleware *middleware2.AdminAuthenticationMiddleware
}

func NewRouter(adminAuthMiddleware middleware2.AdminAuthenticationMiddleware, LoginController v1.LoginController,
	EmployeeController v1.EmployeeController, LeaveController v1.LeaveController) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/common/health", func(c *gin.Context) {
		c.JSON(200, gin.H{})
	})
	r.POST("/auth/v1/login", LoginController.Login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userApiV1 := r.Group("/user/v1")
	userApiV1.Use(middleware2.JWT())
	userApiV1.GET("/contact", EmployeeController.GetContact)
	userApiV1.POST("/leave", LeaveController.CreateLeave)

	adminApiV1 := r.Group("/admin/v1")
	adminApiV1.Use(middleware2.JWT())
	adminApiV1.Use(adminAuthMiddleware.AdminAuthHandler())
	adminApiV1.GET("/contacts", EmployeeController.GetContacts)
	adminApiV1.GET("/leaves", LeaveController.GetLeaves)
	adminApiV1.PUT("/leaves/:id", LeaveController.ReviewLeaves)

	return r
}
