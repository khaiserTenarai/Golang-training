package day8sanskriti
package main

import (
	"log"
	"os"
)

func setupLogger() {
	file, err := os.OpenFile(
		"application.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0666,
	)

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime)
}

func Info(message string) {
	log.Println("[INFO]", message)
}

func Warning(message string) {
	log.Println("[WARNING]", message)
}

func Error(message string) {
	log.Println("[ERROR]", message)
}

package main

func main() {
	setupLogger()

	Info("Application started")
	Info("Employee added")
	Warning("Employee salary is low")
	Error("Employee not found")
}