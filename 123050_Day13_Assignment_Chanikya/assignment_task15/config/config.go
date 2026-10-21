package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DBHost               string
	DBPort               string
	DBUser               string
	DBName               string
	ServerPort           string
	EnableNewValidation  bool
}

type Secrets struct {
	DBPassword string
}

func LoadConfig() (*Config, *Secrets, error) {
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbName := getEnv("DB_NAME", "employee_db")
	serverPort := getEnv("SERVER_PORT", "10015")
	
	valStr := getEnv("ENABLE_NEW_VALIDATION", "true")
	enableVal, err := strconv.ParseBool(valStr)
	if err != nil {
		enableVal = true
	}

	cfg := &Config{
		DBHost:              dbHost,
		DBPort:              dbPort,
		DBUser:              dbUser,
		DBName:              dbName,
		ServerPort:          serverPort,
		EnableNewValidation:  enableVal,
	}

	// Secret isolated out of standard general configs
	sec := &Secrets{
		DBPassword: getEnv("DB_PASSWORD", "admin"),
	}

	if cfg.DBHost == "" || cfg.ServerPort == "" {
		return nil, nil, fmt.Errorf("invalid runtime target configuration criteria detected")
	}

	return cfg, sec, nil
}

func getEnv(key, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}
