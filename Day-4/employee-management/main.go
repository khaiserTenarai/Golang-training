package main

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/service"
	"employee-management/utility"
)

func main() {

	employeeService := service.NewEmployeeService()

	employee := model.Employee{
		ID:     101,
		Name:   "Pallavi",
		Email:  "pallavi@example.com",
		Age:    25,
		Salary: 50000,
	}

	// Add employee
	err := employeeService.AddEmployee(employee)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Employee added successfully")

	// Get employee
	foundEmployee, err := employeeService.GetEmployee(101)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nEmployee Details:")
	fmt.Println("ID:", foundEmployee.ID)
	fmt.Println("Name:", foundEmployee.Name)
	fmt.Println("Email:", foundEmployee.Email)
	fmt.Println("Age:", foundEmployee.Age)
	fmt.Println("Salary:", foundEmployee.Salary)

	// errors.Is
	err = employeeService.DeleteEmployee(999)

	if errors.Is(err, service.ErrEmployeeNotFound) {
		fmt.Println("\nerrors.Is: Employee not found")
	}

	// errors.As
	err = employeeService.AddEmployee(
		model.Employee{
			ID:     102,
			Name:   "",
			Email:  "test@example.com",
			Age:    25,
			Salary: 40000,
		},
	)

	var validationError *utility.ValidationError

	if errors.As(err, &validationError) {
		fmt.Println("\nerrors.As:")
		fmt.Println("Field:", validationError.Field)
		fmt.Println("Message:", validationError.Message)
	}
}
