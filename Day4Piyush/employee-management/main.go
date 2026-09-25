package main

import (
	"employee-management/model"
	"employee-management/service"
	"employee-management/utility"
	"errors"
	"fmt"
)

func main() {
	employeeService := service.NewEmployeeService()

	// Add Employee
	employee := model.Employee{
		ID:     1,
		Name:   "   Piyush   ",
		Email:  "piyush@gmail.com",
		Age:    21,
		Salary: 50000,
	}

	err := employeeService.AddEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully")

	// Get Employee
	foundEmployee, err := employeeService.GetEmployee(1)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee:", foundEmployee)

	// Duplicate Employee Test
	duplicateEmployee := model.Employee{
		ID:     1,
		Name:   "Rahul",
		Email:  "rahul@gmail.com",
		Age:    25,
		Salary: 40000,
	}

	err = employeeService.AddEmployee(duplicateEmployee)

	if err != nil {
		fmt.Println("Duplicate Error:", err)

		if errors.Is(err, service.ErrDuplicateEmployee) {
			fmt.Println("Duplicate employee detected")
		}
	}

	// Get Non-existing Employee
	_, err = employeeService.GetEmployee(999)

	if err != nil {
		fmt.Println("Error:", err)

		if errors.Is(err, service.ErrEmployeeNotFound) {
			fmt.Println("Employee was not found")
		}
	}

	// Delete Employee
	err = employeeService.DeleteEmployee(1)

	if err != nil {
		fmt.Println("Delete Error:", err)
		return
	}

	fmt.Println("Employee deleted successfully")

	// Verify Employee is Deleted
	_, err = employeeService.GetEmployee(1)

	if err != nil {
		fmt.Println("After delete:", err)

		if errors.Is(err, service.ErrEmployeeNotFound) {
			fmt.Println("Confirmed: employee no longer exists")
		}
	}

	invalidEmployee := model.Employee{
		ID:     2,
		Name:   "Rahul",
		Email:  "invalid-email",
		Age:    25,
		Salary: 40000,
	}

	err = employeeService.AddEmployee(invalidEmployee)

	if err != nil {
		fmt.Println("Validation Error:", err)

		var validationErr utility.ValidationError

		if errors.As(err, &validationErr) {
			fmt.Println("Validation field:", validationErr.Field)
			fmt.Println("Validation message:", validationErr.Message)
		}
	}
}
