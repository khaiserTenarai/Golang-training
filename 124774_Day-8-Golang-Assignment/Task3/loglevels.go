package main

import (
	"log"
	"os"
)

func main() {

	logger := log.New(
		os.Stdout,
		"EMPLOYEE-APP ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	logger.Println("INFO: application started")
	logger.Println("INFO: employee request received")

	logger.Println("WARNING: employee salary is very low")

	logger.Println("ERROR: unable to connect to database")

	logger.Println("INFO: application finished")
}
