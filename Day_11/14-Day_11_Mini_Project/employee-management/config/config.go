// Package config contains application configuration loading logic.
package config

// Import os to read environment variables.
import "os"

// Import godotenv to load values from the .env file.
import "github.com/joho/godotenv"

// Config contains all application configuration values.
type Config struct {
	// DBHost contains the PostgreSQL host name.
	DBHost string
	// DBPort contains the PostgreSQL port number.
	DBPort string
	// DBName contains the PostgreSQL database name.
	DBName string
	// DBUser contains the PostgreSQL user name.
	DBUser string
	// DBPassword contains the PostgreSQL password.
	DBPassword string
	// ServerPort contains the HTTP server port.
	ServerPort string
	// LogLevel contains the application log level.
	LogLevel string
}

// Load loads configuration from the .env file and environment variables.
func Load() Config {
	// Ignore the error because environment variables may already be available.
	_ = godotenv.Load()

	// Return a Config object populated from environment variables.
	return Config{
		// Read the PostgreSQL host.
		DBHost: os.Getenv("DB_HOST"),
		// Read the PostgreSQL port.
		DBPort: os.Getenv("DB_PORT"),
		// Read the PostgreSQL database name.
		DBName: os.Getenv("DB_NAME"),
		// Read the PostgreSQL username.
		DBUser: os.Getenv("DB_USER"),
		// Read the PostgreSQL password.
		DBPassword: os.Getenv("DB_PASSWORD"),
		// Read the application server port.
		ServerPort: os.Getenv("SERVER_PORT"),
		// Read the log level.
		LogLevel: os.Getenv("LOG_LEVEL"),
	}
}
