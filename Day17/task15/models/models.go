package models

type Department struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Employee struct {
	ID           string  `json:"id"`
	Name         string  `json:"name" binding:"required"`
	Role         string  `json:"role" binding:"required"`
	DepartmentID string  `json:"departmentId" binding:"required"`
	Salary       float64 `json:"salary" binding:"required,gt=0"`
}

type EmployeeFilter struct {
	DepartmentID *string
	MinSalary    *float64
}

type Pagination struct {
	Limit  int
	Offset int
}

type UserContext struct {
	UserID string
	Role   string
}
