package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Env      string
	AppName  string
	Port     int
	LogLevel string
	Debug    bool
	DBHost   string
	DBName   string
}

func ask(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func testConfig(appName, dbName string) Config {
	return Config{
		Env:      "test",
		AppName:  appName,
		Port:     8081,
		LogLevel: "error",
		Debug:    false,
		DBHost:   "in-memory",
		DBName:   dbName,
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := ask(reader, "Enter the app name: ")
	dbName := ask(reader, "Enter the test database name (press Enter for employees_test): ")
	if dbName == "" {
		dbName = "employees_test"
	}

	cfg := testConfig(name, dbName)
	fmt.Printf("%+v\n", cfg)
}
