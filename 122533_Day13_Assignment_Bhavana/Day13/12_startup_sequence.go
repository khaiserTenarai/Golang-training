package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type step struct {
	name string
	undo func()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Which step should fail? (1-5, or just press Enter for none): ")
	text, _ := reader.ReadString('\n')
	failAt, _ := strconv.Atoi(strings.TrimSpace(text)) // anything invalid becomes 0 = none

	steps := []step{
		{"load configuration", nil},
		{"set up logging", func() { fmt.Println("  - flushing and closing the log file") }},
		{"connect to the database", func() { fmt.Println("  - closing the database connection") }},
		{"create services", func() { fmt.Println("  - stopping the services") }},
		{"start the HTTP server", func() { fmt.Println("  - stopping the HTTP server") }},
	}

	var started []step
	for i, s := range steps {
		number := i + 1
		fmt.Printf("[%d/%d] %s... ", number, len(steps), s.name)

		if number == failAt {
			fmt.Println("FAILED (simulated)")
			fmt.Println("\nRolling back what had already started:")
			for j := len(started) - 1; j >= 0; j-- {
				if started[j].undo != nil {
					started[j].undo()
				}
			}
			os.Exit(1)
		}

		fmt.Println("ok")
		started = append(started, s)
	}

	fmt.Println("\nStartup finished - the application is ready.")
}
