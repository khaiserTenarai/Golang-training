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
	Env            string
	Port           string
	LogLevel       string
	DBName         string
	UseInMemoryDB  bool
	RequestTimeout time.Duration
}

func testConfig() Config {
	return Config{
		Env:            "test",
		Port:           "0",
		LogLevel:       "error",
		DBName:         "employees_test",
		UseInMemoryDB:  true,
		RequestTimeout: 2 * time.Second,
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	cfg := testConfig()

	fmt.Printf("Test database name [%s]: ", cfg.DBName)
	nameText, _ := reader.ReadString('\n')
	if name := strings.TrimSpace(nameText); name != "" {
		cfg.DBName = name
	}

	fmt.Printf("Request timeout in seconds [%d]: ", int(cfg.RequestTimeout.Seconds()))
	secondsText, _ := reader.ReadString('\n')
	if seconds, err := strconv.Atoi(strings.TrimSpace(secondsText)); err == nil && seconds > 0 {
		cfg.RequestTimeout = time.Duration(seconds) * time.Second
	}

	if strings.Contains(strings.ToLower(cfg.DBName), "prod") {
		fmt.Println("\nRefusing to continue: a test configuration must never use a production database.")
		os.Exit(1)
	}

	fmt.Printf("\n%+v\n", cfg)
}
