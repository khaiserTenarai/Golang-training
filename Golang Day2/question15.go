package main

import (
	"fmt"
)

type Employee struct {
	ID   int
	Name string
	Role string
}


var dbMap = make(map[int]Employee)

var idList []int

func main() {
	var choice int

	for {
		fmt.Println("\n=== Day 2: Map & Slice Employee DB ===")
		fmt.Println("1. Add Employee")
		fmt.Println("2. Search Employee (Fast Map Lookup)")
		fmt.Println("3. List All (Ordered by Slice)")
		fmt.Println("4. Delete Employee")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option (1-5): ")
		
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addEmployee()
		case 2:
			searchEmployee()
		case 3:
			listEmployees()
		case 4:
			deleteEmployee()
		case 5:
			fmt.Println(" Shutting down")
			return
		default:
			fmt.Println(" Invalid choice")
		}
	}
}

func addEmployee() {
	var id int
	var name, role string

	fmt.Print("Enter ID: ")
	fmt.Scan(&id)

	
	_, exists := dbMap[id]
	if exists {
		fmt.Println("  An employee with this ID already exists")
		return
	}

	fmt.Print("Enter Name: ")
	fmt.Scan(&name)
	fmt.Print("Enter Role: ")
	fmt.Scan(&role)

	newEmp := Employee{ID: id, Name: name, Role: role}

	
	dbMap[id] = newEmp
	
	
	idList = append(idList, id)

	fmt.Println("Employee added successfully!")
}

func searchEmployee() {
	var id int
	fmt.Print("Enter ID to search: ")
	fmt.Scan(&id)

	
	emp, exists := dbMap[id]
	if exists {
		fmt.Printf(" Found: %s works as a %s (ID: %d)\n", emp.Name, emp.Role, emp.ID)
	} else {
		fmt.Println(" Employee not found in the map.")
	}
}

func listEmployees() {
	
	if len(idList) == 0 {
		fmt.Println(" No employees to display")
		return
	}

	fmt.Println("\n--- Employee Roster ---")
	
	
	for _, id := range idList {
		
		emp := dbMap[id]
		fmt.Printf("ID: %d | Name: %s | Role: %s\n", emp.ID, emp.Name, emp.Role)
	}
	fmt.Println("-----------------------")
}

func deleteEmployee() {
	var id int
	fmt.Print("Enter ID to delete: ")
	fmt.Scan(&id)

	_, exists := dbMap[id]
	if !exists {
		fmt.Println(" Cannot delete: ID not found.")
		return
	}

	
	delete(dbMap, id)

	
	for i, listID := range idList {
		if listID == id {
			
			idList = append(idList[:i], idList[i+1:]...)
			break 
		}
	}

	fmt.Println(" Employee removed from both Map and Slice!")
}