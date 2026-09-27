package view

import (
	"emp_simple/controller"
	"emp_simple/model"
	"fmt"
)

type EmpView struct {
	Controller *controller.EmployeeController
}

func (v *EmpView) ShowMenu() {
	for {
		fmt.Println("1. Add Employee")
		fmt.Println("2. Display Employee")
		fmt.Println("3. Delete Employee")
		fmt.Println("4. Exit")

		var choice int
		fmt.Println("Enter Your Choice")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			v.addemployee()

		case 2:
			v.displayemployee()

		case 3:
			v.deleteemployee()

		case 4:
			fmt.Println("Exit")
			return

		default:
			fmt.Println("Invalid")
		}
	}

}

func (v *EmpView) addemployee() {
	var employee model.Employee

	fmt.Println("Enter ID")
	fmt.Scan(&employee.ID)

	fmt.Println("Enter Name:")
	fmt.Scan(&employee.Name)

	fmt.Println("Enter Age")
	fmt.Scan(&employee.Age)

	fmt.Println("Enter Salary")
	fmt.Scan(&employee.Salary)

	err := v.Controller.AddEmp(&employee)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Employee Added Succefully")
}

func (v *EmpView) displayemployee() {
	employees := v.Controller.DisplayEmployee()

	if len(employees) == 0 {
		fmt.Println("No Employee Found")
		return

	}

	fmt.Println("--Emp List--")
	for _, employee := range employees {
		fmt.Println("ID", employee.ID)
		fmt.Println("Name", employee.Name)
		fmt.Println("Age", employee.Age)
		fmt.Println("Salary", employee.Salary)
	}
}

func (v *EmpView) deleteemployee() {
	var id int

	fmt.Println("Enter Id to delete")
	fmt.Scan(&id)

	if v.Controller.DeleteEmployee(id) {
		fmt.Println("Emp Deleted")

	} else {
		fmt.Println("Not found")
	}
}
