package controller

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MainController struct {
	deptController *DepartmentController
	empController  *EmployeeController
	reader         *bufio.Reader
}

func NewMainController(
	deptController *DepartmentController,
	empController *EmployeeController,
) *MainController {
	return &MainController{
		deptController: deptController,
		empController:  empController,
		reader:         bufio.NewReader(os.Stdin),
	}
}

func (m *MainController) Start() {
	for {
		m.showMainMenu()
		choice := m.readInt("Enter your choice: ")

		switch choice {
		case 1:
			m.deptController.Start()
		case 2:
			m.empController.Start()
		case 3:
			fmt.Println("\nThank you for using Management Console System. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select 1, 2, or 3.")
		}
	}
}

func (m *MainController) showMainMenu() {
	fmt.Println()
	fmt.Println("==============================================")
	fmt.Println("         ENTERPRISE MANAGEMENT SYSTEM        ")
	fmt.Println("==============================================")
	fmt.Println("1. Department Management")
	fmt.Println("2. Employee Management")
	fmt.Println("3. Exit System")
	fmt.Println("==============================================")
}

func (m *MainController) readInt(message string) int {
	for {
		fmt.Print(message)
		val, err := m.reader.ReadString('\n')
		if err != nil {
			continue
		}
		val = strings.TrimSpace(val)
		num, err := strconv.Atoi(val)
		if err == nil {
			return num
		}
		fmt.Println("Please enter a valid integer.")
	}
}