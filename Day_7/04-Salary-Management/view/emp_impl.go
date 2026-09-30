package view
import (
	"fmt"

	"salary-management/model"
)

type EmployeeViewImpl struct {
}

func NewEmployeeView() EmployeeView {
	return &EmployeeViewImpl{}
}

func (v *EmployeeViewImpl) ShowMenu() int {

	fmt.Println()
	fmt.Println("========== Salary Management System ==========")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Find Employee")
	fmt.Println("3. Find All Employees")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Update Salary")
	fmt.Println("6. Salary History")
	fmt.Println("7. Delete Employee")
	fmt.Println("8. Exit")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *EmployeeViewImpl) ReadEmployee() model.Employee {

	var employee model.Employee

	fmt.Println()
	fmt.Println("---------- Add Employee ----------")

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Email: ")
	fmt.Scan(&employee.Email)

	fmt.Print("Enter Salary: ")
	fmt.Scan(&employee.Salary)

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

	fmt.Println()
	fmt.Println("---------- Employee ----------")

	fmt.Println("ID     :", employee.ID)
	fmt.Println("Name   :", employee.Name)
	fmt.Println("Age    :", employee.Age)
	fmt.Println("Email  :", employee.Email)
	fmt.Println("Salary :", employee.Salary)
}

func (v *EmployeeViewImpl) DisplayEmployees(
	employees []model.Employee,
) {

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Println()
	fmt.Println("---------- Employees ----------")

	for _, employee := range employees {

		fmt.Println(
			"ID:",
			employee.ID,
			"| Name:",
			employee.Name,
			"| Age:",
			employee.Age,
			"| Email:",
			employee.Email,
			"| Salary:",
			employee.Salary,
		)
	}
}

func (v *EmployeeViewImpl) ReadEmployeeForUpdate() model.Employee {

	var employee model.Employee

	fmt.Println()
	fmt.Println("---------- Update Employee ----------")

	fmt.Print("Enter ID: ")
	fmt.Scan(&employee.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Email: ")
	fmt.Scan(&employee.Email)

	return employee
}
