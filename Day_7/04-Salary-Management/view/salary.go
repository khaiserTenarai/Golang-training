package view
import "salary-management/model"

type SalaryView interface {

	ReadSalaryUpdate() (
		int,
		float64,
	)

	ReadEmployeeID() int

	DisplayHistory(
		history []model.SalaryHistory,
	)
}
