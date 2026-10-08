package repository

type SalaryRepository interface {
	UpdateSalary(employeeID int, newSalary float64) error
}
