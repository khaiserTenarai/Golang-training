package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task02_department_management/models"
)

type DepartmentView struct {
	scanner *bufio.Scanner
}

func NewDepartmentView() *DepartmentView {
	return &DepartmentView{scanner: bufio.NewScanner(os.Stdin)}
}

func (v *DepartmentView) readLine() string {
	v.scanner.Scan()
	return strings.TrimSpace(v.scanner.Text())
}

func (v *DepartmentView) ShowMenu() int {
	fmt.Println("\n===== DEPARTMENT MANAGEMENT =====")
	fmt.Println("1. Create Department")
	fmt.Println("2. List Departments")
	fmt.Println("3. Update Department")
	fmt.Println("4. Delete Department")
	fmt.Println("5. Add Employee")
	fmt.Println("6. List All Employees with Department")
	fmt.Println("7. List Employees by Department")
	fmt.Println("8. Assign Employee to Department")
	fmt.Println("9. Delete Employee")
	fmt.Println("10. Exit")
	fmt.Print("Enter choice: ")
	choice, _ := strconv.Atoi(v.readLine())
	return choice
}

func (v *DepartmentView) GetDepartmentInput() models.Department {
	var d models.Department
	fmt.Print("Enter Department Name: ")
	d.Name = v.readLine()
	fmt.Print("Enter Location: ")
	d.Location = v.readLine()
	return d
}

func (v *DepartmentView) GetEmployeeInput() models.Employee {
	var emp models.Employee
	fmt.Print("Enter Name: ")
	emp.Name = v.readLine()
	fmt.Print("Enter Email: ")
	emp.Email = v.readLine()
	fmt.Print("Enter Department ID (0 for none): ")
	deptID, _ := strconv.Atoi(v.readLine())
	if deptID > 0 {
		emp.DepartmentID = &deptID
	}
	fmt.Print("Enter Salary: ")
	emp.Salary, _ = strconv.ParseFloat(v.readLine(), 64)
	return emp
}

func (v *DepartmentView) GetID(label string) int {
	fmt.Printf("Enter %s ID: ", label)
	id, _ := strconv.Atoi(v.readLine())
	return id
}

func (v *DepartmentView) GetUpdateDeptInput(existing models.Department) models.Department {
	fmt.Printf("Name (%s) - new value (Enter to keep): ", existing.Name)
	if val := v.readLine(); val != "" {
		existing.Name = val
	}
	fmt.Printf("Location (%s) - new value (Enter to keep): ", existing.Location)
	if val := v.readLine(); val != "" {
		existing.Location = val
	}
	return existing
}

func (v *DepartmentView) ShowDepartments(depts []models.Department) {
	if len(depts) == 0 {
		fmt.Println("\nNo departments found.")
		return
	}
	fmt.Printf("\n%-5s %-25s %-20s\n", "ID", "Name", "Location")
	fmt.Println(strings.Repeat("-", 55))
	for _, d := range depts {
		fmt.Printf("%-5d %-25s %-20s\n", d.ID, d.Name, d.Location)
	}
}

func (v *DepartmentView) ShowEmployees(employees []models.Employee) {
	if len(employees) == 0 {
		fmt.Println("\nNo employees found.")
		return
	}
	fmt.Printf("\n%-5s %-20s %-25s %-20s %-10s\n", "ID", "Name", "Email", "Department", "Salary")
	fmt.Println(strings.Repeat("-", 85))
	for _, e := range employees {
		fmt.Printf("%-5d %-20s %-25s %-20s %-10.2f\n", e.ID, e.Name, e.Email, e.DepartmentName, e.Salary)
	}
}

func (v *DepartmentView) ShowSuccess(msg string) {
	fmt.Println("\n[SUCCESS]", msg)
}

func (v *DepartmentView) ShowError(context string, err error) {
	fmt.Printf("\n[ERROR] %s: %v\n", context, err)
}
