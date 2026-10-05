package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Salary float64
}

func producer(employeeChannel chan<- Employee) {
	employees := []Employee{
		{Name: "max", Salary: 50000},
		{Name: "Rahul", Salary: 60000},
		{Name: "Suresh", Salary: 70000},
	}

	for _, employee := range employees {
		fmt.Println("Producer: sending", employee.Name)

		employeeChannel <- employee
	}

	close(employeeChannel)
}

func consumer(employeeChannel <-chan Employee) {
	for employee := range employeeChannel {
		fmt.Printf(
			"Consumer: processing %s with salary %.2f\n",
			employee.Name,
			employee.Salary,
		)

		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	employeeChannel := make(chan Employee, 2)

	go producer(employeeChannel)

	consumer(employeeChannel)

	fmt.Println("All employees processed")
}
