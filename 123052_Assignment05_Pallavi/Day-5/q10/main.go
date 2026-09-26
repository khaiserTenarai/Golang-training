package main

import "fmt"

// Employee contains common employee information
type Employee struct {
	ID   int
	Name string
}

// Developer contains Employee
type Developer struct {
	Employee
	ProgrammingLanguage string
}

// Manager contains Employee
type Manager struct {
	Employee
	TeamSize int
}

func main() {

	// Developer uses Employee through composition
	developer := Developer{
		Employee: Employee{
			ID:   101,
			Name: "Pallavi",
		},
		ProgrammingLanguage: "Go",
	}

	// Manager uses Employee through composition
	manager := Manager{
		Employee: Employee{
			ID:   102,
			Name: "Rahul",
		},
		TeamSize: 10,
	}

	fmt.Println("Developer")
	fmt.Println("ID:", developer.ID)
	fmt.Println("Name:", developer.Name)
	fmt.Println("Language:", developer.ProgrammingLanguage)

	fmt.Println("\nManager")
	fmt.Println("ID:", manager.ID)
	fmt.Println("Name:", manager.Name)
	fmt.Println("Team Size:", manager.TeamSize)
}