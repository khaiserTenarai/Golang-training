package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/employee_management/model"
	"example.com/employee_management/service"
)

type EmployeeController struct {
	service service.EmployeeService
	reader  *bufio.Reader
}

func NewEmployeeController(employeeService service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: employeeService,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (c *EmployeeController) Start() {
	for {
		c.showMenu()

		choice := c.readInt("Enter your choice: ")

		switch choice {
		case 1:
			c.AddEmployee()

		case 2:
			c.GetEmployee()

		case 3:
			c.GetAllEmployees()

		case 4:
			c.UpdateEmployee()

		case 5:
			c.DeleteEmployee()

		case 6:
			fmt.Println("Thank you. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
	}
}

func (c *EmployeeController) showMenu() {
	fmt.Println("========================================")
	fmt.Println("       EMPLOYEE MANAGEMENT SYSTEM")
	fmt.Println("========================================")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Get Employee")
	fmt.Println("3. Get All Employees")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Delete Employee")
	fmt.Println("6. Exit")
	fmt.Println("========================================")
}

func (c *EmployeeController) AddEmployee() {
	fmt.Println("\nADD EMPLOYEE")

	employee := model.Employee{
		ID:     c.readInt("Enter ID: "),
		Name:   c.readString("Enter Name: "),
		Email:  c.readString("Enter Email: "),
		Age:    c.readInt("Enter Age: "),
		Salary: c.readFloat("Enter Salary: "),
		Phone:  c.readString("Enter Phone: "),
	}

	c.service.AddEmployee(employee)
}

func (c *EmployeeController) GetEmployee() {
	fmt.Println("\nGET EMPLOYEE")

	id := c.readInt("Enter Employee ID: ")

	c.service.GetEmployee(id)
}

func (c *EmployeeController) GetAllEmployees() {
	fmt.Println("\nALL EMPLOYEES")

	c.service.GetAllEmployees()
}

func (c *EmployeeController) UpdateEmployee() {
	fmt.Println("\nUPDATE EMPLOYEE")

	employee := model.Employee{
		ID:     c.readInt("Enter ID: "),
		Name:   c.readString("Enter Name: "),
		Email:  c.readString("Enter Email: "),
		Age:    c.readInt("Enter Age: "),
		Salary: c.readFloat("Enter Salary: "),
		Phone:  c.readString("Enter Phone: "),
	}

	c.service.UpdateEmployee(employee)
}

func (c *EmployeeController) DeleteEmployee() {
	fmt.Println("\nDELETE EMPLOYEE")

	id := c.readInt("Enter Employee ID: ")

	c.service.DeleteEmployee(id)
}

func (c *EmployeeController) readString(message string) string {
	fmt.Print(message)

	value, _ := c.reader.ReadString('\n')

	return strings.TrimSpace(value)
}

func (c *EmployeeController) readInt(message string) int {
	for {
		value := c.readString(message)

		result, err := strconv.Atoi(value)

		if err == nil {
			return result
		}

		fmt.Println("Please enter a valid integer.")
	}
}

func (c *EmployeeController) readFloat(message string) float64 {
	for {
		value := c.readString(message)

		result, err := strconv.ParseFloat(value, 64)

		if err == nil {
			return result
		}

		fmt.Println("Please enter a valid number.")
	}
}