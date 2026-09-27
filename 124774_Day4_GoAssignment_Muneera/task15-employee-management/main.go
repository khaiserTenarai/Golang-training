package main

import (
	"errors"
	"fmt"
	"task15-employee-management/model"
	"task15-employee-management/service"
	"task15-employee-management/utility"
)

func main() {
	employee := model.Employee{
		Id:     101,
		Name:   "Muneera",
		Email:  "muneera@gmail.com",
		Age:    21,
		Salary: 20000,
	}

	err := service.AddEmployee(employee)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	employee, err = service.GetEmployee(101)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("\n Employee Details")
	fmt.Println("ID:", employee.Id)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Salary:", employee.Salary)

	err = service.DeleteEmployee(101)
	if err != nil {
		if errors.Is(err, utility.ErrEmployeeNotFound) {
			fmt.Println("Employee Not found")
		} else {
			fmt.Println("Error:", err)
		}
		return
	}
	fmt.Println("Employee deleted successfully")

}
