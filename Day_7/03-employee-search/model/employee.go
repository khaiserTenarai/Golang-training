package model
type Employee struct {
	ID         int
	Name       string
	Age        int
	Email      string
	Salary     float64
	Department Department
}

type EmployeeSearch struct {
	Name         string
	DepartmentID int
	MinSalary    float64
	MaxSalary    float64
	Page         int
	Size         int
	SortBy       string
	SortOrder    string
}
