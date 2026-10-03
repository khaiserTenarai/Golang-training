package view
import "ems/model"

/*
	EmployeeView defines operations
	of the view layer.
*/
type EmployeeView interface {

	ShowMenu() int

	ReadEmployee() model.Employee

	ReadEmployeeForUpdate() model.Employee

	ReadID() int

	DisplayEmployee(employee model.Employee)

	DisplayEmployees(employees []model.Employee)

	/*
		Day 9:

		Display the result of
		concurrent processing.
	*/
	DisplayProcessingMessage()
}