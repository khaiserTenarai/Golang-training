package main

import (
	"bufio"
	"errors"
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

func productionConfig(appName, dbHost string) (Config, error) {
	if dbHost == "" {
		return Config{}, errors.New("production needs a database host")
	}

	return Config{
		Env:      "production",
		AppName:  appName,
		Port:     80,
		LogLevel: "info",
		Debug:    false,
		DBHost:   dbHost,
		DBName:   "employees",
	}, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	name := ask(reader, "Enter the app name: ")
	host := ask(reader, "Enter the database host: ")

	cfg, err := productionConfig(name, host)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%+v\n", cfg)
}
