package v1

import (
	"github.com/Hsun-Weng/human-resource-service/internal/constants/leave_status"
	"github.com/Hsun-Weng/human-resource-service/internal/constants/leave_type"
	"github.com/Hsun-Weng/human-resource-service/internal/requests"
	"github.com/Hsun-Weng/human-resource-service/internal/response"
	"github.com/Hsun-Weng/human-resource-service/internal/services"
	"github.com/Hsun-Weng/human-resource-service/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type LeaveController interface {
	CreateLeave(c *gin.Context)
	GetLeaves(c *gin.Context)
	ReviewLeaves(c *gin.Context)
}

type leaveController struct {
	service services.LeaveService
}

func NewLeaveController(service services.LeaveService) LeaveController {
	return &leaveController{service: service}
}

// CreateLeave godoc
// @Summary Create a Leave Request
// @Description Create a new leave request for the employee by providing leave details.
// @Tags Leave
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param createLeaveRequest body requests.CreateLeaveRequest true "Leave request details"
// @Success 200 {string} string "create leave successfully"
// @Router /user/v1/leave [post]
func (controller *leaveController) CreateLeave(c *gin.Context) {
	employeeId := c.GetUint("employee_id")
	var createLeave requests.CreateLeaveRequest
	if err := c.ShouldBindJSON(&createLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isValid := leave_type.IsValidType(createLeave.LeaveType)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid leave type"})
		return
	}
	startTime := time.Time(createLeave.StartDate)
	endTime := time.Time(createLeave.EndDate)
	if startTime.After(endTime) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Start time cannot be after end time",
		})
		return
	}

	err := controller.service.CreateLeave(employeeId, &createLeave)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "create leave successfully")
}

// GetLeaves godoc
// @Summary Get a list of leaves by admin
// @Description Get a list of leave details
// @Tags Leave
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param page query int false "Page number" default(1)
// @Param size query int false "Number of leaves per page" default(10)
// @Param status query string false "Status for query"
// @Success 200 {array} response.LeaveWithPagination "List of leaves"
// @Router /admin/v1/leaves [get]
func (controller *leaveController) GetLeaves(c *gin.Context) {
	page := util.QueryParamInt(c, "page", 1)
	size := util.QueryParamInt(c, "size", 10)
	status := c.Query("status")
	if status != "" {
		checked := leave_status.IsValidStatus(status)
		if !checked {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		}
		return
	}
	totalCount, err := controller.service.GetLeaveTotalCount(status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if totalCount == 0 {
		c.JSON(http.StatusOK, response.LeaveWithPagination{Leaves: []*response.Leave{},
			Pagination: response.Pagination{Total: totalCount, Page: page, Size: size}})
		return
	}

	leaves, err := controller.service.GetLeaves(status, page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	var leavesResponse []*response.Leave
	for _, leave := range leaves {
		leaveResponse := response.Leave{
			ID:        leave.ID,
			LeaveType: leave.LeaveType,
			StartDate: util.FormatDate(leave.StartTime),
			EndDate:   util.FormatDate(leave.EndTime),
			Status:    leave.Status,
			Reason:    leave.Reason,
			CreatedAt: util.FormatDateTime(leave.CreatedAt),
		}

		leavesResponse = append(leavesResponse, &leaveResponse)
	}

	c.JSON(http.StatusOK, response.LeaveWithPagination{Leaves: leavesResponse,
		Pagination: response.Pagination{Total: totalCount, Page: page, Size: size}})
}

// ReviewLeaves godoc
// @Summary Review Leave
// @Description Review the employee's leave application by admin
// @Tags Leave
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Param reviewLeaveRequest body requests.ReviewLeaveRequest true "Leave request details"
// @Param id path int true "Leave ID"
// @Success 200 {string} review leave successfully "Success Message"
// @Router /admin/v1/leaves/{id} [put]
func (controller *leaveController) ReviewLeaves(c *gin.Context) {
	idStr := c.Param("id")

	leaveId, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID, must be an integer.",
		})
		return
	}

	var reviewLeave requests.ReviewLeaveRequest
	if err := c.ShouldBindJSON(&reviewLeave); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.service.ReviewLeave(uint(leaveId), reviewLeave.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "review leave successfully")
}
