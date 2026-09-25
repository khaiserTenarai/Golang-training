package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"employee-management/controller"
	"employee-management/model"
)

var reader = bufio.NewReader(os.Stdin)

func Start() {

	for {

		fmt.Println()
		fmt.Println("===== Employee Management =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Get Employee")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. Exit")
		fmt.Print("Enter choice: ")

		choice := readInt()

		switch choice {

		case 1:
			addEmployee()

		case 2:
			getEmployee()

		case 3:
			deleteEmployee()

		case 4:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func addEmployee() {

	fmt.Print("Enter ID: ")
	id := readInt()

	fmt.Print("Enter Name: ")
	name := readString()

	fmt.Print("Enter Email: ")
	email := readString()

	fmt.Print("Enter Age: ")
	age := readInt()

	fmt.Print("Enter Salary: ")
	salary := readFloat()

	employee := model.Employee{
		ID:     id,
		Name:   name,
		Email:  email,
		Age:    age,
		Salary: salary,
	}

	controller.AddEmployee(employee)
}

func getEmployee() {

	fmt.Print("Enter Employee ID: ")

	id := readInt()

	controller.GetEmployee(id)
}

func deleteEmployee() {

	fmt.Print("Enter Employee ID: ")

	id := readInt()

	controller.DeleteEmployee(id)
}

func readString() string {

	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func readInt() int {

	input := readString()

	value, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Please enter a valid number")
		return 0
	}

	return value
}

func readFloat() float64 {

	input := readString()

	value, err := strconv.ParseFloat(input, 64)

	if err != nil {
		fmt.Println("Please enter a valid salary")
		return 0
	}

	return value
}