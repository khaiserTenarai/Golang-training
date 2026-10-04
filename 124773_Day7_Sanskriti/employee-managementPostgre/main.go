package main

import (
	"context"
	"fmt"
	"log"

	"employee-managementPostgre/config"
	"employee-managementPostgre/controller"
	"employee-managementPostgre/database"
	"employee-managementPostgre/repository"
	"employee-managementPostgre/service"
	"employee-managementPostgre/view"
)

func main() {

	// Load configuration
	cfg := config.LoadConfig()

	// Connect to database
	conn, err := database.ConnectDB(cfg)

	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	defer conn.Close(context.Background())

	fmt.Println("Database connected successfully!")

	// Repository
	employeeRepo := &repository.EmployeeRepositoryImpl{
		DB: conn,
	}

	// Service
	employeeService := &service.EmployeeServiceImpl{
		Repository: employeeRepo,
	}

	// Controller
	employeeController := &controller.EmployeeController{
		Service: employeeService,
	}

	// Menu
	for {

		choice := view.ShowMenu()

		switch choice {

		case 1:
			addEmployee(employeeController)

		case 2:
			displayEmployees(employeeController)

		case 3:
			searchEmployee(employeeController)

		case 4:
			updateEmployee(employeeController)

		case 5:
			deleteEmployee(employeeController)

		case 6:
			fmt.Println("Thank you!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func addEmployee(controller *controller.EmployeeController) {

	employee := view.ReadEmployee()

	err := controller.AddEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully!")
	fmt.Println("Employee ID:", employee.ID)
}

func displayEmployees(controller *controller.EmployeeController) {

	employees, err := controller.GetEmployees()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	view.DisplayEmployees(employees)
}

func searchEmployee(controller *controller.EmployeeController) {

	id := view.ReadEmployeeID()

	employee, err := controller.GetEmployee(id)

	if err != nil {
		fmt.Println("Employee not found:", err)
		return
	}

	view.DisplayEmployee(employee)
}

func updateEmployee(controller *controller.EmployeeController) {

	id := view.ReadEmployeeID()

	employee, err := controller.GetEmployee(id)

	if err != nil {
		fmt.Println("Employee not found:", err)
		return
	}

	fmt.Println("\nEnter new employee details:")

	employee.Name = view.ReadString("Enter Name: ")
	employee.Email = view.ReadString("Enter Email: ")
	employee.Age = view.ReadInt("Enter Age: ")
	employee.Salary = view.ReadFloat("Enter Salary: ")
	employee.DepartmentID = view.ReadInt("Enter Department ID: ")

	err = controller.UpdateEmployee(&employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee updated successfully!")
}

func deleteEmployee(controller *controller.EmployeeController) {

	id := view.ReadEmployeeID()

	err := controller.DeleteEmployee(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee deleted successfully!")
}