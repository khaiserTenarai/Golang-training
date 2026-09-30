package view

import (
	"fmt"

	"attendance_leave/model"
)

type EmployeeView interface {
	ShowMenu() int
	ReadEmployee() model.Employee
	ReadID() int
	DisplayEmployee(employee model.Employee)
	DisplayEmployees(employees []model.Employee)
}

type EmployeeViewImpl struct {
}

func NewEmployeeView() EmployeeView {
	return &EmployeeViewImpl{}
}

func (v *EmployeeViewImpl) ShowMenu() int {

	fmt.Println("\n========== Employee Management ==========")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Find Employee")
	fmt.Println("3. Find All Employees")
	fmt.Println("4. Back")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *EmployeeViewImpl) ReadEmployee() model.Employee {

	var employee model.Employee

	fmt.Println("\n---------- Add Employee ----------")

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Email: ")
	fmt.Scan(&employee.Email)

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

	fmt.Println("ID    :", employee.ID)
	fmt.Println("Name  :", employee.Name)
	fmt.Println("Email :", employee.Email)
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
			"ID:",
			employee.ID,
			"| Name:",
			employee.Name,
			"| Email:",
			employee.Email,
		)
	}
}
