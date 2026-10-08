package service

type SalaryService interface {
	UpdateSalary(employeeID int, newSalary float64) error
}
