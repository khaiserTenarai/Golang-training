package controller

type SalaryController interface {
	UpdateSalary(employeeID int, newSalary float64) error
}
