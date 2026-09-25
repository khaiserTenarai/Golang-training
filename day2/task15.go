package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
}

var employeeMap = make(map[int]Employee)
var employeeIDs = []int{}
var reader = bufio.NewReader(os.Stdin)

func main() {
	addRecord(Employee{ID: 101, Name: "Aarav Sharma", Department: "Engineering", Salary: 85000})
	addRecord(Employee{ID: 102, Name: "Priya Patel", Department: "HR", Salary: 65000})

	for {

		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee by ID")
		fmt.Println("3. Display All Employees")
		fmt.Println("4. Filter Employees by Department")
		fmt.Println("5. Delete Employee")
		fmt.Println("6. Exit")
		fmt.Print("Choose an option (1-6): ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			addEmployeeHandler()
		case "2":
			searchEmployeeHandler()
		case "3":
			displayAllEmployees()
		case "4":
			filterByDepartmentHandler()
		case "5":
			deleteEmployeeHandler()
		case "6":
			fmt.Println("Exiting system. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addRecord(emp Employee) {
	employeeMap[emp.ID] = emp
	employeeIDs = append(employeeIDs, emp.ID)
}

func addEmployeeHandler() {
	fmt.Println("\n--- Add Employee ---")
	fmt.Print("Enter ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil || id <= 0 {
		fmt.Println("Error: Invalid ID.")
		return
	}

	if _, exists := employeeMap[id]; exists {
		fmt.Println("Error: Employee ID already exists.")
		return
	}

	fmt.Print("Enter Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Print("Enter Department: ")
	dept, _ := reader.ReadString('\n')
	dept = strings.TrimSpace(dept)

	fmt.Print("Enter Salary: ")
	salStr, _ := reader.ReadString('\n')
	salary, err := strconv.ParseFloat(strings.TrimSpace(salStr), 64)
	if err != nil {
		fmt.Println("Error: Invalid Salary.")
		return
	}

	addRecord(Employee{ID: id, Name: name, Department: dept, Salary: salary})
	fmt.Println("Employee added successfully!")
}

func searchEmployeeHandler() {

	fmt.Print("Enter Employee ID: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	emp, found := employeeMap[id]
	if !found {
		fmt.Println("Employee not found.")
		return
	}

	fmt.Printf("Found -> ID: %d | Name: %s | Dept: %s | Salary: $%.2f\n",
		emp.ID, emp.Name, emp.Department, emp.Salary)
}

func displayAllEmployees() {
	fmt.Println("\n--- All Employees ---")
	if len(employeeIDs) == 0 {
		fmt.Println("No employee records found.")
		return
	}

	for _, id := range employeeIDs {
		emp := employeeMap[id]
		fmt.Printf("ID: %-5d | Name: %-15s | Dept: %-12s | Salary: $%.2f\n",
			emp.ID, emp.Name, emp.Department, emp.Salary)
	}
}

func filterByDepartmentHandler() {
	fmt.Println("\n--- Filter by Department ---")
	fmt.Print("Enter Department Name: ")
	deptInput, _ := reader.ReadString('\n')
	targetDept := strings.ToLower(strings.TrimSpace(deptInput))

	count := 0
	for _, id := range employeeIDs {
		emp := employeeMap[id]
		if strings.ToLower(emp.Department) == targetDept {
			fmt.Printf("ID: %-5d | Name: %-15s | Salary: $%.2f\n", emp.ID, emp.Name, emp.Salary)
			count++
		}
	}

	if count == 0 {
		fmt.Println("No employees found in this department.")
	}
}

func deleteEmployeeHandler() {
	fmt.Println("\n--- Delete Employee ---")
	fmt.Print("Enter Employee ID to delete: ")
	idStr, _ := reader.ReadString('\n')
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	if _, found := employeeMap[id]; !found {
		fmt.Println("Employee not found.")
		return
	}

	delete(employeeMap, id)

	for i, val := range employeeIDs {
		if val == id {
			employeeIDs = append(employeeIDs[:i], employeeIDs[i+1:]...)
			break
		}
	}

	fmt.Println("Employee deleted successfully.")
}