package view

import (
	"fmt"

	"department-management/model"
)

type DepartmentViewImpl struct {
}

func NewDepartmentView() DepartmentView {
	return &DepartmentViewImpl{}
}

func (v *DepartmentViewImpl) ShowMenu() int {

	fmt.Println("\n========== Department Management ==========")
	fmt.Println("1. Save Department")
	fmt.Println("2. Find Department")
	fmt.Println("3. Find All Departments")
	fmt.Println("4. Update Department")
	fmt.Println("5. Delete Department")
	fmt.Println("6. Exit")

	var choice int

	fmt.Print("Enter choice: ")
	fmt.Scan(&choice)

	return choice
}

func (v *DepartmentViewImpl) ReadDepartment() model.Department {

	var department model.Department

	fmt.Println("\n---------- Enter Department ----------")

	fmt.Print("Enter Name: ")
	fmt.Scan(&department.Name)

	return department
}

func (v *DepartmentViewImpl) ReadDepartmentForUpdate() model.Department {

	var department model.Department

	fmt.Println("\n---------- Update Department ----------")

	fmt.Print("Enter ID: ")
	fmt.Scan(&department.ID)

	fmt.Print("Enter Name: ")
	fmt.Scan(&department.Name)

	return department
}

func (v *DepartmentViewImpl) ReadID() int {

	var id int

	fmt.Print("Enter ID: ")
	fmt.Scan(&id)

	return id
}

func (v *DepartmentViewImpl) DisplayDepartment(
	department model.Department,
) {

	fmt.Println("\n---------- Department ----------")

	fmt.Println("ID   :", department.ID)
	fmt.Println("Name :", department.Name)
}

func (v *DepartmentViewImpl) DisplayDepartments(
	departments []model.Department,
) {

	if len(departments) == 0 {
		fmt.Println("No departments found.")
		return
	}

	fmt.Println("\n---------- Departments ----------")

	for _, department := range departments {

		fmt.Println(
			department.ID,
			department.Name,
		)
	}
}
