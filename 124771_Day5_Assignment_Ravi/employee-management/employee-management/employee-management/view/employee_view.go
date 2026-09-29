package view

import (
	"fmt"

	"employee-management/controller"
	"employee-management/model"
)

type EmployeeView struct {
	controller *controller.EmployeeController
}

func NewEmployeeView(
	employeeController *controller.EmployeeController,
) *EmployeeView {

	return &EmployeeView{
		controller: employeeController,
	}
}

func (view *EmployeeView) Start() {

	for {

		view.ShowMenu()

		var choice int

		_, err := fmt.Scanf("%d", &choice)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			fmt.Scanln() 
			continue
		}

		switch choice {

		case 1:
			view.AddEmployee()

		case 2:
			view.GetEmployee()

		case 3:
			view.GetAllEmployees()

		case 4:
			view.UpdateEmployee()

		case 5:
			view.DeleteEmployee()

		case 6:
			view.SearchEmployee()

		case 7:
			fmt.Println("Thank you for using Employee Management System.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (view *EmployeeView) ShowMenu() {

	fmt.Println()
	fmt.Println("-----EMPLOYEE MANAGEMENT SYSTEM-----")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Get Employee")
	fmt.Println("3. Get All Employees")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Delete Employee")
	fmt.Println("6. Search Employee")
	fmt.Println("7. Exit")
	fmt.Print("Enter your choice: ")
}

func (view *EmployeeView) AddEmployee() {

	fmt.Println()
	fmt.Println("----- Add Employee -----")

	var id int
	var name string
	var age int
	var department string
	var salary float64
	var email string

	fmt.Print("Enter Employee ID: ")
	_, err := fmt.Scanf("%d", &id)
	if err != nil {
		fmt.Println("Invalid ID.")
		fmt.Scanln()
		return
	}

	fmt.Print("Enter Name: ")
	fmt.Scanf(" %[^\n]", &name)

	fmt.Print("Enter Age: ")
	_, err = fmt.Scanf("%d", &age)
	if err != nil {
		fmt.Println("Invalid age.")
		fmt.Scanln()
		return
	}

	fmt.Print("Enter Department: ")
	fmt.Scanf(" %[^\n]", &department)

	fmt.Print("Enter Salary: ")
	_, err = fmt.Scanf("%f", &salary)
	if err != nil {
		fmt.Println("Invalid salary.")
		fmt.Scanln()
		return
	}

	fmt.Print("Enter Email: ")
	fmt.Scanf(" %[^\n]", &email)

	employee := model.Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
		Salary:     salary,
		Email:      email,
	}

	err = view.controller.AddEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully.")
}

func (view *EmployeeView) GetEmployee() {

	fmt.Println()
	fmt.Println("----- Get Employee -----")

	var id int

	fmt.Print("Enter Employee ID: ")

	_, err := fmt.Scanf("%d", &id)

	if err != nil {
		fmt.Println("Invalid ID.")
		fmt.Scanln()
		return
	}

	employee, err := view.controller.GetEmployee(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	view.PrintEmployee(employee)
}

func (view *EmployeeView) GetAllEmployees() {

	fmt.Println()
	fmt.Println("----- All Employees -----")

	employees := view.controller.GetAllEmployees()

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for _, employee := range employees {
		view.PrintEmployee(employee)
	}
}

func (view *EmployeeView) UpdateEmployee() {

	fmt.Println()
	fmt.Println("----- Update Employee -----")

	var id int
	var name string
	var age int
	var department string
	var salary float64
	var email string

	fmt.Print("Enter Employee ID: ")

	_, err := fmt.Scanf("%d", &id)

	if err != nil {
		fmt.Println("Invalid ID.")
		fmt.Scanln()
		return
	}

	fmt.Print("Enter New Name: ")
	fmt.Scanf(" %[^\n]", &name)

	fmt.Print("Enter New Age: ")
	_, err = fmt.Scanf("%d", &age)

	if err != nil {
		fmt.Println("Invalid age.")
		fmt.Scanln()
		return
	}

	fmt.Print("Enter New Department: ")
	fmt.Scanf(" %[^\n]", &department)

	fmt.Print("Enter New Salary: ")
	_, err = fmt.Scanf("%f", &salary)

	if err != nil {
		fmt.Println("Invalid salary.")
		fmt.Scanln()
		return
	}

	fmt.Print("Enter New Email: ")
	fmt.Scanf(" %[^\n]", &email)

	employee := model.Employee{
		ID:         id,
		Name:       name,
		Age:        age,
		Department: department,
		Salary:     salary,
		Email:      email,
	}

	err = view.controller.UpdateEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee updated successfully.")
}

func (view *EmployeeView) DeleteEmployee() {

	fmt.Println()
	fmt.Println("----- Delete Employee -----")

	var id int

	fmt.Print("Enter Employee ID: ")

	_, err := fmt.Scanf("%d", &id)

	if err != nil {
		fmt.Println("Invalid ID.")
		fmt.Scanln()
		return
	}

	err = view.controller.DeleteEmployee(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee deleted successfully.")
}

func (view *EmployeeView) SearchEmployee() {

	fmt.Println()
	fmt.Println("----- Search Employee -----")

	var name string

	fmt.Print("Enter employee name: ")

	fmt.Scanf(" %[^\n]", &name)

	employees := view.controller.SearchEmployee(name)

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	for _, employee := range employees {
		view.PrintEmployee(employee)
	}
}

func (view *EmployeeView) PrintEmployee(
	employee model.Employee,
) {

	fmt.Println("------------------------------------")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Department:", employee.Department)
	fmt.Println("Salary:", employee.Salary)
	fmt.Println("Email:", employee.Email)
	fmt.Println("------------------------------------")
}