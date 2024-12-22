package leave_type

type LeaveType string

const (
	Sick   = "SICK"
	Annual = "ANNUAL"
)

var LeaveTypes = []string{Sick, Annual}

func IsValidType(input string) bool {
	for _, leaveType := range LeaveTypes {
		if input == leaveType {
			return true
		}
	}
	return false
}
