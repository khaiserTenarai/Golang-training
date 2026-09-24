// Command envconfig reads its configuration from environment variables,
// falling back to defaults (12-factor app style).
//
//	APP_NAME="Employee API" APP_PORT=9090 go run .
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Config holds application settings.
type Config struct {
	AppName string
	Port    int
}

// getEnv returns the value of key, or def if it is unset or empty.
func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok && v != "" {
		return v
	}
	return def
}

// LoadConfig builds a Config from the environment and validates it.
func LoadConfig() (Config, error) {
	portStr := getEnv("APP_PORT", "8080")
	port, err := strconv.Atoi(portStr)
	if err != nil || port < 1 || port > 65535 {
		return Config{}, fmt.Errorf("invalid APP_PORT %q: must be a number 1-65535", portStr)
	}
	return Config{
		AppName: getEnv("APP_NAME", "EmployeeApp"),
		Port:    port,
	}, nil
}

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Starting %s on port %d\n", cfg.AppName, cfg.Port)
}
