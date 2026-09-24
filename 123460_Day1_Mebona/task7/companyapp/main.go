package main

import (
	"companyapp/employee"
	"fmt"
)

func main() {
	
	message := employee.HireEmployee("John Doe", "Systems Administrator")
	
	fmt.Println(message)
}