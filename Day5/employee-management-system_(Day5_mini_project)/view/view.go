package view

import (
	"fmt"

	"employee-management/controller"
	"employee-management/model"
)

type EmployeeView struct {
	controller controller.EmployeeController
}

func NewEmployeeView(controller controller.EmployeeController) *EmployeeView {
	return &EmployeeView{
		controller: controller,
	}
}

func (v *EmployeeView) Start() {

	for {

		fmt.Println()
		fmt.Println("========== Employee Management ==========")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Get All Employees")
		fmt.Println("3. Get Employee By ID")
		fmt.Println("4. Update Employee")
		fmt.Println("5. Delete Employee")
		fmt.Println("6. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {

		case 1:
			v.addEmployee()

		case 2:
			v.getAllEmployees()

		case 3:
			v.getEmployeeByID()

		case 4:
			v.updateEmployee()

		case 5:
			v.deleteEmployee()

		case 6:
			fmt.Println("Thank you!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func (v *EmployeeView) addEmployee() {

	fmt.Println()
	fmt.Println("------ Add Employee ------")

	var employee model.Employee

	fmt.Print("Enter ID: ")
	fmt.Scan(&employee.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&employee.Salary)

	v.controller.AddEmployee(employee)
}

func (v *EmployeeView) getAllEmployees() {

	v.controller.GetAllEmployees()
}

func (v *EmployeeView) getEmployeeByID() {

	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	v.controller.GetEmployeeByID(id)
}

func (v *EmployeeView) updateEmployee() {

	fmt.Println()
	fmt.Println("------ Update Employee ------")

	var employee model.Employee

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&employee.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&employee.Salary)

	v.controller.UpdateEmployee(employee)
}

func (v *EmployeeView) deleteEmployee() {

	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	v.controller.DeleteEmployee(id)
}