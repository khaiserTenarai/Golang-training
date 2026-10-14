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
	DBHost  string
}

type Secrets struct {
	DBPassword string
	APIKey     string
}

func (s Secrets) String() string {
	return "Secrets{DBPassword: ****, APIKey: ****}"
}

func ask(reader *bufio.Reader, label string) string {
	fmt.Print(label)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func secret(reader *bufio.Reader, envName, label string) string {
	if value := os.Getenv(envName); value != "" {
		fmt.Printf("%s read from the environment\n", envName)
		return value
	}
	return ask(reader, label)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	port, _ := strconv.Atoi(ask(reader, "Enter the port: "))
	cfg := Config{
		AppName: ask(reader, "Enter the app name: "),
		Port:    port,
		DBHost:  ask(reader, "Enter the database host: "),
	}

	secrets := Secrets{
		DBPassword: secret(reader, "DB_PASSWORD", "Enter the database password: "),
		APIKey:     secret(reader, "API_KEY", "Enter the API key: "),
	}

	fmt.Println()
	fmt.Printf("config  : %+v\n", cfg)
	fmt.Printf("secrets : %v\n", secrets)
	fmt.Printf("the database password has %d characters\n", len(secrets.DBPassword))
}
