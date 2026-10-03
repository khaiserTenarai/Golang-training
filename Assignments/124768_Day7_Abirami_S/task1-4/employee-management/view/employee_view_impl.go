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

type EmployeeViewImpl struct {
	controller controller.EmployeeController
	reader     *bufio.Reader
}

func NewEmployeeView(controller controller.EmployeeController) EmployeeView {
	return &EmployeeViewImpl{
		controller: controller,
		reader:     bufio.NewReader(os.Stdin),
	}
}
func (v *EmployeeViewImpl) Start() {
	for {
		fmt.Println("\n----- EMPLOYEE MANAGEMENT -----")
		fmt.Println("1. Create Employee")
		fmt.Println("2. Get Employee")
		fmt.Println("3. Get All Employees")
		fmt.Println("4. Update Employee")
		fmt.Println("5. Delete Employee")
		fmt.Println("6. Search Employees")
		fmt.Println("7. Exit")

		fmt.Print("Enter choice: ")
		choice := v.readInt()

		switch choice {
		case 1:
			v.CreateEmployee()
		case 2:
			v.GetEmployee()
		case 3:
			v.GetAllEmployees()
		case 4:
			v.UpdateEmployee()
		case 5:
			v.DeleteEmployee()
		case 6:
			v.SearchEmployees()
		case 7:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
func (v *EmployeeViewImpl) CreateEmployee() {
	var employee model.Employee

	fmt.Print("Enter Name: ")
	employee.Name = v.readString()

	fmt.Print("Enter Email: ")
	employee.Email = v.readString()

	fmt.Print("Enter Age: ")
	employee.Age = v.readInt()

	fmt.Print("Enter Salary: ")
	employee.Salary = v.readFloat()

	fmt.Print("Enter Department ID: ")
	employee.DepartmentID = v.readInt()

	err := v.controller.CreateEmployee(employee)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee created successfully")
}

func (v *EmployeeViewImpl) GetEmployee() {
	fmt.Print("Enter Employee ID: ")
	id := v.readInt()

	employee, err := v.controller.GetEmployee(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayEmployee(employee)
}

func (v *EmployeeViewImpl) GetAllEmployees() {
	employees, err := v.controller.GetAllEmployees()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayEmployees(employees)
}

func (v *EmployeeViewImpl) UpdateEmployee() {
	var employee model.Employee

	fmt.Print("Enter Employee ID: ")
	employee.ID = v.readInt()

	fmt.Print("Enter Name: ")
	employee.Name = v.readString()

	fmt.Print("Enter Email: ")
	employee.Email = v.readString()

	fmt.Print("Enter Age: ")
	employee.Age = v.readInt()

	fmt.Print("Enter Salary: ")
	employee.Salary = v.readFloat()

	fmt.Print("Enter Department ID: ")
	employee.DepartmentID = v.readInt()

	err := v.controller.UpdateEmployee(employee)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee updated successfully")
}

func (v *EmployeeViewImpl) DeleteEmployee() {
	fmt.Print("Enter Employee ID: ")
	id := v.readInt()

	err := v.controller.DeleteEmployee(id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee deleted successfully")
}

func (v *EmployeeViewImpl) SearchEmployees() {
	fmt.Println("\n----- SEARCH EMPLOYEES -----")

	fmt.Print("Enter name (press Enter to skip): ")
	name := v.readString()

	fmt.Print("Enter Department ID (0 to skip): ")
	departmentID := v.readInt()

	fmt.Print("Enter minimum Salary (0 to skip): ")
	salary := v.readFloat()

	fmt.Print("Enter Page Number: ")
	page := v.readInt()

	fmt.Print("Enter Page Size: ")
	size := v.readInt()

	fmt.Print("Enter Sort By (id/name/age/salary): ")
	sortBy := v.readString()

	fmt.Print("Enter Sort Order (ASC/DESC): ")
	sortOrder := v.readString()

	employees, err := v.controller.SearchEmployees(
		name,
		departmentID,
		salary,
		page,
		size,
		sortBy,
		sortOrder,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayEmployees(employees)
}

func (v *EmployeeViewImpl) DisplayEmployee(employee *model.Employee) {
	fmt.Println("\n----- EMPLOYEE DETAILS -----")
	fmt.Println("ID           :", employee.ID)
	fmt.Println("Name         :", employee.Name)
	fmt.Println("Email        :", employee.Email)
	fmt.Println("Age          :", employee.Age)
	fmt.Println("Salary       :", employee.Salary)
	fmt.Println("Department ID:", employee.DepartmentID)
}

func (v *EmployeeViewImpl) DisplayEmployees(employees []model.Employee) {
	fmt.Println("\n----- EMPLOYEE LIST -----")

	if len(employees) == 0 {
		fmt.Println("No employees found")
		return
	}

	for _, employee := range employees {
		fmt.Println("----------------------------")
		fmt.Println("ID           :", employee.ID)
		fmt.Println("Name         :", employee.Name)
		fmt.Println("Email        :", employee.Email)
		fmt.Println("Age          :", employee.Age)
		fmt.Println("Salary       :", employee.Salary)
		fmt.Println("Department ID:", employee.DepartmentID)
	}
}

func (v *EmployeeViewImpl) readString() string {
	value, _ := v.reader.ReadString('\n')
	return strings.TrimSpace(value)
}

func (v *EmployeeViewImpl) readInt() int {
	value := v.readString()
	number, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return number
}

func (v *EmployeeViewImpl) readFloat() float64 {
	value := v.readString()
	number, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0
	}
	return number
}
