package main

import (
	"log"
	"os"
)

func main() {
	// Task 3: Simple log levels using separate standard loggers.
	info := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warning := log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

	info.Println("Application started")
	warning.Println("Employee department is missing")
	errorLog.Println("Could not process sample employee")
}
