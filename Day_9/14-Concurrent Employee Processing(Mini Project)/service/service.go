package service
import "ems/model"

/*
	EmployeeService defines employee operations.
*/
type EmployeeService interface {

	/*
		Existing Day 7 CRUD operations.
	*/

	Save(employee model.Employee) error

	FindByID(id int) (model.Employee, error)

	FindAll() ([]model.Employee, error)

	Update(employee model.Employee) error

	Delete(id int) error

	/*
		New Day 9 operation.

		ProcessEmployees processes employees
		concurrently using workers.
	*/
	ProcessEmployees()
}