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
	DBPassword     string
	RequestTimeout time.Duration
}

func setting(reader *bufio.Reader, key string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	fmt.Printf("%s is not set - type a value (or press Enter to leave it empty): ", key)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func productionConfig(reader *bufio.Reader) Config {
	logLevel := setting(reader, "LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	return Config{
		Env:            "production",
		Port:           "8080",
		Debug:          false,
		LogLevel:       logLevel,
		DBHost:         setting(reader, "DB_HOST"),
		DBPassword:     setting(reader, "DB_PASSWORD"),
		RequestTimeout: 10 * time.Second,
	}
}

func (c Config) Validate() []string {
	var problems []string

	if c.DBHost == "" || c.DBHost == "localhost" {
		problems = append(problems, "DB_HOST must point at the real production database, not localhost")
	}
	if len(c.DBPassword) < 12 {
		problems = append(problems, "DB_PASSWORD must be at least 12 characters")
	}
	if c.Debug {
		problems = append(problems, "debug mode must be off in production")
	}
	if c.LogLevel == "debug" {
		problems = append(problems, "LOG_LEVEL=debug is too noisy (and may leak data) in production")
	}
	return problems
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	cfg := productionConfig(reader)

	if problems := cfg.Validate(); len(problems) > 0 {
		fmt.Println("\nNot safe to start in production:")
		for _, p := range problems {
			fmt.Println(" -", p)
		}
		os.Exit(1)
	}

	fmt.Printf("\nProduction config accepted: env=%s port=%s db=%s log=%s password=%s\n",
		cfg.Env, cfg.Port, cfg.DBHost, cfg.LogLevel, strings.Repeat("*", len(cfg.DBPassword)))
}
