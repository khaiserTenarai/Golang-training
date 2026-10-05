package main

import "fmt"

func main() {
	employeeChannel := make(chan string)

	go func() {
		employeeChannel <- "max"
		employeeChannel <- "Rahul"
		employeeChannel <- "Suresh"

		close(employeeChannel)
	}()

	for employee := range employeeChannel {
		fmt.Println("Processing employee:", employee)
	}

	fmt.Println("All employees processed")
}
