
package main

import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/service"
	"employee-management/utility"
)

func main() {
	svc := service.NewEmployeeService()

	ranjitha := model.Employee{
		ID:     1,
		Name:   "  Ranjitha   Ray ",
		Email:  "ray.s@example.com",
		Age:    27,
		Salary: 65000,
	}

	if err := svc.AddEmployee(ranjitha); err != nil {
		fmt.Println("Failed to add employee:", err)
	} else {
		fmt.Println("Employee added successfully.")
	}

	err := svc.AddEmployee(ranjitha)
	if errors.Is(err, service.ErrDuplicateEmployee) {
		fmt.Println("As expected, duplicate employee was rejected:", err)
	}

	found, err := svc.GetEmployee(1)
	if err == nil {
		fmt.Println("Fetched:", found)
	}

	_, err = svc.GetEmployee(99)
	if errors.Is(err, service.ErrEmployeeNotFound) {
		fmt.Println("As expected, employee 99 was not found:", err)
	}

	badEmployee := model.Employee{
		ID:     2,
		Name:   "Meera",
		Email:  "not-an-email",
		Age:    30,
		Salary: 40000,
	}
	err = svc.AddEmployee(badEmployee)
	var valErr *utility.ValidationError
	if errors.As(err, &valErr) {
		fmt.Println("Validation error caught -> Field:", valErr.Field, "Message:", valErr.Message)
	}

	if err := svc.DeleteEmployee(1); err == nil {
		fmt.Println("Employee 1 deleted.")
	}
}
