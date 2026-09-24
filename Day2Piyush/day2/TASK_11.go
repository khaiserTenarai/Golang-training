package main

import "fmt"

func main() {

	// Employee map
	employees := map[int]string{
		101: "Piyush",
		102: "Prachi",
		103: "yash",
		104: "Priya",
	}

	// Employee ID to search
	id := 103

	// Lookup employee
	name, found := employees[id]

	if found {
		fmt.Println("Employee Found:", name)
	} else {
		fmt.Println("Employee Not Found")
	}
}