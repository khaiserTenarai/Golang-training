package main

import (
	"log"
	"os"
)

func main() {
	
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		
		log.Fatal("Could not open log file: ", err)
	}

	defer file.Close()

	log.SetOutput(file)

	log.SetPrefix("MyApp: ")

	log.Println("Application started successfully.")
	log.Printf("User %s logged in from IP %s\n", "admin", "192.168.1.50")
	
	log.Println("WARNING: High memory usage detected.")
}