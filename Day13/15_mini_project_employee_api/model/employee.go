package model

type Employee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

type SalaryReport struct {
	Employees     int     `json:"employees"`
	TotalSalary   float64 `json:"total_salary"`
	AverageSalary float64 `json:"average_salary"`
}
