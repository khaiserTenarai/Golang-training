package view

import (
	"bufio"
	"emp_management_Updated/controller"
	"emp_management_Updated/model"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EmployeeView struct {
	controller controller.EmployeeController
}

// Constructor
func NewEmployeeView(c controller.EmployeeController) *EmployeeView {
	return &EmployeeView{
		controller: c,
	}
}

// Menu
func (v *EmployeeView) ShowMenu() {

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Println()
		fmt.Println("===== Employee Management =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Display Employees")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. Update Employee")
		fmt.Println("5. Exit")

		fmt.Print("Enter choice: ")

		input, _ := reader.ReadString('\n')
		choice, _ := strconv.Atoi(strings.TrimSpace(input))

		switch choice {

		case 1:
			v.addEmployee(reader)

		case 2:
			v.displayEmployees()

		case 3:
			v.deleteEmployee(reader)

		case 4:
			v.updateEmployee(reader)

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

// Add employee
func (v *EmployeeView) addEmployee(reader *bufio.Reader) {

	var employee model.Employee

	// ID
	for {
		fmt.Print("Enter ID: ")

		input, _ := reader.ReadString('\n')

		id, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Invalid ID. Please enter a number.")
			continue
		}

		employee.ID = id
		break
	}

	// Name
	for {
		fmt.Print("Enter Name: ")

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			fmt.Println("Name cannot be empty.")
			continue
		}

		employee.Name = name
		break
	}

	// Age
	for {
		fmt.Print("Enter Age: ")

		input, _ := reader.ReadString('\n')

		age, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Invalid age. Please enter a number.")
			continue
		}

		employee.Age = age
		break
	}

	// Salary
	for {
		fmt.Print("Enter Salary: ")

		input, _ := reader.ReadString('\n')

		salary, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

		if err != nil {
			fmt.Println("Invalid salary. Please enter a number.")
			continue
		}

		employee.Salary = salary
		break
	}

	// Send employee to controller
	err := v.controller.AddEmployee(&employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully!")
}

// Display employees
func (v *EmployeeView) displayEmployees() {

	employees := v.controller.DisplayEmployees()

	if len(employees) == 0 {
		fmt.Println("No Employee Found")
		return
	}

	fmt.Println()
	fmt.Println("===== Employee List =====")

	for _, employee := range employees {

		fmt.Println("ID:", employee.ID)
		fmt.Println("Name:", employee.Name)
		fmt.Println("Age:", employee.Age)
		fmt.Println("Salary:", employee.Salary)
		fmt.Println("-------------------------")
	}
}

// Delete employee
func (v *EmployeeView) deleteEmployee(reader *bufio.Reader) {

	fmt.Print("Enter Employee ID to delete: ")

	input, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(input))

	if v.controller.DeleteEmployee(id) {
		fmt.Println("Employee deleted successfully!")
	} else {
		fmt.Println("Employee not found!")
	}
}

// Update employee
func (v *EmployeeView) updateEmployee(reader *bufio.Reader) {

	var employee model.Employee

	// ID
	for {
		fmt.Print("Enter Employee ID to update: ")

		input, _ := reader.ReadString('\n')

		id, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Invalid ID. Please enter a number.")
			continue
		}

		employee.ID = id
		break
	}

	// Name
	for {
		fmt.Print("Enter New Name: ")

		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		if name == "" {
			fmt.Println("Name cannot be empty.")
			continue
		}

		employee.Name = name
		break
	}

	// Age
	for {
		fmt.Print("Enter New Age: ")

		input, _ := reader.ReadString('\n')

		age, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			fmt.Println("Invalid age. Please enter a number.")
			continue
		}

		employee.Age = age
		break
	}

	// Salary
	for {
		fmt.Print("Enter New Salary: ")

		input, _ := reader.ReadString('\n')

		salary, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

		if err != nil {
			fmt.Println("Invalid salary. Please enter a number.")
			continue
		}

		employee.Salary = salary
		break
	}

	err := v.controller.UpdateEmployee(&employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee updated successfully!")
}
