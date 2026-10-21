package main

import "fmt"

// Task 1: Basic Layered Architecture Setup
// Layered architecture separates concerns into discrete levels of responsibility.
// Typically: Presentation (UI/API) -> Business Logic (Service) -> Data Access (Repository).

type PresentationLayer struct {
	Service BusinessLayer
}

type BusinessLayer struct {
	Data string
}

func main() {
	// Setup simple dependencies directly representing static layers
	bl := BusinessLayer{Data: "Layered Data Setup"}
	pl := PresentationLayer{Service: bl}
	fmt.Printf("Task 1 - Presentation rendering: %s\n", pl.Service.Data)
}
