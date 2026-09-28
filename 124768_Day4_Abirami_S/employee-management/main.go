package main

import (
	"employee-management/model"
	"employee-management/service"
	"errors"
	"fmt"
)

func main() {
	employee1 := model.Employee{
		ID:     1,
		Name:   "Tom",
		Email:  "tom@gmail.com",
		Age:    22,
		Salary: 80000,
	}

	err := service.AddEmployee(employee1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Employee added successfully")
	}

	err = service.AddEmployee(employee1)

	if errors.Is(err, service.ErrDuplicateEmployee) {
		fmt.Println("Duplicate Employee")
	}

	employee, err := service.GetEmployee(1)

	if errors.Is(err, service.ErrEmployeeNotFound) {
		fmt.Println("Employee not found")
	} else if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Employee: ", employee)
	}

	fmt.Print("All Employees: ")

	for _, emp := range service.GetAllEmployees() {
		fmt.Println(emp)
	}

	employee1.Salary = 90000

	err = service.UpdateEmployee(employee1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Employee updated successfully")
	}

	err = service.DeleteEmployee(1)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Employee deleted successfully")
	}

	_, err = service.GetEmployee(1)

	if errors.Is(err, service.ErrEmployeeNotFound) {
		fmt.Println("Employee not found")
	}
}
