package view

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task04_salary_management/models"
)

type SalaryView struct {
	scanner *bufio.Scanner
}

func NewSalaryView() *SalaryView {
	return &SalaryView{scanner: bufio.NewScanner(os.Stdin)}
}

func (v *SalaryView) readLine() string {
	v.scanner.Scan()
	return strings.TrimSpace(v.scanner.Text())
}

func (v *SalaryView) ShowMenu() int {
	fmt.Println("\n===== SALARY MANAGEMENT =====")
	fmt.Println("1. Add Employee")
	fmt.Println("2. List Employees")
	fmt.Println("3. Update Salary (Transaction)")
	fmt.Println("4. View Salary History")
	fmt.Println("5. Exit")
	fmt.Print("Enter choice: ")
	choice, _ := strconv.Atoi(v.readLine())
	return choice
}

func (v *SalaryView) GetEmployeeInput() models.Employee {
	var emp models.Employee
	fmt.Print("Enter Name: ")
	emp.Name = v.readLine()
	fmt.Print("Enter Department: ")
	emp.Department = v.readLine()
	fmt.Print("Enter Salary: ")
	emp.Salary, _ = strconv.ParseFloat(v.readLine(), 64)
	return emp
}

func (v *SalaryView) GetID(label string) int {
	fmt.Printf("Enter %s ID: ", label)
	id, _ := strconv.Atoi(v.readLine())
	return id
}

func (v *SalaryView) ShowCurrentSalary(emp models.Employee) {
	fmt.Printf("\nEmployee: %s | Department: %s | Current Salary: %.2f\n", emp.Name, emp.Department, emp.Salary)
}

func (v *SalaryView) GetSalaryUpdateInput() (float64, string) {
	fmt.Print("Enter New Salary: ")
	salary, _ := strconv.ParseFloat(v.readLine(), 64)
	fmt.Print("Enter Reason for Change: ")
	reason := v.readLine()
	return salary, reason
}

func (v *SalaryView) ShowEmployees(employees []models.Employee) {
	if len(employees) == 0 {
		fmt.Println("\nNo employees found.")
		return
	}
	fmt.Printf("\n%-5s %-20s %-20s %-12s\n", "ID", "Name", "Department", "Salary")
	fmt.Println(strings.Repeat("-", 62))
	for _, e := range employees {
		fmt.Printf("%-5d %-20s %-20s %-12.2f\n", e.ID, e.Name, e.Department, e.Salary)
	}
}

func (v *SalaryView) ShowHistory(history []models.SalaryHistory) {
	if len(history) == 0 {
		fmt.Println("\nNo salary history found.")
		return
	}
	fmt.Printf("\n%-5s %-15s %-12s %-12s %-25s %-20s\n", "ID", "Employee", "Old Salary", "New Salary", "Reason", "Changed At")
	fmt.Println(strings.Repeat("-", 95))
	for _, h := range history {
		fmt.Printf("%-5d %-15s %-12.2f %-12.2f %-25s %-20s\n",
			h.ID, h.EmployeeName, h.OldSalary, h.NewSalary, h.ChangeReason, h.ChangedAt.Format("2006-01-02 15:04"))
	}
}

func (v *SalaryView) ShowSuccess(msg string) {
	fmt.Println("\n[SUCCESS]", msg)
}

func (v *SalaryView) ShowError(context string, err error) {
	fmt.Printf("\n[ERROR] %s: %v\n", context, err)
}
