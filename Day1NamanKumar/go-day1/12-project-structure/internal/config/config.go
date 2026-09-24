// Package config loads application configuration.
// It is internal: only code inside this module can import it.
package config

import "os"

// Config holds runtime settings.
type Config struct {
	AppName string
	Env     string
}

// Load reads configuration from environment variables with defaults.
func Load() Config {
	c := Config{AppName: "EmployeeApp", Env: "development"}
	if v := os.Getenv("APP_NAME"); v != "" {
		c.AppName = v
	}
	if v := os.Getenv("APP_ENV"); v != "" {
		c.Env = v
	}
	return c
}
