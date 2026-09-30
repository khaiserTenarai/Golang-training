package repository
import "salary-management/model"

type SalaryRepository interface {

	UpdateSalary(
		employeeID int,
		newSalary float64,
	) error

	FindHistory(
		employeeID int,
	) ([]model.SalaryHistory, error)
}
