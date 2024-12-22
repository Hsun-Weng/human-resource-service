package leave_status

type LeaveStatus string

const (
	Pending  = "PENDING"
	Approved = "APPROVED"
	Declined = "DECLINED"
)

var LeaveStatuses = []string{Pending, Approved, Declined}

func IsValidStatus(input string) bool {
	for _, leaveStatus := range LeaveStatuses {
		if input == leaveStatus {
			return true
		}
	}
	return false
}
