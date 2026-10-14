package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type step struct {
	name string
	run  func() error
}

func ask(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	appName := ask(reader, "Enter the app name: ")
	failDB := strings.ToLower(ask(reader, "Should the database connection fail? (y/n): ")) == "y"

	steps := []step{
		{"load configuration", func() error {
			fmt.Println("   app name:", appName)
			return nil
		}},
		{"set up logger", func() error {
			fmt.Println("   logger ready")
			return nil
		}},
		{"connect to database", func() error {
			if failDB {
				return errors.New("connection refused")
			}
			fmt.Println("   database connected")
			return nil
		}},
		{"start http server", func() error {
			fmt.Println("   listening for requests")
			return nil
		}},
	}

	for i, s := range steps {
		fmt.Printf("[%d/%d] %s\n", i+1, len(steps), s.name)
		if err := s.run(); err != nil {
			fmt.Printf("startup failed at '%s': %v\n", s.name, err)
			os.Exit(1)
		}
	}

	fmt.Println(appName, "started successfully")
}
