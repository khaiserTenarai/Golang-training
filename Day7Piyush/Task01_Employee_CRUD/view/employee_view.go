package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task01_employee_crud/models"
)

type EmployeeView struct {
	scanner *bufio.Scanner
}

func NewEmployeeView() *EmployeeView {
	return &EmployeeView{scanner: bufio.NewScanner(os.Stdin)}
}

func (v *EmployeeView) readLine() string {
	v.scanner.Scan()
	return strings.TrimSpace(v.scanner.Text())
}

func (v *EmployeeView) ShowMenu() int {
	fmt.Println("\n========== EMPLOYEE CRUD ==========")
	fmt.Println("1. Create Employee")
	fmt.Println("2. List All Employees")
	fmt.Println("3. Get Employee by ID")
	fmt.Println("4. Update Employee")
	fmt.Println("5. Delete Employee")
	fmt.Println("6. Exit")
	fmt.Print("Enter choice: ")
	choice, _ := strconv.Atoi(v.readLine())
	return choice
}

func (v *EmployeeView) GetEmployeeInput() models.Employee {
	var emp models.Employee
	fmt.Print("Enter Name: ")
	emp.Name = v.readLine()
	fmt.Print("Enter Email: ")
	emp.Email = v.readLine()
	fmt.Print("Enter Age: ")
	emp.Age, _ = strconv.Atoi(v.readLine())
	fmt.Print("Enter Department: ")
	emp.Department = v.readLine()
	fmt.Print("Enter Salary: ")
	emp.Salary, _ = strconv.ParseFloat(v.readLine(), 64)
	return emp
}

func (v *EmployeeView) GetID() int {
	fmt.Print("Enter Employee ID: ")
	id, _ := strconv.Atoi(v.readLine())
	return id
}

func (v *EmployeeView) GetUpdateInput(existing models.Employee) models.Employee {
	fmt.Printf("Name (%s) - new value (Enter to keep): ", existing.Name)
	if val := v.readLine(); val != "" {
		existing.Name = val
	}
	fmt.Printf("Email (%s) - new value (Enter to keep): ", existing.Email)
	if val := v.readLine(); val != "" {
		existing.Email = val
	}
	fmt.Printf("Age (%d) - new value (Enter to keep): ", existing.Age)
	if val := v.readLine(); val != "" {
		existing.Age, _ = strconv.Atoi(val)
	}
	fmt.Printf("Department (%s) - new value (Enter to keep): ", existing.Department)
	if val := v.readLine(); val != "" {
		existing.Department = val
	}
	fmt.Printf("Salary (%.2f) - new value (Enter to keep): ", existing.Salary)
	if val := v.readLine(); val != "" {
		existing.Salary, _ = strconv.ParseFloat(val, 64)
	}
	return existing
}

func (v *EmployeeView) ShowEmployee(emp models.Employee) {
	fmt.Printf("\nID: %d | Name: %s | Email: %s | Age: %d | Dept: %s | Salary: %.2f | Created: %s\n",
		emp.ID, emp.Name, emp.Email, emp.Age, emp.Department, emp.Salary, emp.CreatedAt.Format("2006-01-02 15:04"))
}

func (v *EmployeeView) ShowEmployees(employees []models.Employee) {
	if len(employees) == 0 {
		fmt.Println("\nNo employees found.")
		return
	}
	fmt.Printf("\n%-5s %-20s %-25s %-5s %-15s %-12s\n", "ID", "Name", "Email", "Age", "Department", "Salary")
	fmt.Println(strings.Repeat("-", 90))
	for _, e := range employees {
		fmt.Printf("%-5d %-20s %-25s %-5d %-15s %-12.2f\n", e.ID, e.Name, e.Email, e.Age, e.Department, e.Salary)
	}
}

func (v *EmployeeView) ShowSuccess(msg string) {
	fmt.Println("\n[SUCCESS]", msg)
}

func (v *EmployeeView) ShowError(context string, err error) {
	fmt.Printf("\n[ERROR] %s: %v\n", context, err)
}
