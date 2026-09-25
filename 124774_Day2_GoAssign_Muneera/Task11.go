package main

import "fmt"

func main() {
	employees := map[int]string{
		1: "Muneera",
		2: "Ayesha",
		3: "Amaira",
	}
	var id int
	fmt.Println("Enter employee id:")
	fmt.Scanf("%d", &id)

	if id == 1 {
		fmt.Println("Employeee:", employees[1])
	} else if id == 2 {
		fmt.Println("Employeee:", employees[2])
	} else if id == 3 {
		fmt.Println("Employeee:", employees[2])
	} else {
		fmt.Println("Employeee not found")
	}
}
