package main

import (
	"fmt"
	"os"
)

func main() {
	id := [5]string{"101", "102", "103"}
	name := [5]string{"Adam", "James", "Tom", "Allen", "Max"}
	salary := [5]int{30000, 60000, 25000, 75000, 40000}
	searchId := os.Args[1]
	for i := 0; i < len(id); i++ {
		if searchId == id[i] {
			fmt.Println("Employee ID: ", id[i])
			fmt.Println("Employee Name: ", name[i])
			fmt.Println("Employee Salary: ", salary[i])
			return
		}
	}
	fmt.Println("Employee Not Found")
}
