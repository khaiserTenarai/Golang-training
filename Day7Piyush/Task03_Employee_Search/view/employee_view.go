package view

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"task03_employee_search/models"
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
	fmt.Println("\n===== EMPLOYEE SEARCH =====")
	fmt.Println("1. Add Employee")
	fmt.Println("2. Search Employees")
	fmt.Println("3. Exit")
	fmt.Print("Enter choice: ")
	choice, _ := strconv.Atoi(v.readLine())
	return choice
}

func (v *EmployeeView) GetEmployeeInput() models.Employee {
	var emp models.Employee
	fmt.Print("Enter Name: ")
	emp.Name = v.readLine()
	fmt.Print("Enter Department: ")
	emp.Department = v.readLine()
	fmt.Print("Enter Salary: ")
	emp.Salary, _ = strconv.ParseFloat(v.readLine(), 64)
	return emp
}

func (v *EmployeeView) GetSearchParams() models.SearchParams {
	var p models.SearchParams
	fmt.Print("Search by Name (Enter to skip): ")
	p.Name = v.readLine()
	fmt.Print("Search by Department (Enter to skip): ")
	p.Department = v.readLine()
	fmt.Print("Min Salary (0 to skip): ")
	p.MinSalary, _ = strconv.ParseFloat(v.readLine(), 64)
	fmt.Print("Max Salary (0 to skip): ")
	p.MaxSalary, _ = strconv.ParseFloat(v.readLine(), 64)
	fmt.Print("Sort by (name/department/salary/id): ")
	p.SortBy = v.readLine()
	fmt.Print("Sort order (ASC/DESC): ")
	p.SortOrder = v.readLine()
	fmt.Print("Page number: ")
	p.Page, _ = strconv.Atoi(v.readLine())
	fmt.Print("Page size: ")
	p.PageSize, _ = strconv.Atoi(v.readLine())
	return p
}

func (v *EmployeeView) ShowEmployees(employees []models.Employee, total, page, pageSize int) {
	if len(employees) == 0 {
		fmt.Println("\nNo employees found.")
		return
	}
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))
	fmt.Printf("\n--- Page %d of %d (Total: %d records) ---\n", page, totalPages, total)
	fmt.Printf("%-5s %-20s %-20s %-12s\n", "ID", "Name", "Department", "Salary")
	fmt.Println(strings.Repeat("-", 62))
	for _, e := range employees {
		fmt.Printf("%-5d %-20s %-20s %-12.2f\n", e.ID, e.Name, e.Department, e.Salary)
	}
}

func (v *EmployeeView) ShowSuccess(msg string) {
	fmt.Println("\n[SUCCESS]", msg)
}

func (v *EmployeeView) ShowError(context string, err error) {
	fmt.Printf("\n[ERROR] %s: %v\n", context, err)
}
