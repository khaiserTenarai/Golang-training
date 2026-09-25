package controller


import (
	"errors"
	"fmt"

	"employee-management/model"
	"employee-management/service"
	"employee-management/utility"
)

func AddEmployee(employee model.Employee) {

	err := service.AddEmployee(employee)

	if err != nil {

		var validationErr utility.ValidationError

		if errors.As(err, &validationErr) {

			fmt.Println(
				"Validation Error:",
				validationErr,
			)

			return
		}

		if errors.Is(err, utility.ErrDuplicateEmployee) {

			fmt.Println("Employee ID already exists")

			return
		}

		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Employee added successfully")
}

func GetEmployee(id int) {

	employee, err := service.GetEmployee(id)

	if err != nil {

		if errors.Is(err, utility.ErrEmployeeNotFound) {
			fmt.Println("Employee not found")
			return
		}

		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Employee Details")
	fmt.Println("-----------------")
	fmt.Println("ID:", employee.ID)
	fmt.Println("Name:", employee.Name)
	fmt.Println("Email:", employee.Email)
	fmt.Println("Age:", employee.Age)
	fmt.Println("Salary:", employee.Salary)
}

func DeleteEmployee(id int) {

	err := service.DeleteEmployee(id)

	if err != nil {

		if errors.Is(err, utility.ErrEmployeeNotFound) {
			fmt.Println("Employee not found")
			return
		}

		fmt.Println("Error:", err)

		return
	}

	fmt.Println("Employee deleted successfully")
}