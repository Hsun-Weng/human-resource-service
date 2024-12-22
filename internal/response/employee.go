package response

type EmployeeContact struct {
	Name          string  `json:"name"`
	Email         string  `json:"company_email"`
	Phone         string  `json:"phone"`
	LivingAddress string  `json:"living_address"`
	Salary        float32 `json:"salary"`
	Role          string  `json:"job_role"`
}

type EmployeeContactWithPagination struct {
	Employees  []*EmployeeContact `json:"employees"`
	Pagination Pagination         `json:"pagination"`
}
