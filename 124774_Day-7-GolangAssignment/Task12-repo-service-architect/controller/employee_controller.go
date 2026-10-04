package controller

import (
	"fmt"

	"employee-app/model"
	"employee-app/service"
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

func (c *EmployeeController) Start() {

	var choice int

	for {

		fmt.Println()
		fmt.Println("===== EMPLOYEE =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employees")
		fmt.Println("3. Exit")

		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:

			var name string
			var salary float64

			fmt.Print("Enter name: ")
			fmt.Scan(&name)

			fmt.Print("Enter salary: ")
			fmt.Scan(&salary)

			employee := model.Employee{
				Name:   name,
				Salary: salary,
			}

			err := c.service.AddEmployee(employee)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Employee added successfully.")
			}

		case 2:

			employees := c.service.GetEmployees()

			for _, employee := range employees {

				fmt.Println(
					employee.ID,
					employee.Name,
					employee.Salary,
				)
			}

		case 3:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}
