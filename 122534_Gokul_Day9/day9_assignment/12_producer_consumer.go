package main

import (
	"fmt"
	"time"
)

type queuedEmployee struct {
	ID   int
	Name string
}

func employeeProducer(out chan<- queuedEmployee) {
	names := []string{"Asha", "Ravi", "Meena", "Karthik", "Divya"}
	for i, n := range names {
		emp := queuedEmployee{ID: i + 1, Name: n}
		fmt.Println("Producer: created", emp.Name)
		out <- emp
		time.Sleep(200 * time.Millisecond)
	}
	close(out)
}

func employeeConsumer(in <-chan queuedEmployee, done chan<- bool) {
	for emp := range in {
		fmt.Printf("Consumer: processed employee %d (%s)\n", emp.ID, emp.Name)
	}
	done <- true
}

func main() {
	employees := make(chan queuedEmployee)
	done := make(chan bool)

	go employeeProducer(employees)
	go employeeConsumer(employees, done)

	<-done
	fmt.Println("All employees processed")
}
