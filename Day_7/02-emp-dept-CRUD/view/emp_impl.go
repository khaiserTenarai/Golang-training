package view

import (
	"fmt"

	"department-management/model"
)

type EmployeeViewImpl struct {
}

func NewEmployeeView() EmployeeView {
	return &EmployeeViewImpl{}
}

func (v *EmployeeViewImpl) ShowMenu() int {

	fmt.Println("\n========== Employee Management ==========")
	fmt.Println("1. Save Employee")
	fmt.Println("2. Find Employee")
	fmt.Println("3. Find All Employees")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Delete Employee")
	fmt.Println("6. Exit")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *EmployeeViewImpl) ReadEmployee() model.Employee {

	var employee model.Employee

	fmt.Println("\n---------- Enter Employee ----------")

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Email: ")
	fmt.Scan(&employee.Email)

	fmt.Print("Enter Department ID: ")
	fmt.Scan(&employee.DepartmentID)

	return employee
}

func (v *EmployeeViewImpl) ReadEmployeeForUpdate() model.Employee {

	var employee model.Employee

	fmt.Println("\n---------- Update Employee ----------")

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&employee.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Email: ")
	fmt.Scan(&employee.Email)

	fmt.Print("Enter Department ID: ")
	fmt.Scan(&employee.DepartmentID)

	return employee
}

func (v *EmployeeViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	return id
}

func (v *EmployeeViewImpl) DisplayEmployee(
	employee model.Employee,
) {

	fmt.Println("\n---------- Employee ----------")

	fmt.Println("ID           :", employee.ID)
	fmt.Println("Name         :", employee.Name)
	fmt.Println("Age          :", employee.Age)
	fmt.Println("Email        :", employee.Email)
	fmt.Println("Department ID:", employee.DepartmentID)
}

func (v *EmployeeViewImpl) DisplayEmployees(
	employees []model.Employee,
) {

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Println("\n---------- Employees ----------")

	for _, employee := range employees {

		fmt.Println(
			employee.ID,
			employee.Name,
			employee.Age,
			employee.Email,
			employee.DepartmentID,
		)
	}
}
