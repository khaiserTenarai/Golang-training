package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"
)

type EmployeeController struct {
	service service.EmployeeService
	reader  *bufio.Reader
}

func NewEmployeeController(
	service service.EmployeeService,
) *EmployeeController {

	return &EmployeeController{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

// Start is the console application's main loop.
// Go does not have a while keyword; for {} is used as a while-style loop.
func (c *EmployeeController) Start() {

	for {

		c.showMenu()

		choice := c.readInt("Enter your choice: ")

		switch choice {

		case 1:
			c.addEmployee()

		case 2:
			c.getEmployee()

		case 3:
			c.getAllEmployees()

		case 4:
			c.updateEmployee()

		case 5:
			c.deleteEmployee()

		case 6:
			fmt.Println()
			fmt.Println("Thank you for using Employee Management System.")
			return

		default:
			fmt.Println("Invalid choice. Please select 1 to 6.")
		}

		c.pause()
	}
}

func (c *EmployeeController) showMenu() {

	fmt.Println()
	fmt.Println("==============================================")
	fmt.Println("       EMPLOYEE MANAGEMENT SYSTEM")
	fmt.Println("==============================================")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Find Employee")
	fmt.Println("3. Find All Employees")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Delete Employee")
	fmt.Println("6. Exit")
	fmt.Println("==============================================")
}

func (c *EmployeeController) addEmployee() {

	fmt.Println()
	fmt.Println("---------- ADD EMPLOYEE ----------")

	employee := model.Employee{}

	employee.Name = c.readString("Enter Name: ")
	employee.Email = c.readString("Enter Email: ")
	employee.Age = c.readInt("Enter Age: ")
	employee.Salary = c.readFloat("Enter Salary: ")

	employee.Address.City =
		c.readString("Enter City: ")

	employee.Address.State =
		c.readString("Enter State: ")

	employee.Address.Pincode =
		c.readString("Enter Pincode: ")

	err := c.service.AddEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("Employee added successfully.")
}

func (c *EmployeeController) getEmployee() {

	fmt.Println()
	fmt.Println("---------- FIND EMPLOYEE ----------")

	id := c.readInt64("Enter Employee ID: ")

	employee, err := c.service.GetEmployee(id)

	if errors.Is(err, repository.ErrNotFound) {
		fmt.Println("Employee not found.")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	employee.Display()
}

func (c *EmployeeController) getAllEmployees() {

	fmt.Println()
	fmt.Println("---------- ALL EMPLOYEES ----------")

	employees, err := c.service.GetAllEmployees()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Println()
	fmt.Printf(
		"%-5s %-20s %-28s %-5s %-12s %-18s\n",
		"ID",
		"NAME",
		"EMAIL",
		"AGE",
		"SALARY",
		"CITY",
	)

	fmt.Println(
		strings.Repeat("-", 95),
	)

	for _, employee := range employees {

		fmt.Printf(
			"%-5d %-20s %-28s %-5d %-12.2f %-18s\n",
			employee.ID,
			employee.Name,
			employee.Email,
			employee.Age,
			employee.Salary,
			employee.Address.City,
		)
	}
}

func (c *EmployeeController) updateEmployee() {

	fmt.Println()
	fmt.Println("---------- UPDATE EMPLOYEE ----------")

	id := c.readInt64("Enter Employee ID: ")

	employee, err := c.service.GetEmployee(id)

	if errors.Is(err, repository.ErrNotFound) {
		fmt.Println("Employee not found.")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("Current employee:")
	employee.Display()

	fmt.Println()
	fmt.Println("Enter new values.")

	employee.Name = c.readString("Enter Name: ")
	employee.Email = c.readString("Enter Email: ")
	employee.Age = c.readInt("Enter Age: ")
	employee.Salary = c.readFloat("Enter Salary: ")

	employee.Address.City =
		c.readString("Enter City: ")

	employee.Address.State =
		c.readString("Enter State: ")

	employee.Address.Pincode =
		c.readString("Enter Pincode: ")

	err = c.service.UpdateEmployee(employee)

	if errors.Is(err, repository.ErrNotFound) {
		fmt.Println("Employee not found.")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee updated successfully.")
}

func (c *EmployeeController) deleteEmployee() {

	fmt.Println()
	fmt.Println("---------- DELETE EMPLOYEE ----------")

	id := c.readInt64("Enter Employee ID: ")

	employee, err := c.service.GetEmployee(id)

	if errors.Is(err, repository.ErrNotFound) {
		fmt.Println("Employee not found.")
		return
	}

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	employee.Display()

	confirmation := c.readString(
		"Are you sure you want to delete? (y/n): ",
	)

	switch strings.ToLower(confirmation) {

	case "y", "yes":

		err = c.service.DeleteEmployee(id)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Employee deleted successfully.")

	case "n", "no":
		fmt.Println("Delete operation cancelled.")

	default:
		fmt.Println("Invalid response. Delete operation cancelled.")
	}
}

func (c *EmployeeController) readString(
	message string,
) string {

	for {

		fmt.Print(message)

		value, err := c.reader.ReadString('\n')

		if err != nil {
			continue
		}

		value = strings.TrimSpace(value)

		if value != "" {
			return value
		}

		fmt.Println("Value cannot be empty.")
	}
}

func (c *EmployeeController) readInt(
	message string,
) int {

	for {

		value := c.readString(message)

		number, err := strconv.Atoi(value)

		if err == nil {
			return number
		}

		fmt.Println("Please enter a valid integer.")
	}
}

func (c *EmployeeController) readInt64(
	message string,
) int64 {

	for {

		value := c.readString(message)

		number, err := strconv.ParseInt(value, 10, 64)

		if err == nil {
			return number
		}

		fmt.Println("Please enter a valid number.")
	}
}

func (c *EmployeeController) readFloat(
	message string,
) float64 {

	for {

		value := c.readString(message)

		number, err := strconv.ParseFloat(value, 64)

		if err == nil {
			return number
		}

		fmt.Println("Please enter a valid decimal number.")
	}
}

func (c *EmployeeController) pause() {

	fmt.Println()
	fmt.Print("Press ENTER to continue...")

	_, _ = c.reader.ReadString('\n')
}
