package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"employee-management-app/model"
)

type EmployeeView struct {
	reader *bufio.Reader
}

// Creates View.
func NewEmployeeView() *EmployeeView {

	return &EmployeeView{
		reader: bufio.NewReader(os.Stdin),
	}
}

// Displays menu.
func (v *EmployeeView) ShowMenu() {

	fmt.Println()
	fmt.Println("========== EMPLOYEE MANAGEMENT ==========")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Delete Employee")
	fmt.Println("3. Update Employee")
	fmt.Println("4. Find Employee By ID")
	fmt.Println("5. Find All Employees")
	fmt.Println("6. Exit")
	fmt.Println("=========================================")
}

// Reads integer.
func (v *EmployeeView) ReadInt(message string) int {

	for {

		fmt.Print(message)

		input, _ := v.reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)

		if err == nil {
			return value
		}

		fmt.Println("Please enter a valid number.")
	}
}

// Reads decimal number.
func (v *EmployeeView) ReadFloat(message string) float64 {

	for {

		fmt.Print(message)

		input, _ := v.reader.ReadString('\n')

		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)

		if err == nil {
			return value
		}

		fmt.Println("Please enter a valid salary.")
	}
}

// Reads text.
func (v *EmployeeView) ReadString(message string) string {

	fmt.Print(message)

	input, _ := v.reader.ReadString('\n')

	return strings.TrimSpace(input)
}

// Reads complete employee.
func (v *EmployeeView) ReadEmployee() model.Employee {

	return model.Employee{
		ID:     v.ReadInt("Enter ID: "),
		Name:   v.ReadString("Enter Name: "),
		Email:  v.ReadString("Enter Email: "),
		Age:    v.ReadInt("Enter Age: "),
		Salary: v.ReadFloat("Enter Salary: "),
	}
}

// Displays one employee.
func (v *EmployeeView) ShowEmployee(
	employee model.Employee,
) {

	fmt.Println("--------------------------------")
	fmt.Println("ID     :", employee.ID)
	fmt.Println("Name   :", employee.Name)
	fmt.Println("Email  :", employee.Email)
	fmt.Println("Age    :", employee.Age)
	fmt.Println("Salary :", employee.Salary)
	fmt.Println("--------------------------------")
}

// Displays all employees.
func (v *EmployeeView) ShowEmployees(
	employees []model.Employee,
) {

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for _, employee := range employees {
		v.ShowEmployee(employee)
	}
}

// Displays a message.
func (v *EmployeeView) ShowMessage(message string) {

	fmt.Println(message)
}
