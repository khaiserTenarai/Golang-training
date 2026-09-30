package view
import (
	"fmt"

	"employee-management/model"
)

type DepartmentViewImpl struct {
}

func NewDepartmentView() DepartmentView {
	return &DepartmentViewImpl{}
}

func (v *DepartmentViewImpl) ShowMenu() int {

	fmt.Println()
	fmt.Println("========== Department Management ==========")
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

	fmt.Print("Enter Department ID: ")
	fmt.Scan(&department.ID)

	fmt.Print("Enter Department Name: ")
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

	fmt.Println()
	fmt.Println("ID:", department.ID)
	fmt.Println("Name:", department.Name)
}

func (v *DepartmentViewImpl) DisplayDepartments(
	departments []model.Department,
) {

	fmt.Println()
	fmt.Println("---------- Departments ----------")

	if len(departments) == 0 {
		fmt.Println("No departments found.")
		return
	}

	for _, department := range departments {

		fmt.Printf(
			"ID: %d | Name: %s\n",
			department.ID,
			department.Name,
		)
	}
}
