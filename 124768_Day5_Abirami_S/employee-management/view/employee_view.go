package view

import (
	"employee-management/controller"
	"employee-management/model"
	"fmt"
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
		fmt.Println("1. Add Employee")
		fmt.Println("2. View Employee")
		fmt.Println("3. View All Employee")
		fmt.Println("4. Update Employee")
		fmt.Println("5. Delete Employee")
		fmt.Println("6. Exit")

		var choice int
		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			v.addEmployee()
		case 2:
			v.viewEmployee()
		case 3:
			v.viewAllEmployee()
		case 4:
			v.updateEmployee()
		case 5:
			v.deleteEmployee()
		case 6:
			fmt.Println("Exiting")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (v *EmployeeView) addEmployee() {
	var employee model.Employee

	fmt.Println("Enter ID:")
	fmt.Scan(&employee.ID)

	fmt.Println("Enter Name:")
	fmt.Scan(&employee.Name)

	fmt.Println("Enter Email:")
	fmt.Scan(&employee.Email)

	fmt.Println("Enter Age:")
	fmt.Scan(&employee.Age)

	fmt.Println("Enter Salary:")
	fmt.Scan(&employee.Salary)

	err := v.controller.Add(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Employee added successfully")
}
func (v *EmployeeView) viewEmployee() {
	var id int
	fmt.Println("Enter ID:")
	fmt.Scan(&id)
	employee, err := v.controller.GetEmployeeById(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("\nEmployee Details")
	fmt.Println("ID: ", employee.ID)
	fmt.Println("Name: ", employee.Name)
	fmt.Println("Email: ", employee.Email)
	fmt.Println("Age: ", employee.Age)
	fmt.Println("Salary: ", employee.Salary)
}
func (v *EmployeeView) viewAllEmployee() {
	employees := v.controller.GetEmployee()
	if len(employees) == 0 {
		fmt.Println("No employees found")
		return
	}
	for _, employee := range employees {
		fmt.Println("ID: ", employee.ID)
		fmt.Println("Name: ", employee.Name)
		fmt.Println("Email: ", employee.Email)
		fmt.Println("Age: ", employee.Age)
		fmt.Println("Salary: ", employee.Salary)
	}
}
func (v *EmployeeView) updateEmployee() {
	var employee model.Employee

	fmt.Println("Enter ID:")
	fmt.Scan(&employee.ID)

	fmt.Println("Enter Name:")
	fmt.Scan(&employee.Name)

	fmt.Println("Enter Email:")
	fmt.Scan(&employee.Email)

	fmt.Println("Enter Age:")
	fmt.Scan(&employee.Age)

	fmt.Println("Enter Salary:")
	fmt.Scan(&employee.Salary)

	err := v.controller.UpdateEmp(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Employee updated successfully")
}
func (v *EmployeeView) deleteEmployee() {
	var id int
	fmt.Println("Enter ID:")
	fmt.Scan(&id)

	err := v.controller.DeleteEmp(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Employee deleted successfully")
}
