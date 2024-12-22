package requests

type CreateLeaveRequest struct {
	StartDate CustomDate `json:"start_time"`
	EndDate   CustomDate `json:"end_time"`
	LeaveType string     `json:"leave_type"`
	Reason    string     `json:"reason"`
}

type ReviewLeaveRequest struct {
	Status string `json:"status"`
}
