package main

import (
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "APP ", log.Ldate|log.Ltime|log.Lshortfile)

	logger.Println("Application started")

	empID := 101
	logger.Printf("Processing employee ID: %d", empID)

	if empID <= 0 {
		logger.Printf("Error: Invalid employee ID %d", empID)
	} else {
		logger.Printf("Successfully processed employee ID: %d", empID)
	}

	logger.Println("Application finished")
}