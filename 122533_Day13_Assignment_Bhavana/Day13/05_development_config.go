package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	Env            string
	Port           string
	Debug          bool
	LogLevel       string
	DBHost         string
	DBName         string
	RequestTimeout time.Duration
}

func developmentConfig() Config {
	return Config{
		Env:            "development",
		Port:           "8080",
		Debug:          true,
		LogLevel:       "debug",
		DBHost:         "localhost",
		DBName:         "employees_dev",
		RequestTimeout: 60 * time.Second,
	}
}

func override(reader *bufio.Reader, label, current string) string {
	fmt.Printf("%s [%s]: ", label, current)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return current
	}
	return text
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	cfg := developmentConfig()

	fmt.Println("Development configuration - press Enter to keep a value.")
	cfg.Port = override(reader, "Port", cfg.Port)
	cfg.DBHost = override(reader, "Database host", cfg.DBHost)
	cfg.DBName = override(reader, "Database name", cfg.DBName)

	fmt.Printf("\n%+v\n", cfg)
}
