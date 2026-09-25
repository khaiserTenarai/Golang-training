package main

import "fmt"

func main() {
	employees := map[int]string{101: "Adam", 102: "James", 103: "Max"}
	var id int
	fmt.Println("Enter ID: ")
	fmt.Scan(&id)
	name, found := employees[id]
	if found {
		fmt.Println("Employee Name: ", name)
	} else {
		fmt.Println("Employee not found")
	}
}
