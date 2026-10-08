package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"employee-management/controller"
	"employee-management/model"
)

type DepartmentViewImpl struct {
	controller controller.DepartmentController
	reader     *bufio.Reader
}

func NewDepartmentView(controller controller.DepartmentController) DepartmentView {
	return &DepartmentViewImpl{
		controller: controller,
		reader:     bufio.NewReader(os.Stdin),
	}
}

func (v *DepartmentViewImpl) Start() {
	for {

		fmt.Println("\n===== DEPARTMENT MANAGEMENT =====")
		fmt.Println("1. Create department")
		fmt.Println("2. Get department")
		fmt.Println("3. Get all departments")
		fmt.Println("4. Update department")
		fmt.Println("5. Delete department")
		fmt.Println("6. Exit")

		fmt.Print("Enter choice: ")

		choice, _ := strconv.Atoi(v.readInput())

		switch choice {

		case 1:
			v.CreateDepartment()

		case 2:
			v.GetDepartment()

		case 3:
			v.GetAllDepartments()

		case 4:
			v.UpdateDepartment()

		case 5:
			v.DeleteDepartment()

		case 6:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (v *DepartmentViewImpl) CreateDepartment() {
	var department model.Department

	fmt.Println("\n----- CREATE DEPARTMENT -----")

	fmt.Print("Enter department name: ")
	department.Name = v.readInput()

	err := v.controller.CreateDepartment(department)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Department created successfully")
}

func (v *DepartmentViewImpl) GetDepartment() {
	fmt.Println("\n----- GET DEPARTMENT -----")

	fmt.Print("Enter department ID: ")

	id, err := strconv.Atoi(v.readInput())

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	department, err := v.controller.GetDepartment(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayDepartment(*department)
}

func (v *DepartmentViewImpl) GetAllDepartments() {
	fmt.Println("\n----- GET ALL DEPARTMENTS -----")

	departments, err := v.controller.GetAllDepartments()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	v.DisplayDepartments(departments)
}

func (v *DepartmentViewImpl) UpdateDepartment() {
	var department model.Department

	fmt.Println("\n----- UPDATE DEPARTMENT -----")

	fmt.Print("Enter department ID: ")

	id, err := strconv.Atoi(v.readInput())

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	department.ID = id

	fmt.Print("Enter department name: ")
	department.Name = v.readInput()

	err = v.controller.UpdateDepartment(department)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Department updated successfully")
}

func (v *DepartmentViewImpl) DeleteDepartment() {
	fmt.Println("\n----- DELETE DEPARTMENT -----")

	fmt.Print("Enter department ID: ")

	id, err := strconv.Atoi(v.readInput())

	if err != nil {
		fmt.Println("Invalid ID")
		return
	}

	err = v.controller.DeleteDepartment(id)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Department deleted successfully")
}

func (v *DepartmentViewImpl) DisplayDepartment(department model.Department) {
	fmt.Println("\nDepartment Details")
	fmt.Println("ID     :", department.ID)
	fmt.Println("Name   :", department.Name)
}

func (v *DepartmentViewImpl) DisplayDepartments(departments []model.Department) {
	if len(departments) == 0 {
		fmt.Println("No departments found")
		return
	}

	for _, department := range departments {
		v.DisplayDepartment(department)
	}
}

func (v *DepartmentViewImpl) readInput() string {
	input, _ := v.reader.ReadString('\n')
	return strings.TrimSpace(input)
}
