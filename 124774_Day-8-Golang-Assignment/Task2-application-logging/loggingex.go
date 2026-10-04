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

	logger.Println("application started")
	logger.Println("connecting to database")

	// database connection code

	logger.Println("database connected successfully")
	logger.Println("processing employee request")

	// employee operation code

	logger.Println("employee request completed")
	logger.Println("application finished")
}
