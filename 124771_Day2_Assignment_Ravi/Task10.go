package main
import "fmt"

type Employee struct {
	ID     int
	Name   string
	Salary float64
}

var employees []Employee

// Create
func addEmployee(e Employee) {
	employees = append(employees, e)
}

// Read
func getEmployees() []Employee {
	return employees
}

// Update
func updateEmployee(id int, name string, salary float64) {
	for i := range employees {
		if employees[i].ID == id {
			employees[i].Name = name
			employees[i].Salary = salary
			return
		}
	}
}

// Delete
func deleteEmployee(id int) {
	for i := range employees {
		if employees[i].ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			return
		}
	}
}

func main() {
	addEmployee(Employee{1, "Ravi", 50000})
	fmt.Println("Employees:", getEmployees())
	updateEmployee(1, "Amit", 55000)
	fmt.Println("After Update:", getEmployees())
	deleteEmployee(2)
	fmt.Println("After Delete:", getEmployees())
}
