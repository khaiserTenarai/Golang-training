package day8sanskriti
package main

import "log"

func logEmployee(id int, name string, salary float64) {
	log.Printf(
		"[EMPLOYEE] id=%d name=%s salary=%.2f",
		id,
		name,
		salary,
	)
}

func main() {
	setupLogger()

	logEmployee(101, "Sanskriti", 50000)
}