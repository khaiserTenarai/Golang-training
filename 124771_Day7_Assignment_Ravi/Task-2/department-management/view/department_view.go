package view

import (
	"fmt"

	"department-management/model"
)

func ShowDepartment(department *model.Department) {

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("Department Details")
	fmt.Println("======================================")

	fmt.Println("ID          :", department.ID)
	fmt.Println("Name        :", department.Name)
	fmt.Println("Description :", department.Description)

	fmt.Println("======================================")
}

func ShowDepartments(
	departments []model.Department,
) {

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("Departments")
	fmt.Println("======================================")

	if len(departments) == 0 {
		fmt.Println("No departments found.")
		return
	}

	fmt.Printf(
		"%-5s %-20s %-30s\n",
		"ID",
		"NAME",
		"DESCRIPTION",
	)

	fmt.Println("--------------------------------------")

	for _, department := range departments {

		fmt.Printf(
			"%-5d %-20s %-30s\n",
			department.ID,
			department.Name,
			department.Description,
		)
	}

	fmt.Println("======================================")
}
