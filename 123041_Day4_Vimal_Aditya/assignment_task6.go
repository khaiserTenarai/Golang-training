package main

import "fmt"

func empIds() func() int{
	id := 0

	return func() int{
		id ++
		return id
	}
}

func main() {

	fmt.Println("\n**********************************************")
	fmt.Println("6. Create a closure for generating employee IDs.")
	fmt.Println("************************************************")

	var numberOfIDs int

	fmt.Print("Enter how many IDs to generate: ")
	fmt.Scan(&numberOfIDs)

	nextID := empIds()

	fmt.Println("Generated Employee IDs: ")
	for i := 0; i < numberOfIDs; i++ {
		generatedID := nextID()
		fmt.Println("Employee ID:", generatedID)
	}
}
