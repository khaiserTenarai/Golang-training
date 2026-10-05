package main

import "fmt"

func main() {
	employeeChannel := make(chan string)

	go func() {
		employeeChannel <- "Employee 1 completed"
		employeeChannel <- "Employee 2 completed"
		employeeChannel <- "Employee 3 completed"

		close(employeeChannel)
	}()

	for {
		employee, ok := <-employeeChannel

		if !ok {
			fmt.Println("Channel is closed")
			break
		}

		fmt.Println(employee)
	}
}
