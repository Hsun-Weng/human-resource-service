package requests

type CreateLeaveRequest struct {
	StartDate CustomDate `json:"start_date"`
	EndDate   CustomDate `json:"end_date"`
	LeaveType string     `json:"leave_type"`
	Reason    string     `json:"reason"`
}

type ReviewLeaveRequest struct {
	Status string `json:"status"`
}
