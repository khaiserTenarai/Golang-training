package main

import "fmt"

type Address struct {
	Street  string
	City    string
	State   string
	ZipCode string
}

type Department struct {
	Name    string
	Manager string
	Floor   int
}

type Employee struct {
	ID       int
	FullName string
	Location Address
	Role     Department
}

func main() {
	
	emp := Employee{
		ID:       101,
		FullName: "lakshmi Shibu",
		
		
		Location: Address{
			Street:  "MG Street",
			City:    "Kochi",
			State:   "Kerala",
			ZipCode: "685618",
		},
		
		
		Role: Department{
			Name:    "Software Engineering",
			Manager: "Rahul",
			Floor:   4,
		},
	}

	
	fmt.Println("--- Employee Overview ---")
	fmt.Printf("Name: %s (ID: %d)\n", emp.FullName, emp.ID)
	
	
	fmt.Printf("Works in: %s (Floor %d)\n", emp.Role.Name, emp.Role.Floor)
	fmt.Printf("Reports to: %s\n", emp.Role.Manager)
	fmt.Printf("Lives in: %s, %s\n", emp.Location.City, emp.Location.State)

	fmt.Println("\n--- Updating Data ---")
	fmt.Println("Promoting Lakshmi to a new floor...")
	
	emp.Role.Floor = 5 
	
	fmt.Printf("Lakshmi is now on Floor %d.\n", emp.Role.Floor)
}