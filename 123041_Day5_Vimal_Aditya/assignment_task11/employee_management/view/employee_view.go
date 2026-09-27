package view

import (
	"employee_management/controller"
	"employee_management/models"
	"fmt"
)

type EmployeeView struct{
	controller *controller.EmployeeController
}

func NewEmployeeView(controller *controller.EmployeeController) *EmployeeView{
	return &EmployeeView{
		controller: controller,
	}
}

func (v *EmployeeView) showMenu(){
	fmt.Println("========================================")
	fmt.Println("       EMPLOYEE MANAGEMENT SYSTEM       ")
	fmt.Println("========================================")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Get Employee")
	fmt.Println("3. Get All Employees")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Delete Employee")
	fmt.Println("6. Exit")
	fmt.Println("========================================")
}

func (v *EmployeeView) addEmployee(){
	fmt.Println("Add Employee: ")
	var emp models.Employee

	fmt.Println("Enter ID: ")
	fmt.Scan(&emp.ID)

	fmt.Println("Enter Name: ")
	fmt.Scan(&emp.Name)

	fmt.Println("Enter Email: ")
	fmt.Scan(&emp.Email)

	fmt.Println("Enter Age: ")
	fmt.Scan(&emp.Age)

	fmt.Println("Enter Salary: ")
	fmt.Scan(&emp.Salary)

	if err := v.controller.AddEmployee(emp); err != nil{
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println("Employee Added successfully")
	}
}

func (v *EmployeeView) getEmployee(){
	fmt.Println("Get Employee: ")
	var id int 

	fmt.Print("Enter Employee ID: ")
	fmt.Scan(&id)

	if emp, err := v.controller.GetEmployee(id); err != nil{
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("ID: %d | Name: %s | Email: %s | Age: %d | Salary: %f | ", emp.ID, emp.Name, emp.Email, emp.Age, emp.Salary)
	}
}

func (v *EmployeeView) getAllEmployee(){
	fmt.Println("All Employees: ")
	employees, err := v.controller.GetAllEmployee()
	if err != nil{
		fmt.Printf("Error: %v", err)
	}
	if len(employees) == 0{
		fmt.Println("No records found")
	}
	for _, emp := range employees{
		fmt.Printf("ID: %d | Name: %s | Email: %s | Age: %d | Salary: %f | ", emp.ID, emp.Name, emp.Email, emp.Age, emp.Salary)
	}
}

func (v *EmployeeView) updateEmployee(){
	fmt.Println("Update Employee: ")

	var emp models.Employee

	fmt.Print("Enter ID to update: ")
	fmt.Scan(&emp.ID)

	fmt.Print("Enter New Name: ")
	fmt.Scan(&emp.Name)

	fmt.Print("Enter New Email: ")
	fmt.Scan(&emp.Email)

	fmt.Print("Enter New Age: ")
	fmt.Scan(&emp.Age)

	fmt.Print("Enter New Salary: ")
	fmt.Scan(&emp.Salary)

	if err := v.controller.UpdateEmployee(emp); err != nil{
		fmt.Printf("Error: %v",err)
	} else {
		fmt.Println("Updated Successfully")
	} 
}

func (v *EmployeeView) deleteEmployee(){
	fmt.Println("Delete Employee:")

	var id int

	fmt.Print("Enter ID to Delete: ")
	fmt.Scan(&id)

	if err := v.controller.DeleteEmployee(id); err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Println("Deleted Successfully")
	}
}

func (v *EmployeeView) Start(){
	for {
		v.showMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice{
		case 1: 
			v.addEmployee()
		case 2:
			v.getEmployee()
		case 3:
			v.getAllEmployee()
		case 4:
			v.updateEmployee()
		case 5: 
			v.deleteEmployee()
		case 6:
			fmt.Println("Thank You")
			return
		default:
			fmt.Println("Invalid Choice")
		}
		fmt.Println()
	}
}