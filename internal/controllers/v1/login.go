package v1

import (
	"github.com/Hsun-Weng/human-resource-service/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController interface {
	Login(c *gin.Context)
}

type loginController struct {
	service services.LoginService
}

func NewLoginController(service services.LoginService) LoginController {
	return &loginController{service: service}
}

type LoginRequest struct {
	CompanyEmail string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
}

// Login godoc
// @Summary User Login
// @Description Login with company email and password to receive an authentication token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginRequest body LoginRequest true "Login request with company email and password"
// @Success 200 {string} json "{"token": "{token}"}"
// @Router /auth/v1/login [post]
func (controller *loginController) Login(c *gin.Context) {
	var loginRequest LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	loginAuthentication := services.Authentication{CompanyEmail: loginRequest.CompanyEmail, Password: loginRequest.Password}
	token, err := controller.service.Login(c.Request.Context(), loginAuthentication)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username or Password is incorrect"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
