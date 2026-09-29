package main

import "fmt"

type Employee struct {
	ID   int
	Name string
}

var employees []Employee 

func create(e Employee) {
	employees = append(employees, e)
}

func read(id int) (Employee, bool) {
	for _, e := range employees {
		if e.ID == id {
			return e, true
		}
	}
	return Employee{}, false
}

func update(id int, newName string) bool {
	for i, e := range employees {
		if e.ID == id {
			employees[i].Name = newName
			return true
		}
	}
	return false
}

func delete_(id int) bool {
	for i, e := range employees {
		if e.ID == id {
			employees = append(employees[:i], employees[i+1:]...)
			return true
		}
	}
	return false
}

func main() {
	create(Employee{ID: 1, Name: "Pooja"})
	create(Employee{ID: 2, Name: "Amruth"})
	fmt.Println("After create:", employees)

	e, _ := read(2)
	fmt.Println("Read ID 2   :", e)

	update(2, "Gokul Reddy")
	fmt.Println("After update:", employees)

	delete_(1)
	fmt.Println("After delete:", employees)
}
