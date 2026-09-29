package main

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/service"
	"employee-management/utility"
)

func main() {

	emp1 := model.Employee{ID: 1, Name: "  Anita  ", Email: "anita@example.com", Age: 28, Salary: 45000}
	err := service.AddEmployee(emp1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee added successfully")
	}

	fmt.Println()

	err = service.AddEmployee(emp1)
	if errors.Is(err, service.ErrDuplicateEmployee) {
		fmt.Println("Duplicate employee detected:", err)
	}

	fmt.Println()

	badEmp := model.Employee{ID: 2, Name: "", Email: "bad-email", Age: 10, Salary: -500}
	err = service.AddEmployee(badEmp)

	var validationErr *utility.ValidationError
	if errors.As(err, &validationErr) {
		fmt.Println("Validation error on field:", validationErr.Field)
		fmt.Println("Reason:", validationErr.Message)
	}

	fmt.Println()

	found, err := service.GetEmployee(1)
	if err == nil {
		fmt.Println("Found employee:", found)
	}

	fmt.Println()

	_, err = service.GetEmployee(99)
	if errors.Is(err, service.ErrEmployeeNotFound) {
		fmt.Println("Lookup failed as expected:", err)
	}

	fmt.Println()

	err = service.DeleteEmployee(1)
	if err == nil {
		fmt.Println("Employee deleted successfully")
	}

	err = service.DeleteEmployee(1)
	if errors.Is(err, service.ErrEmployeeNotFound) {
		fmt.Println("Delete failed as expected:", err)
	}
}
