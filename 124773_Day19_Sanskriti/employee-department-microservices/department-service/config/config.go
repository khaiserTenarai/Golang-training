package config

import "os"

func Port() string {
	if v := os.Getenv("DEPARTMENT_PORT"); v != "" {
		return v
	}
	return "8082"
}
