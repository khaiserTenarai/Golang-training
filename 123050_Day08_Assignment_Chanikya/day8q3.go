package main

import (
	"log"
)

func info(message string) {
	log.Println("[INFO]", message)
}

func warning(message string) {
	log.Println("[WARNING]", message)
}

func errorLog(message string) {
	log.Println("[ERROR]", message)
}

func main() {

	info("Application started")

	info("Database connected successfully")

	warning("Employee salary is below expected range")

	errorLog("Failed to save employee")

	info("Application finished")
}
