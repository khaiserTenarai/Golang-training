package repository
import "ems/model"

type EmployeeRepository interface {
	Save(employee model.Employee) error
	FindByID(id int) (model.Employee, error)
	FindAll() ([]model.Employee, error)
	Update(employee model.Employee) error
	Delete(id int) error
}
