package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/employee-management/model"
	"example.com/employee-management/service"
	"example.com/employee-management/util"
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
			c.GetEmployee()
		case 3:
			c.GetAllEmployees()
		case 4:
			c.UpdateEmployee()
		case 5:
			c.DeleteEmployee()
		case 6:
			c.JSONDemo()
		case 7:
			fmt.Println("Thank you. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice")
		}

		fmt.Println()
	}
}

func (c *EmployeeController) showMenu() {
	util.Divider()
	fmt.Println("       EMPLOYEE MANAGEMENT SYSTEM")
	util.Divider()
	fmt.Println("1. Add Employee")
	fmt.Println("2. Get Employee")
	fmt.Println("3. Get All Employees")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Delete Employee")
	fmt.Println("6. JSON Demo (Marshal/Unmarshal an Employee)")
	fmt.Println("7. Exit")
	util.Divider()
}

func (c *EmployeeController) readEmployeeDetails() model.Employee {
	e := model.Employee{}
	e.ID = c.readInt("Enter ID (0 to auto-generate): ")
	e.Person = model.Person{
		Name:  c.readString("Enter Name: "),
		Email: c.readString("Enter Email: "),
	}
	e.Salary = c.readFloat("Enter Salary: ")
	e.Age = c.readInt("Enter Age: ")
	e.Address = model.Address{
		City:    c.readString("Enter City: "),
		Pincode: c.readString("Enter Pincode: "),
	}
	e.Department = model.Department{
		Name: c.readString("Enter Department Name: "),
	}
	e.IsActive = true
	return e
}

func (c *EmployeeController) AddEmployee() {
	fmt.Println("\nHello from Controller - Add Employee")

	employee := c.readEmployeeDetails()

	saved, err := c.service.AddEmployee(employee)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Added:", saved)
}

func (c *EmployeeController) GetEmployee() {
	fmt.Println("\nHello from Controller - Get Employee")
	id := c.readInt("Enter Employee ID: ")

	employee, err := c.service.GetEmployee(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Found:", employee)
}

func (c *EmployeeController) GetAllEmployees() {
	fmt.Println("\nHello from Controller - Get All Employees")

	employees := c.service.GetAllEmployees()
	if len(employees) == 0 {
		fmt.Println("No employees yet.")
		return
	}
	for _, e := range employees {
		fmt.Println(e)
	}
}

func (c *EmployeeController) UpdateEmployee() {
	fmt.Println("\nHello from Controller - Update Employee")

	employee := c.readEmployeeDetails()

	updated, err := c.service.UpdateEmployee(employee)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Updated:", updated)
}

func (c *EmployeeController) DeleteEmployee() {
	fmt.Println("\nHello from Controller - Delete Employee")
	id := c.readInt("Enter Employee ID: ")

	if err := c.service.DeleteEmployee(id); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Deleted employee", id)
}

// JSONDemo shows Marshal (struct -> JSON) then Unmarshal (JSON -> struct)
// using the util package's generic helpers.
func (c *EmployeeController) JSONDemo() {
	fmt.Println("\nHello from Controller - JSON Demo")

	sample := model.Employee{
		Person:     model.Person{Name: "Demo User", Email: "demo@example.com", Age: 30},
		ID:         999,
		Salary:     50000,
		Address:    model.Address{City: "Bengaluru", Pincode: "560001"},
		Department: model.Department{Name: "Engineering"},
		IsActive:   true,
	}

	jsonStr, err := util.ToJSON(sample)
	if err != nil {
		fmt.Println("Marshal error:", err)
		return
	}
	fmt.Println("Marshaled JSON:")
	fmt.Println(jsonStr)

	var roundTrip model.Employee
	if err := util.FromJSON(jsonStr, &roundTrip); err != nil {
		fmt.Println("Unmarshal error:", err)
		return
	}
	fmt.Println("Unmarshaled back into Employee struct:")
	fmt.Println(roundTrip)
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
