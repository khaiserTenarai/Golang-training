package internal

import (
	"fmt"
	"report/pkg"
)

func PrintSummary() {
	pkg.Title("--- REPORT ---")

	timestamps := [...]string{"10:00 am", "10:05 am", "10:12 am"}
	levels := [...]string{"info", "warning", "err"}
	messages := [...]string{"started", "running faster", "connection failed"}

	for index, time := range timestamps {
		fmt.Println("Time:   ", time)
		fmt.Println("Level:  ", levels[index])
		fmt.Println("Message:", messages[index])
		fmt.Println("--------------------")
	}
}

// Explaination: 
// internal/: It contains the core logic for the project module.
