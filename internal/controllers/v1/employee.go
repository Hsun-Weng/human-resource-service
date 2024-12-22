package v1

import (
	"github.com/Hsun-Weng/human-resource-service/internal/response"
	"github.com/Hsun-Weng/human-resource-service/internal/services"
	"github.com/Hsun-Weng/human-resource-service/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmployeeController interface {
	GetContact(c *gin.Context)
	GetContacts(c *gin.Context)
}

type employeeController struct {
	service services.EmployeeService
}

func NewEmployeeController(service services.EmployeeService) EmployeeController {
	return &employeeController{service: service}
}

// GetContact godoc
// @Summary Get Employee Contact Information
// @Description Retrieve contact details of an employee by their Token.
// @Tags Contacts
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} response.EmployeeContact
// @Router /user/v1/contact [get]
func (controller *employeeController) GetContact(c *gin.Context) {
	employee, err := controller.service.GetEmployee(c.GetUint("employee_id"))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "can't find contact"})
		return
	}
	employeeContact := response.EmployeeContact{
		Name:          employee.Name,
		Email:         employee.CompanyEmail,
		Phone:         employee.Phone,
		LivingAddress: employee.LivingAddress,
		Role:          employee.Role,
		Salary:        employee.Salary,
	}

	c.JSON(http.StatusOK, employeeContact)
}

// GetContacts godoc
// @Summary Get a list of employee contacts by admin
// @Description Get a list of employees with their basic contact information like name, email, phone, etc.
// @Tags Contacts
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param page query int false "Page number" default(1)
// @Param size query int false "Number of contacts per page" default(10)
// @Success 200 {object} response.EmployeeContactWithPagination
// @Router /admin/v1/contacts [get]
func (controller *employeeController) GetContacts(c *gin.Context) {
	page := util.QueryParamInt(c, "page", 1)
	size := util.QueryParamInt(c, "size", 10)
	totalCount, err := controller.service.GetEmployeeTotalCount()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var employeeContacts []*response.EmployeeContact
	if totalCount == 0 {
		c.JSON(http.StatusOK, response.EmployeeContactWithPagination{
			Employees: employeeContacts,
			Pagination: response.Pagination{Total: totalCount,
				Page: page,
				Size: size},
		})
		return
	}
	employees, err := controller.service.GetEmployees(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	for _, employee := range employees {
		employeeContact := response.EmployeeContact{
			Name:          employee.Name,
			Email:         employee.CompanyEmail,
			Phone:         employee.Phone,
			LivingAddress: employee.LivingAddress,
			Role:          employee.Role,
			Salary:        employee.Salary,
		}

		employeeContacts = append(employeeContacts, &employeeContact)
	}

	c.JSON(http.StatusOK, response.EmployeeContactWithPagination{
		Employees: employeeContacts,
		Pagination: response.Pagination{Total: totalCount,
			Page: page,
			Size: size},
	})
}
