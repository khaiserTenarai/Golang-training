package view

import (
	"fmt"

	"employee-management/controller"
	"employee-management/module"
)

func Start() {

	for {
		fmt.Println()
		fmt.Println("===== Employee Management System =====")
		fmt.Println("1. Add Employee")
		fmt.Println("2. View All Employees")
		fmt.Println("3. Find Employee")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			addEmployee()

		case 2:
			viewEmployees()

		case 3:
			findEmployee()

		case 4:
			deleteEmployee()

		case 5:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func addEmployee() {

	var employee module.Employee

	fmt.Print("Enter ID: ")
	fmt.Scanln(&employee.ID)

	fmt.Print("Enter Name: ")
	fmt.Scanln(&employee.Name)

	fmt.Print("Enter Age: ")
	fmt.Scanln(&employee.Age)

	fmt.Print("Enter Salary: ")
	fmt.Scanln(&employee.Salary)

	fmt.Print("Enter Position: ")
	fmt.Scanln(&employee.Position)

	err := controller.AddEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully")
}

func viewEmployees() {

	employees := controller.GetAllEmployees()

	if len(employees) == 0 {
		fmt.Println("No employees found")
		return
	}

	fmt.Println()
	fmt.Println("ID\tName\tAge\tSalary\tPosition")

	for i := 0; i < len(employees); i++ {

		fmt.Printf(
			"%d\t%s\t%d\t%.2f\t%s\n",
			employees[i].ID,
			employees[i].Name,
			employees[i].Age,
			employees[i].Salary,
			employees[i].Position,
		)
	}
}

func findEmployee() {

	var id int

	fmt.Print("Enter employee ID: ")
	fmt.Scanln(&id)

	employee, err := controller.GetEmployeeByID(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println()
	fmt.Println("Employee Details")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Salary:", employee.Salary)
	fmt.Println("Position:", employee.Position)
}

func deleteEmployee() {

	var id int

	fmt.Print("Enter employee ID: ")
	fmt.Scanln(&id)

	err := controller.DeleteEmployee(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee deleted successfully")
}
