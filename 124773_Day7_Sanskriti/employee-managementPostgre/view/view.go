package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"employee-managementPostgre/model"
)

var reader = bufio.NewReader(os.Stdin)

func ShowMenu() int {

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("     EMPLOYEE MANAGEMENT")
	fmt.Println("================================")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Display Employees")
	fmt.Println("3. Search Employee By ID")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Delete Employee")
	fmt.Println("6. Exit")
	fmt.Println("================================")

	return ReadInt("Enter choice: ")
}

func ReadInt(message string) int {

	for {
		fmt.Print(message)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		return value
	}
}

func ReadString(message string) string {

	fmt.Print(message)

	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func ReadFloat(message string) float64 {

	for {
		fmt.Print(message)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		value, err := strconv.ParseFloat(input, 64)

		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		return value
	}
}

func ReadEmployee() *model.Employee {

	employee := &model.Employee{}

	employee.Name = ReadString("Enter Name: ")

	employee.Email = ReadString("Enter Email: ")

	employee.Age = ReadInt("Enter Age: ")

	employee.Salary = ReadFloat("Enter Salary: ")

	employee.DepartmentID = ReadInt("Enter Department ID: ")

	return employee
}

func ReadEmployeeID() int {

	return ReadInt("Enter Employee ID: ")
}

func DisplayEmployee(employee model.Employee) {

	fmt.Println("--------------------------------")
	fmt.Println("ID           :", employee.ID)
	fmt.Println("Name         :", employee.Name)
	fmt.Println("Email        :", employee.Email)
	fmt.Println("Age          :", employee.Age)
	fmt.Println("Salary       :", employee.Salary)
	fmt.Println("Department ID:", employee.DepartmentID)
	fmt.Println("--------------------------------")
}

func DisplayEmployees(employees []model.Employee) {

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Println("\n========== EMPLOYEES ==========")

	for _, employee := range employees {
		DisplayEmployee(employee)
	}
}

func ShowMessage(message string) {

	fmt.Println(message)
}