package main

import "fmt"

type Employee struct {
	id     int
	name   string
	salary int
}

func main() {
	employees := []Employee{
		{1, "Muneera", 30000},
		{2, "Rahul", 25000},
	}
	//Create
	employees = append(employees, Employee{3, "Anu", 28000})
	//Read
	fmt.Println("Employees:")
	for _, e := range employees {
		fmt.Println(e.id, e.name, e.salary)
	}
	//update
	for i := range employees {
		if employees[i].id == 2 {
			employees[i].salary = 30000
		}
	}
	fmt.Println("\n After Update:")
	for _, e := range employees {
		fmt.Println(e.id, e.name, e.salary)
	}
	//Delete
	for i := range employees {
		if employees[i].id == 1 {
			employees = append(employees[:i], employees[i+1:]...)
			break
		}
	}
	fmt.Println("\n After Delete:")
	for _, e := range employees {
		fmt.Println(e.id, e.name, e.salary)
	}

}
