package view

import (
	"fmt"

	"department-management/model"
)

func ShowEmployee(
	employee *model.EmployeeWithDepartment,
) {

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("Employee Details")
	fmt.Println("======================================")

	fmt.Println("ID              :", employee.ID)
	fmt.Println("Name            :", employee.Name)
	fmt.Println("Email           :", employee.Email)
	fmt.Printf(
		"Salary          : %.2f\n",
		employee.Salary,
	)
	fmt.Println("Department ID   :", employee.DepartmentID)
	fmt.Println("Department Name :", employee.DepartmentName)

	fmt.Println("======================================")
}

func ShowEmployees(
	employees []model.EmployeeWithDepartment,
) {

	fmt.Println()
	fmt.Println("==============================================================")
	fmt.Println("Employees")
	fmt.Println("==============================================================")

	if len(employees) == 0 {
		fmt.Println("No employees found.")
		return
	}

	fmt.Printf(
		"%-5s %-18s %-25s %-12s %-15s\n",
		"ID",
		"NAME",
		"EMAIL",
		"SALARY",
		"DEPARTMENT",
	)

	fmt.Println("--------------------------------------------------------------")

	for _, employee := range employees {

		fmt.Printf(
			"%-5d %-18s %-25s %-12.2f %-15s\n",
			employee.ID,
			employee.Name,
			employee.Email,
			employee.Salary,
			employee.DepartmentName,
		)
	}

	fmt.Println("==============================================================")
}
