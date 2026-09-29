package main

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/service"
	"employee-management/utility"
)

func main() {

	// Add Employee
	employee := model.Employee{
		ID:     1,
		Name:   "  Vittesh  ",
		Email:  "vittesh@gmail.com",
		Age:    23,
		Salary: 30000,
	}

	err := service.AddEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully")

	// Get Employee
	employee, err = service.GetEmployee(1)

	if err != nil {

		if errors.Is(err, utility.ErrEmployeeNotFound) {
			fmt.Println("Employee was not found")
		} else {
			fmt.Println("Error:", err)
		}

	} else {

		fmt.Println("Employee Details:")
		fmt.Println("ID:", employee.ID)
		fmt.Println("Name:", employee.Name)
		fmt.Println("Email:", employee.Email)
		fmt.Println("Age:", employee.Age)
		fmt.Println("Salary:", employee.Salary)
	}

	// Add duplicate employee
	err = service.AddEmployee(employee)

	if err != nil {

		if errors.Is(err, utility.ErrDuplicateEmployee) {
			fmt.Println("Duplicate employee detected")
		}
	}

	// Validation error
	invalidEmployee := model.Employee{
		ID:     2,
		Name:   "",
		Email:  "wrongemail",
		Age:    20,
		Salary: 20000,
	}

	err = service.AddEmployee(invalidEmployee)

	if err != nil {

		var validationError utility.ValidationError

		if errors.As(err, &validationError) {

			fmt.Println("Validation Error")
			fmt.Println("Field:", validationError.Field)
			fmt.Println("Message:", validationError.Message)

		} else {
			fmt.Println("Error:", err)
		}
	}

	// Delete Employee
	err = service.DeleteEmployee(1)

	if err != nil {
		fmt.Println("Delete Error:", err)
	} else {
		fmt.Println("Employee deleted successfully")
	}
}