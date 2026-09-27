package main

import "fmt"

type Employee struct {
	ID   int
	Name string
	Dept string
}

var employees []Employee

func addEmployee(e Employee) {
	employees = append(employees, e)
}

func getEmployee(id int) (Employee, bool) {
	for _, e := range employees {
		if e.ID == id {
			return e, true
		}
	}
	return Employee{}, false
}

func updateEmployee(id int, dept string) bool {
	for i, e := range employees {
		if e.ID == id {
			employees[i].Dept = dept
			return true
		}
	}
	return false
}

func deleteEmployee(id int) bool {
	for i, e := range employees {
		if e.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	addEmployee(Employee{ID: 1, Name: "Ranjitha", Dept: "Data Engineering"})
	addEmployee(Employee{ID: 2, Name: "Meera", Dept: "QA"})

	e, _ := getEmployee(1)
	fmt.Println("Found:", e)

	updateEmployee(1, "Analytics")
	e, _ = getEmployee(1)
	fmt.Println("After update:", e)

	deleteEmployee(2)
	fmt.Println("All employees:", employees)
}
