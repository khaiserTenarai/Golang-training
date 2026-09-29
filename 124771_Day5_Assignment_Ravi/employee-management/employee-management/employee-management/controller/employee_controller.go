package controller

import (
	"employee-management/model"
	"employee-management/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(
	employeeService service.EmployeeService,
) *EmployeeController {

	return &EmployeeController{
		service: employeeService,
	}
}

func (controller *EmployeeController) AddEmployee(
	employee model.Employee,
) error {

	return controller.service.AddEmployee(employee)
}

func (controller *EmployeeController) GetEmployee(
	id int,
) (model.Employee, error) {

	return controller.service.GetEmployee(id)
}

func (controller *EmployeeController) GetAllEmployees() []model.Employee {

	return controller.service.GetAllEmployees()
}

func (controller *EmployeeController) UpdateEmployee(
	employee model.Employee,
) error {

	return controller.service.UpdateEmployee(employee)
}

func (controller *EmployeeController) DeleteEmployee(
	id int,
) error {

	return controller.service.DeleteEmployee(id)
}

func (controller *EmployeeController) SearchEmployee(
	name string,
) []model.Employee {

	return controller.service.SearchEmployee(name)
}