
package main

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/service"
	"employee-management/utility"
)

func main() {
	empService := service.NewEmployeeService()

	emp1 := model.Employee{ID: 1, Name: "sachin ms", Email: "sachin@example.com", Age: 23, Salary: 35000}
	if err := empService.AddEmployee(emp1); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee added:", emp1.Name)
	}

	if err := empService.AddEmployee(emp1); err != nil {
		fmt.Println("Error:", err)

		if errors.Is(err, service.ErrDuplicateEmployee) {
			fmt.Println("-> This employee already exists in our records.")
		}
	}

	emp2 := model.Employee{ID: 2, Name: "Sneha Patil", Email: "not-an-email", Age: 25, Salary: 40000}
	if err := empService.AddEmployee(emp2); err != nil {
		fmt.Println("Error:", err)

		var valErr *utility.ValidationError
		if errors.As(err, &valErr) {
			fmt.Println("-> Problem with field:", valErr.Field)
		}
	}

	found, err := empService.GetEmployee(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Found employee:", found.Name)
	}

	_, err = empService.GetEmployee(99)
	if err != nil {
		fmt.Println("Error:", err)
		if errors.Is(err, service.ErrEmployeeNotFound) {
			fmt.Println("-> No employee with that ID.")
		}
	}

	if err := empService.DeleteEmployee(1); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Employee deleted successfully.")
	}

	if err := empService.DeleteEmployee(1); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Remaining employees:", empService.ListEmployees())
}
