package repository
import "ems/model"

/*
	EmployeeRepository defines database
	operations for Employee.
*/
type EmployeeRepository interface {

	// Day 7 CRUD operations.
	Save(employee model.Employee) error

	FindByID(id int) (model.Employee, error)

	FindAll() ([]model.Employee, error)

	Update(employee model.Employee) error

	Delete(id int) error
}