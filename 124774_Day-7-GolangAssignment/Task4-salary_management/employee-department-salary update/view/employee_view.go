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

// ==================================================
// CONSTRUCTOR
// ==================================================

func NewEmployeeView() *EmployeeView {

	return &EmployeeView{
		reader: bufio.NewReader(os.Stdin),
	}
}

// ==================================================
// MENU
// ==================================================

func (v *EmployeeView) ShowMenu() {

	fmt.Println()
	fmt.Println("========== EMPLOYEE MANAGEMENT ==========")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Delete Employee")
	fmt.Println("3. Update Employee")
	fmt.Println("4. Find Employee By ID")
	fmt.Println("5. Find All Employees")
	fmt.Println("6. Search Employees")
	fmt.Println("7. Update Employee Salary")
	fmt.Println("8. Back to Main Menu")
	fmt.Println("=========================================")
}

// ==================================================
// READ INTEGER
// ==================================================

func (v *EmployeeView) ReadInt(message string) int {

	for {

		fmt.Print(message)

		input, err := v.reader.ReadString('\n')

		if err != nil {
			fmt.Println("Unable to read input.")
			continue
		}

		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)

		if err == nil {
			return value
		}

		fmt.Println("Please enter a valid number.")
	}
}

// ==================================================
// READ FLOAT
// ==================================================

func (v *EmployeeView) ReadFloat(message string) float64 {

	for {

		fmt.Print(message)

		input, err := v.reader.ReadString('\n')

		if err != nil {
			fmt.Println("Unable to read input.")
			continue
		}

		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)

		if err == nil {
			return value
		}

		fmt.Println("Please enter a valid salary.")
	}
}

// ==================================================
// READ STRING
// ==================================================

func (v *EmployeeView) ReadString(message string) string {

	fmt.Print(message)

	input, _ := v.reader.ReadString('\n')

	return strings.TrimSpace(input)
}

// ==================================================
// READ EMPLOYEE FOR ADD
// ==================================================

func (v *EmployeeView) ReadEmployee() model.Employee {

	return model.Employee{
		Name:         v.ReadString("Enter Name: "),
		Email:        v.ReadString("Enter Email: "),
		Age:          v.ReadInt("Enter Age: "),
		Salary:       v.ReadFloat("Enter Salary: "),
		DepartmentID: v.ReadInt("Enter Department ID: "),
	}
}

// ==================================================
// READ EMPLOYEE FOR UPDATE
// ==================================================

func (v *EmployeeView) ReadEmployeeForUpdate() model.Employee {

	return model.Employee{
		ID:           v.ReadInt("Enter Employee ID: "),
		Name:         v.ReadString("Enter Name: "),
		Email:        v.ReadString("Enter Email: "),
		Age:          v.ReadInt("Enter Age: "),
		Salary:       v.ReadFloat("Enter Salary: "),
		DepartmentID: v.ReadInt("Enter Department ID: "),
	}
}

// ==================================================
// SHOW ONE EMPLOYEE
// ==================================================

func (v *EmployeeView) ShowEmployee(
	employee model.Employee,
) {

	fmt.Println("--------------------------------")
	fmt.Println("ID            :", employee.ID)
	fmt.Println("Name          :", employee.Name)
	fmt.Println("Email         :", employee.Email)
	fmt.Println("Age           :", employee.Age)
	fmt.Println("Salary        :", employee.Salary)
	fmt.Println("Department ID :", employee.DepartmentID)
	fmt.Println("Department Name :", employee.DepartmentName)
	fmt.Println("--------------------------------")
}

// ==================================================
// SHOW ALL EMPLOYEES
// ==================================================

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

// ==================================================
// SHOW MESSAGE
// ==================================================

func (v *EmployeeView) ShowMessage(
	message string,
) {

	fmt.Println(message)
}
func (v *EmployeeView) ReadSearchCriteria() (
	string,
	string,
	float64,
) {

	name := v.ReadString(
		"Enter Name (leave empty for all): ",
	)

	department := v.ReadString(
		"Enter Department (leave empty for all): ",
	)

	minSalary := v.ReadFloat(
		"Enter Minimum Salary (0 for all): ",
	)

	return name, department, minSalary
}
