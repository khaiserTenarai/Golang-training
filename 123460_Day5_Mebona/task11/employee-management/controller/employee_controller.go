package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"employee-management/model"
	"employee-management/service"
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
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nEmployee Management")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employee")
		fmt.Println("3. View All Employees")
		fmt.Println("4. Update Employee")
		fmt.Println("5. Delete Employee")
		fmt.Println("6. Exit")

		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid choice.")
			continue
		}

		switch choice {
		case 1:
			c.addEmployee(reader)

		case 2:
			c.getEmployee(reader)

		case 3:
			c.getAllEmployees()

		case 4:
			c.updateEmployee(reader)

		case 5:
			c.deleteEmployee(reader)

		case 6:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (c *EmployeeController) addEmployee(
	reader *bufio.Reader,
) {
	var employee model.Employee

	fmt.Print("Enter ID: ")
	fmt.Fscan(reader, &employee.ID)
	reader.ReadString('\n')

	fmt.Print("Enter Name: ")
	employee.Name, _ = reader.ReadString('\n')
	employee.Name = strings.TrimSpace(employee.Name)

	fmt.Print("Enter Email: ")
	employee.Email, _ = reader.ReadString('\n')
	employee.Email = strings.TrimSpace(employee.Email)

	fmt.Print("Enter Age: ")
	fmt.Fscan(reader, &employee.Age)

	fmt.Print("Enter Salary: ")
	fmt.Fscan(reader, &employee.Salary)
	reader.ReadString('\n')

	if c.service.AddEmployee(employee) {
		fmt.Println("Employee added successfully.")
	} else {
		fmt.Println("Employee ID already exists.")
	}
}

func (c *EmployeeController) getEmployee(
	reader *bufio.Reader,
) {
	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Fscan(reader, &id)
	reader.ReadString('\n')

	employee, found := c.service.GetEmployee(id)

	if !found {
		fmt.Println("Employee not found.")
		return
	}

	employee.Display()
}

func (c *EmployeeController) getAllEmployees() {
	employees := c.service.GetAllEmployees()

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for _, employee := range employees {
		fmt.Println("--------------------")
		employee.Display()
	}
}

func (c *EmployeeController) updateEmployee(
	reader *bufio.Reader,
) {
	var employee model.Employee

	fmt.Print("Enter Employee ID: ")
	fmt.Fscan(reader, &employee.ID)
	reader.ReadString('\n')

	fmt.Print("Enter New Name: ")
	employee.Name, _ = reader.ReadString('\n')
	employee.Name = strings.TrimSpace(employee.Name)

	fmt.Print("Enter New Email: ")
	employee.Email, _ = reader.ReadString('\n')
	employee.Email = strings.TrimSpace(employee.Email)

	fmt.Print("Enter New Age: ")
	fmt.Fscan(reader, &employee.Age)

	fmt.Print("Enter New Salary: ")
	fmt.Fscan(reader, &employee.Salary)
	reader.ReadString('\n')

	if c.service.UpdateEmployee(employee) {
		fmt.Println("Employee updated successfully.")
	} else {
		fmt.Println("Employee not found.")
	}
}

func (c *EmployeeController) deleteEmployee(
	reader *bufio.Reader,
) {
	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Fscan(reader, &id)
	reader.ReadString('\n')

	if c.service.DeleteEmployee(id) {
		fmt.Println("Employee deleted successfully.")
	} else {
		fmt.Println("Employee not found.")
	}
}