package controller

import (
	"bufio"
	"context"
	"fmt"
	"strconv"
	"strings"

	"department-management/service"
	"department-management/view"
)

type EmployeeController struct {
	Service *service.EmployeeService
	Reader  *bufio.Reader
}

func NewEmployeeController(
	service *service.EmployeeService,
	reader *bufio.Reader,
) *EmployeeController {

	return &EmployeeController{
		Service: service,
		Reader:  reader,
	}
}

func (c *EmployeeController) Create() {

	fmt.Println()
	fmt.Println("===== Create Employee =====")

	fmt.Print("Enter employee name: ")
	name := c.readLine()

	fmt.Print("Enter employee email: ")
	email := c.readLine()

	salary := c.readFloat(
		"Enter employee salary: ",
	)

	departmentID := c.readInt(
		"Enter department ID: ",
	)

	employee, err := c.Service.Create(
		context.Background(),
		name,
		email,
		salary,
		departmentID,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(
		"Employee created successfully. ID:",
		employee.ID,
	)
}

func (c *EmployeeController) GetByID() {

	fmt.Println()
	fmt.Println("===== Get Employee =====")

	id := c.readInt(
		"Enter employee ID: ",
	)

	employee, err := c.Service.GetByID(
		context.Background(),
		id,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	view.ShowEmployee(employee)
}

func (c *EmployeeController) GetAll() {

	fmt.Println()
	fmt.Println("===== Get All Employees =====")

	employees, err := c.Service.GetAll(
		context.Background(),
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	view.ShowEmployees(employees)
}

func (c *EmployeeController) Update() {

	fmt.Println()
	fmt.Println("===== Update Employee =====")

	id := c.readInt(
		"Enter employee ID: ",
	)

	fmt.Print("Enter new employee name: ")
	name := c.readLine()

	fmt.Print("Enter new employee email: ")
	email := c.readLine()

	salary := c.readFloat(
		"Enter new salary: ",
	)

	departmentID := c.readInt(
		"Enter new department ID: ",
	)

	err := c.Service.Update(
		context.Background(),
		id,
		name,
		email,
		salary,
		departmentID,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(
		"Employee updated successfully.",
	)
}

func (c *EmployeeController) Delete() {

	fmt.Println()
	fmt.Println("===== Delete Employee =====")

	id := c.readInt(
		"Enter employee ID: ",
	)

	err := c.Service.Delete(
		context.Background(),
		id,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(
		"Employee deleted successfully.",
	)
}

func (c *EmployeeController) ChangeDepartment() {

	fmt.Println()
	fmt.Println("===== Change Employee Department =====")

	employeeID := c.readInt(
		"Enter employee ID: ",
	)

	departmentID := c.readInt(
		"Enter new department ID: ",
	)

	err := c.Service.ChangeDepartment(
		context.Background(),
		employeeID,
		departmentID,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(
		"Employee department changed successfully.",
	)
}

func (c *EmployeeController) readLine() string {

	input, _ := c.Reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func (c *EmployeeController) readInt(
	message string,
) int {

	for {

		fmt.Print(message)

		input := c.readLine()

		value, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println(
				"Please enter a valid number.",
			)
			continue
		}

		return value
	}
}

func (c *EmployeeController) readFloat(
	message string,
) float64 {

	for {

		fmt.Print(message)

		input := c.readLine()

		value, err := strconv.ParseFloat(
			input,
			64,
		)

		if err != nil {
			fmt.Println(
				"Please enter a valid salary.",
			)
			continue
		}

		return value
	}
}
