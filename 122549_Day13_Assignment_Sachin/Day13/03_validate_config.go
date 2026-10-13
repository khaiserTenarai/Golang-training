package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	AppName        string
	Port           int
	LogLevel       string
	RequestTimeout time.Duration
}

func (c Config) Validate() []string {
	var problems []string

	if strings.TrimSpace(c.AppName) == "" {
		problems = append(problems, "APP_NAME can't be empty")
	}
	if c.Port < 1 || c.Port > 65535 {
		problems = append(problems, fmt.Sprintf("PORT must be between 1 and 65535 (got %d)", c.Port))
	}
	switch c.LogLevel {
	case "debug", "info", "warn", "error":
	default:
		problems = append(problems, fmt.Sprintf("LOG_LEVEL must be debug, info, warn or error (got %q)", c.LogLevel))
	}
	if c.RequestTimeout <= 0 {
		problems = append(problems, "REQUEST_TIMEOUT_SECONDS must be greater than 0")
	}
	return problems
}

// setting returns an environment variable if it's set. If it isn't,
// it asks the user to type the value in instead.
func setting(reader *bufio.Reader, key string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	fmt.Printf("%s is not set - type a value (or press Enter to leave it empty): ", key)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// a value that isn't a number just becomes 0 here, and Validate() complains about it
	port, _ := strconv.Atoi(setting(reader, "PORT"))
	seconds, _ := strconv.Atoi(setting(reader, "REQUEST_TIMEOUT_SECONDS"))

	cfg := Config{
		AppName:        setting(reader, "APP_NAME"),
		Port:           port,
		LogLevel:       setting(reader, "LOG_LEVEL"),
		RequestTimeout: time.Duration(seconds) * time.Second,
	}

	if problems := cfg.Validate(); len(problems) > 0 {
		fmt.Println("\nThe configuration is not valid:")
		for _, p := range problems {
			fmt.Println(" -", p)
		}
		os.Exit(1)
	}

	fmt.Println("\nConfiguration is valid, safe to start:", cfg)
}
