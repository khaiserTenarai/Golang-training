package controller


import (
	"fmt"

	"employee-management/model"
	"employee-management/service"
)

type EmployeeControllerImpl struct {
	service service.EmployeeService
}

func NewEmployeeController(
	service service.EmployeeService,
) EmployeeController {

	return &EmployeeControllerImpl{
		service: service,
	}
}

func (c *EmployeeControllerImpl) AddEmployee(
	employee model.Employee,
) {

	err := c.service.AddEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully.")
}

func (c *EmployeeControllerImpl) GetAllEmployees() {

	employees := c.service.GetAllEmployees()

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Println()
	fmt.Println("========== Employees ==========")

	for _, employee := range employees {
		c.printEmployee(employee)
	}
}

func (c *EmployeeControllerImpl) GetEmployeeByID(
	id int,
) {

	employee := c.service.GetEmployeeByID(id)

	if employee.ID == 0 {
		fmt.Println("Employee not found.")
		return
	}

	fmt.Println()
	fmt.Println("========== Employee ==========")
	c.printEmployee(employee)
}

func (c *EmployeeControllerImpl) UpdateEmployee(
	employee model.Employee,
) {

	err := c.service.UpdateEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee updated successfully.")
}

func (c *EmployeeControllerImpl) DeleteEmployee(
	id int,
) {

	err := c.service.DeleteEmployee(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee deleted successfully.")
}

func (c *EmployeeControllerImpl) printEmployee(
	employee model.Employee,
) {

	fmt.Println("ID     :", employee.ID)
	fmt.Println("Name   :", employee.Name)
	fmt.Println("Age    :", employee.Age)
	fmt.Println("Salary :", employee.Salary)
	fmt.Println("-------------------------------")
}