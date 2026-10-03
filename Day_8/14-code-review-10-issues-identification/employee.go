package main

import "fmt"

// REVIEW ISSUE #1:
// Global variable usage.
// Makes testing difficult and can lead to race conditions.
var employees []Employee

type Employee struct {
	ID     int
	Name   string
	Salary float64
	Email  string
}

// REVIEW ISSUE #5:
// Function does not return an error.
// Invalid employee data can be added silently.
func AddEmployee(emp Employee) {

	// REVIEW ISSUE #2:
	// No validation for employee name.
	// Empty names are allowed.

	// REVIEW ISSUE #3:
	// No validation for salary.
	// Negative salary values are allowed.

	// REVIEW ISSUE #4:
	// No email validation.
	// Invalid emails such as "ganeshgmail.com" are accepted.

	// REVIEW ISSUE #9:
	// Shared slice updated without synchronization.
	// Concurrent goroutines can cause race conditions.
	employees = append(employees, emp)

	// REVIEW ISSUE #8:
	// Using fmt.Println instead of proper logging.
	fmt.Println("Employee added successfully")
}

func GetEmployee(id int) Employee {
	for _, emp := range employees {
		if emp.ID == id {
			return emp
		}
	}

	// REVIEW ISSUE #6:
	// Returning empty Employee does not clearly indicate
	// whether employee was not found.
	return Employee{}
}

func UpdateSalary(id int, salary float64) {
	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			employees[i].Salary = salary
			fmt.Println("Salary updated")
			return
		}
	}
}

func DeleteEmployee(id int) {
	for i := 0; i < len(employees); i++ {
		if employees[i].ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			fmt.Println("Employee deleted")
			return
		}
	}
}

func CalculateNetSalary(salary float64) float64 {

	// REVIEW ISSUE #7:
	// Hardcoded tax rate.
	// Business rules should be stored in constants/configuration.
	tax := salary * 0.10

	return salary - tax
}

func DisplayEmployees() {
	fmt.Println("\nEmployee List")
	fmt.Println("--------------")

	for _, emp := range employees {
		fmt.Printf(
			"ID: %d, Name: %s, Salary: %.2f, Email: %s\n",
			emp.ID,
			emp.Name,
			emp.Salary,
			emp.Email,
		)
	}
}

func main() {

	AddEmployee(Employee{
		ID: 1,

		// REVIEW ISSUE #2:
		// Empty employee name accepted.
		Name: "",

		// REVIEW ISSUE #3:
		// Negative salary accepted.
		Salary: -5000,

		// REVIEW ISSUE #4:
		// Invalid email format accepted.
		Email: "ganeshgmail.com",
	})

	AddEmployee(Employee{
		ID:     2,
		Name:   "Rahul",
		Salary: 50000,
		Email:  "rahul@gmail.com",
	})

	DisplayEmployees()

	UpdateSalary(2, 60000)

	emp := GetEmployee(2)

	fmt.Println("\nEmployee Found:")
	fmt.Println(emp)

	netSalary := CalculateNetSalary(emp.Salary)

	fmt.Println("\nNet Salary:", netSalary)

	DeleteEmployee(1)

	DisplayEmployees()

	// REVIEW ISSUE #10:
	// No unit tests are available for:
	// - AddEmployee
	// - GetEmployee
	// - UpdateSalary
	// - DeleteEmployee
	// - CalculateNetSalary
}
