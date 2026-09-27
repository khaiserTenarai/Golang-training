package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/employee-management/model"
	"example.com/employee-management/service"
)

type EmployeeController struct {
	service service.EmployeeService
	reader  *bufio.Reader
}

func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		service: service,
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
			c.DisplayEmployees()
		case 3:
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
	fmt.Println("2. Display Employees")
	fmt.Println("3. Exit")
	fmt.Println("========================================")
}

func (c *EmployeeController) AddEmployee() {
	fmt.Println("\nAdd Employee")

	employee := model.Employee{
		ID:     c.readInt("Enter ID: "),
		Name:   c.readString("Enter Name: "),
		Email:  c.readString("Enter Email: "),
		Age:    c.readInt("Enter Age: "),
		Salary: c.readFloat("Enter Salary: "),
	}

	c.service.AddEmployee(employee)
	fmt.Println("Employee added successfully.")
}

func (c *EmployeeController) DisplayEmployees() {
	fmt.Println("\nDisplay Employees")

	employees := c.service.GetAllEmployees()
	if len(employees) == 0 {
		fmt.Println("No employees added yet.")
		return
	}

	for _, e := range employees {
		fmt.Printf("ID: %d | Name: %s | Email: %s | Age: %d | Salary: %.2f\n",
			e.ID, e.Name, e.Email, e.Age, e.Salary)
	}
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
