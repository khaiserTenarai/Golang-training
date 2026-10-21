package config

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	DBHost              string
	DBPort              string
	DBUser              string
	DBPassword          string
	DBName              string
	ServerPort          string
	EnableNewValidation bool
}

// LoadConfig reads the local assignment_task15/.env configuration settings fallback parsing style
func LoadConfig() *Config {
	cfg := &Config{
		DBHost:              "localhost",
		DBPort:              "5432",
		DBUser:              "postgres",
		DBPassword:          "admin",
		DBName:              "employee_db",
		ServerPort:          "10015",
		EnableNewValidation: true,
	}

	file, err := os.Open(".env")
	if err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "#") || !strings.Contains(line, "=") {
				continue
			}
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				val := strings.TrimSpace(parts[1])
				switch key {
				case "DB_HOST":
					cfg.DBHost = val
				case "DB_PORT":
					cfg.DBPort = val
				case "DB_USER":
					cfg.DBUser = val
				case "DB_PASSWORD":
					cfg.DBPassword = val
				case "DB_NAME":
					cfg.DBName = val
				case "SERVER_PORT":
					cfg.ServerPort = val
				case "ENABLE_NEW_VALIDATION":
					cfg.EnableNewValidation = (val == "true")
				}
			}
		}
	}
	return cfg
}
