package model

type Employee struct {
	ID           int
	Name         string
	Email        string
	Salary       float64
	DepartmentID int
}

type EmployeeWithDepartment struct {
	ID             int
	Name           string
	Email          string
	Salary         float64
	DepartmentID   int
	DepartmentName string
}
