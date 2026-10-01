package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"example.com/employee-management/model"
	"example.com/employee-management/repository"
	"example.com/employee-management/service"
)

type DepartmentController struct {
	service service.DepartmentService
	reader  *bufio.Reader
}

func NewDepartmentController(service service.DepartmentService) *DepartmentController {
	return &DepartmentController{
		service: service,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (c *DepartmentController) Start() {
	for {
		c.showMenu()
		choice := c.readInt("Enter your choice: ")
		switch choice {
		case 1:
			c.addDepartment()
		case 2:
			c.getDepartment()
		case 3:
			c.getAllDepartments()
		case 4:
			c.updateDepartment()
		case 5:
			c.deleteDepartment()
		case 6:
			fmt.Println("\nReturning to Main Menu...")
			return
		default:
			fmt.Println("Invalid choice. Please select 1 to 6.")
		}
		c.pause()
	}
}

func (c *DepartmentController) showMenu() {
	fmt.Println()
	fmt.Println("==============================================")
	fmt.Println("       DEPARTMENT MANAGEMENT SYSTEM")
	fmt.Println("==============================================")
	fmt.Println("1. Add Department")
	fmt.Println("2. Find Department")
	fmt.Println("3. Find All Departments")
	fmt.Println("4. Update Department")
	fmt.Println("5. Delete Department")
	fmt.Println("6. Back to Main Menu")
	fmt.Println("==============================================")
}

func (c *DepartmentController) addDepartment() {
	fmt.Println("\n---------- ADD DEPARTMENT ----------")
	dept := model.Department{}
	dept.Name = c.readString("Enter Department Name: ")
	dept.Code = c.readString("Enter Department Code: ")

	err := c.service.AddDepartment(&dept)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("\nDepartment added successfully.")
}

func (c *DepartmentController) getDepartment() {
	fmt.Println("\n---------- FIND DEPARTMENT ----------")
	id := c.readInt64("Enter Department ID: ")
	dept, err := c.service.GetDepartment(id)
	if errors.Is(err, repository.ErrDepartmentNotFound) {
		fmt.Println("Department not found.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	dept.Display()
}

func (c *DepartmentController) getAllDepartments() {
	fmt.Println("\n---------- ALL DEPARTMENTS ----------")
	departments, err := c.service.GetAllDepartments()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	if len(departments) == 0 {
		fmt.Println("No departments found.")
		return
	}
	fmt.Println()
	fmt.Printf("%-5s %-30s %-15s\n", "ID", "NAME", "CODE")
	fmt.Println(strings.Repeat("-", 55))
	for _, dept := range departments {
		fmt.Printf("%-5d %-30s %-15s\n", dept.ID, dept.Name, dept.Code)
	}
}

func (c *DepartmentController) updateDepartment() {
	fmt.Println("\n---------- UPDATE DEPARTMENT ----------")
	id := c.readInt64("Enter Department ID: ")
	dept, err := c.service.GetDepartment(id)
	if errors.Is(err, repository.ErrDepartmentNotFound) {
		fmt.Println("Department not found.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nCurrent department detail:")
	dept.Display()

	fmt.Println("\nEnter new values.")
	dept.Name = c.readString("Enter Department Name: ")
	dept.Code = c.readString("Enter Department Code: ")

	err = c.service.UpdateDepartment(dept)
	if errors.Is(err, repository.ErrDepartmentNotFound) {
		fmt.Println("Department not found.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Department updated successfully.")
}

func (c *DepartmentController) deleteDepartment() {
	fmt.Println("\n---------- DELETE DEPARTMENT ----------")
	id := c.readInt64("Enter Department ID: ")
	dept, err := c.service.GetDepartment(id)
	if errors.Is(err, repository.ErrDepartmentNotFound) {
		fmt.Println("Department not found.")
		return
	}
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	dept.Display()

	confirmation := c.readString("Are you sure you want to delete? (y/n): ")
	switch strings.ToLower(confirmation) {
	case "y", "yes":
		err = c.service.DeleteDepartment(id)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Department deleted successfully.")
	default:
		fmt.Println("Delete operation cancelled.")
	}
}

func (c *DepartmentController) readString(message string) string {
	for {
		fmt.Print(message)
		value, err := c.reader.ReadString('\n')
		if err != nil {
			continue
		}
		value = strings.TrimSpace(value)
		if value != "" {
			return value
		}
		fmt.Println("Value cannot be empty.")
	}
}

func (c *DepartmentController) readInt(message string) int {
	for {
		val := c.readString(message)
		num, err := strconv.Atoi(val)
		if err == nil {
			return num
		}
		fmt.Println("Please enter a valid integer.")
	}
}

func (c *DepartmentController) readInt64(message string) int64 {
	for {
		val := c.readString(message)
		num, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			return num
		}
		fmt.Println("Please enter a valid number.")
	}
}

func (c *DepartmentController) pause() {
	fmt.Println()
	fmt.Print("Press ENTER to continue...")
	_, _ = c.reader.ReadString('\n')
}