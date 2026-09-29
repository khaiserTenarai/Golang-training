package controller

import (
	"fmt"

	"employee-management/model"
	"employee-management/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(svc service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: svc}
}

func (c *EmployeeController) HandleAddEmployee(emp model.Employee) {
	err := c.service.AddEmployee(emp)
	if err != nil {
		fmt.Println("Add failed:", err)
		return
	}
	fmt.Println("Employee added:", emp.Name)
}

func (c *EmployeeController) HandleGetEmployee(id int) {
	emp, err := c.service.GetEmployee(id)
	if err != nil {
		fmt.Println("Get failed:", err)
		return
	}
	fmt.Println("Employee found:", emp.Describe())
}

func (c *EmployeeController) HandleDeleteEmployee(id int) {
	err := c.service.DeleteEmployee(id)
	if err != nil {
		fmt.Println("Delete failed:", err)
		return
	}
	fmt.Println("Employee deleted, ID:", id)
}
