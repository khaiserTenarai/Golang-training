package view
import "employee-management/model"

type EmployeeView interface {
	ShowMenu() int
	ReadEmployee() model.Employee
	ReadID() int
	ReadSearch() model.EmployeeSearch

	DisplayEmployee(employee model.Employee)
	DisplayEmployees(employees []model.Employee)

	DisplaySearchResult(
		employees []model.Employee,
		total int,
		search model.EmployeeSearch,
	)
}
