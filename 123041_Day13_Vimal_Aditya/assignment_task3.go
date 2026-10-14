package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppName string
	Port    int
	Env     string
}

func ask(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func (c Config) validate() []string {
	var problems []string

	if c.AppName == "" {
		problems = append(problems, "app name is required")
	}
	if c.Port < 1 || c.Port > 65535 {
		problems = append(problems, "port must be a number between 1 and 65535")
	}
	switch c.Env {
	case "development", "test", "production":
	default:
		problems = append(problems, "env must be development, test or production")
	}

	return problems
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := ask(reader, "Enter the app name: ")
	port, _ := strconv.Atoi(ask(reader, "Enter the port: "))
	env := strings.ToLower(ask(reader, "Enter the environment: "))

	cfg := Config{AppName: name, Port: port, Env: env}

	problems := cfg.validate()
	if len(problems) == 0 {
		fmt.Println("Configuration is valid")
		return
	}

	fmt.Println("Configuration is not valid:")
	for _, p := range problems {
		fmt.Println(" -", p)
	}
}
