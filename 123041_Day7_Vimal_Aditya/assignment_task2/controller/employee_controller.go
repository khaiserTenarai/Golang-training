package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/employee-management/model"
	"example.com/employee-management/repository"
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
			fmt.Println("\nReturning to Main Menu...")
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
	fmt.Println("6. Back to Main Menu")
	fmt.Println("==============================================")
}

func (c *EmployeeController) addEmployee() {
	fmt.Println("\n---------- ADD EMPLOYEE ----------")
	employee := model.Employee{}
	employee.Name = c.readString("Enter Name: ")
	employee.Email = c.readString("Enter Email: ")
	employee.Age = c.readInt("Enter Age: ")
	employee.Salary = c.readFloat("Enter Salary: ")
	employee.Address.City = c.readString("Enter City: ")
	employee.Address.State = c.readString("Enter State: ")
	employee.Address.Pincode = c.readString("Enter Pincode: ")

	deptID := c.readInt("Enter Department ID (0 if none): ")
	if deptID > 0 {
		employee.DepartmentID = &deptID
	}

	err := c.service.AddEmployee(employee)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\nEmployee added successfully.")
}

func (c *EmployeeController) getEmployee() {
	fmt.Println("\n---------- FIND EMPLOYEE ----------")
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
	fmt.Println("\n---------- ALL EMPLOYEES ----------")
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
	fmt.Printf("%-5s %-18s %-22s %-5s %-10s %-16s %-12s\n", "ID", "NAME", "EMAIL", "AGE", "SALARY", "DEPARTMENT", "CITY")
	fmt.Println(strings.Repeat("-", 95))
	for _, emp := range employees {
		fmt.Printf("%-5d %-18s %-22s %-5d %-10.2f %-16s %-12s\n", emp.ID, emp.Name, emp.Email, emp.Age, emp.Salary, emp.DepartmentName, emp.Address.City)
	}
}

func (c *EmployeeController) updateEmployee() {
	fmt.Println("\n---------- UPDATE EMPLOYEE ----------")
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

	fmt.Println("\nCurrent employee detail:")
	employee.Display()

	fmt.Println("\nEnter new values.")
	employee.Name = c.readString("Enter Name: ")
	employee.Email = c.readString("Enter Email: ")
	employee.Age = c.readInt("Enter Age: ")
	employee.Salary = c.readFloat("Enter Salary: ")
	employee.Address.City = c.readString("Enter City: ")
	employee.Address.State = c.readString("Enter State: ")
	employee.Address.Pincode = c.readString("Enter Pincode: ")

	deptID := c.readInt("Enter Department ID (0 if none): ")
	if deptID > 0 {
		employee.DepartmentID = &deptID
	} else {
		employee.DepartmentID = nil
	}

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
	fmt.Println("\n---------- DELETE EMPLOYEE ----------")
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

	confirmation := c.readString("Are you sure you want to delete? (y/n): ")
	switch strings.ToLower(confirmation) {
	case "y", "yes":
		err = c.service.DeleteEmployee(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Employee deleted successfully.")
	default:
		fmt.Println("Delete operation cancelled.")
	}
}

func (c *EmployeeController) readString(message string) string {
	for {
		fmt.Print(message)
		val, err := c.reader.ReadString('\n')
		if err != nil {
			continue
		}
		val = strings.TrimSpace(val)
		if val != "" {
			return val
		}
		fmt.Println("Value cannot be empty.")
	}
}

func (c *EmployeeController) readInt(message string) int {
	for {
		val := c.readString(message)
		num, err := strconv.Atoi(val)
		if err == nil {
			return num
		}
		fmt.Println("Please enter a valid integer.")
	}
}

func (c *EmployeeController) readInt64(message string) int64 {
	for {
		val := c.readString(message)
		num, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return num
		}
		fmt.Println("Please enter a valid number.")
	}
}

func (c *EmployeeController) readFloat(message string) float64 {
	for {
		val := c.readString(message)
		num, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return num
		}
		fmt.Println("Please enter a valid decimal number.")
	}
}

func (c *EmployeeController) pause() {
	fmt.Println()
	fmt.Print("Press ENTER to continue...")
	_, _ = c.reader.ReadString('\n')
}