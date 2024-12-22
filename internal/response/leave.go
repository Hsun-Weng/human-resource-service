package response

type Leave struct {
	ID        uint   `json:"id"`
	LeaveType string `json:"leave_type"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Status    string `json:"status"`
	Reason    string `json:"reason"`
	CreatedAt string `json:"created_at"`
}

type LeaveWithPagination struct {
	Leaves     []*Leave   `json:"leaves"`
	Pagination Pagination `json:"pagination"`
}
