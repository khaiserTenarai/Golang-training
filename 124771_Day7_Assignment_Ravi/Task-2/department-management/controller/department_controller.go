package controller

import (
	"bufio"
	"context"
	"fmt"
	"strconv"
	"strings"

	"department-management/service"
	"department-management/view"
)

type DepartmentController struct {
	Service *service.DepartmentService
	Reader  *bufio.Reader
}

func NewDepartmentController(
	service *service.DepartmentService,
	reader *bufio.Reader,
) *DepartmentController {

	return &DepartmentController{
		Service: service,
		Reader:  reader,
	}
}

func (c *DepartmentController) Create() {

	fmt.Println()
	fmt.Println("===== Create Department =====")

	fmt.Print("Enter department name: ")
	name := c.readLine()

	fmt.Print("Enter description: ")
	description := c.readLine()

	department, err := c.Service.Create(
		context.Background(),
		name,
		description,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(
		"Department created successfully. ID:",
		department.ID,
	)
}

func (c *DepartmentController) GetByID() {

	fmt.Println()
	fmt.Println("===== Get Department =====")

	id := c.readInt(
		"Enter department ID: ",
	)

	department, err := c.Service.GetByID(
		context.Background(),
		id,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	view.ShowDepartment(department)
}

func (c *DepartmentController) GetAll() {

	fmt.Println()
	fmt.Println("===== Get All Departments =====")

	departments, err := c.Service.GetAll(
		context.Background(),
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	view.ShowDepartments(departments)
}

func (c *DepartmentController) Update() {

	fmt.Println()
	fmt.Println("===== Update Department =====")

	id := c.readInt(
		"Enter department ID: ",
	)

	fmt.Print("Enter new department name: ")
	name := c.readLine()

	fmt.Print("Enter new description: ")
	description := c.readLine()

	err := c.Service.Update(
		context.Background(),
		id,
		name,
		description,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(
		"Department updated successfully.",
	)
}

func (c *DepartmentController) Delete() {

	fmt.Println()
	fmt.Println("===== Delete Department =====")

	id := c.readInt(
		"Enter department ID: ",
	)

	err := c.Service.Delete(
		context.Background(),
		id,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(
		"Department deleted successfully.",
	)
}

func (c *DepartmentController) readLine() string {

	input, _ := c.Reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func (c *DepartmentController) readInt(
	message string,
) int {

	for {

		fmt.Print(message)

		input := c.readLine()

		value, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println(
				"Please enter a valid number.",
			)
			continue
		}

		return value
	}
}
