package controller

import (
	"fmt"

	"day7q12/model"

	"day7q12/service"
)

type EmployeeController struct {
	service service.EmployeeService
}

func NewEmployeeController(
	service service.EmployeeService,
) *EmployeeController {

	return &EmployeeController{

		service: service,
	}

}

func (c *EmployeeController) CreateEmployee() {

	employee := model.Employee{

		Name: "Rajesh",

		Email: "rajesh@gmail.com",

		Salary: 60000,
	}

	err := c.service.CreateEmployee(employee)

	if err != nil {

		fmt.Println("Create error:", err)

		return

	}

	fmt.Println("Employee created successfully")

}

func (c *EmployeeController) GetEmployees() {

	employees, err := c.service.GetEmployees()

	if err != nil {

		fmt.Println(err)

		return

	}

	fmt.Println("\nEmployees:")

	for _, emp := range employees {

		fmt.Println("----------------")

		fmt.Println("ID:", emp.ID)

		fmt.Println("Name:", emp.Name)

		fmt.Println("Email:", emp.Email)

		fmt.Println("Salary:", emp.Salary)

	}

}
