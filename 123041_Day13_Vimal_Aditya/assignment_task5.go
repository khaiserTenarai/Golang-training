package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func developmentConfig(appName string, port int) Config {
	return Config{
		Env:      "development",
		AppName:  appName,
		Port:     port,
		LogLevel: "debug",
		Debug:    true,
		DBHost:   "localhost",
		DBName:   "employees_dev",
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := ask(reader, "Enter the app name: ")

	port := 8080
	if text := ask(reader, "Enter the port (press Enter for 8080): "); text != "" {
		if n, err := strconv.Atoi(text); err == nil {
			port = n
		} else {
			fmt.Println("not a number, using 8080")
		}
	}

	cfg := developmentConfig(name, port)
	fmt.Printf("%+v\n", cfg)
}
