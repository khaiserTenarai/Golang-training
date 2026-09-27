package controller

import (
	"employee_management/models"
	"employee_management/service"
	"fmt"
)

type EmployeeController struct{
	service service.EmployeeService
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController{
	return &EmployeeController{service: service}
}

func (c *EmployeeController) AddEmployee(employee models.Employee) error {
	fmt.Println("Hello from Contorller: Add Employee")

	return c.service.AddEmployee(employee)
	// fmt.Println("Employee Added successfully")

}

func (c *EmployeeController) GetEmployee(id int) (models.Employee, error) {
	fmt.Println("Hello from controller: Get employee")

	return c.service.GetEmployee(id)
}

func (c *EmployeeController) GetAllEmployee() ([]models.Employee, error){
	fmt.Println("Hello from contorller: Get all employee")

	return c.service.GetAllEmployee()
}

func (c *EmployeeController) UpdateEmployee(employee models.Employee) error{
	fmt.Println("Hello from contorller: Update Employee")

	return c.service.UpdateEmployee(employee)
}

func (c *EmployeeController) DeleteEmployee(id int) error{
	fmt.Println("Hello from contorller: Delete employee")

	return c.service.DeleteEmployee(id)
}

