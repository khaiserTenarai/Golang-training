package controller
import (
	"emp_simple/service"
	"emp_simple/model"
)

type EmployeeController struct{
	Service service.Emp_service
}

func (c *EmployeeController) AddEmp(employee *model.Employee) error{
	return c.Service.AddEmp(employee)
} 

func (c *EmployeeController) DisplayEmployee() []model.Employee{
	return c.Service.GetEmployee()
}

func (c *EmployeeController) DeleteEmployee(id int) bool{
	return c.Service.DeleteEmployee(id)
}

