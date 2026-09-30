package service
import "employee-management/model"

type EmployeeService interface {
	Save(employee model.Employee) error
	FindByID(id int) (model.Employee, error)
	FindAll() ([]model.Employee, error)
	Update(employee model.Employee) error
	Delete(id int) error
	Search(search model.EmployeeSearch) ([]model.Employee, int, error)
}
