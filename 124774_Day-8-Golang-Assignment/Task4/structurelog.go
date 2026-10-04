package main

import (
	"log"
	"os"
)

func main() {

	logger := log.New(
		os.Stdout,
		"EMPLOYEE-APP ",
		log.Ldate|log.Ltime,
	)

	employeeID := 101
	name := "Muneera"
	status := "success"

	logger.Printf(
		"employee_id=%d name=%s action=add status=%s",
		employeeID,
		name,
		status,
	)
}
