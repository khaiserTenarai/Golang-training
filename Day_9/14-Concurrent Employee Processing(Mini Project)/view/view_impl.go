package view
import (
	"fmt"

	"ems/model"
)

/*
	EmployeeViewImpl implements
	EmployeeView.
*/
type EmployeeViewImpl struct {
}

/*
	Constructor.
*/
func NewEmployeeView() EmployeeView {

	return &EmployeeViewImpl{}
}

/*
	ShowMenu displays the menu.
*/
func (v EmployeeViewImpl) ShowMenu() int {

	fmt.Println()
	fmt.Println("========== Employee Management System ==========")

	fmt.Println("1. Save Employee")

	fmt.Println("2. Find Employee")

	fmt.Println("3. Find All Employees")

	fmt.Println("4. Update Employee")

	fmt.Println("5. Delete Employee")

	/*
		Day 9 new option.
	*/
	fmt.Println("6. Process Employees Concurrently")

	fmt.Println("7. Exit")

	var choice int

	fmt.Print("Enter choice: ")

	fmt.Scan(&choice)

	return choice
}

/*
	ReadEmployee reads employee information.
*/
func (v EmployeeViewImpl) ReadEmployee() model.Employee {

	var employee model.Employee

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Employee Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Employee Email: ")
	fmt.Scan(&employee.Email)

	fmt.Print("Enter Employee Salary: ")
	fmt.Scan(&employee.Salary)

	return employee
}

/*
	Read employee data for update.
*/
func (v EmployeeViewImpl) ReadEmployeeForUpdate() model.Employee {

	var employee model.Employee

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&employee.ID)

	fmt.Print("Enter Employee Name: ")
	fmt.Scan(&employee.Name)

	fmt.Print("Enter Employee Age: ")
	fmt.Scan(&employee.Age)

	fmt.Print("Enter Employee Email: ")
	fmt.Scan(&employee.Email)

	fmt.Print("Enter Employee Salary: ")
	fmt.Scan(&employee.Salary)

	return employee
}

/*
	Read employee ID.
*/
func (v EmployeeViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter Employee ID: ")

	fmt.Scan(&id)

	return id
}

/*
	Display one employee.
*/
func (v EmployeeViewImpl) DisplayEmployee(
	employee model.Employee,
) {

	fmt.Println()
	fmt.Println("Employee ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Salary:", employee.Salary)
}

/*
	Display multiple employees.
*/
func (v EmployeeViewImpl) DisplayEmployees(
	employees []model.Employee,
) {

	fmt.Println()

	for _, employee := range employees {

		v.DisplayEmployee(employee)
	}
}

/*
	Display processing message.
*/
func (v EmployeeViewImpl) DisplayProcessingMessage() {

	fmt.Println()
	fmt.Println("Concurrent employee processing started...")
}