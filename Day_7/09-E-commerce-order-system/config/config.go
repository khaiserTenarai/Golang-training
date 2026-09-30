package config
import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() Config {

	// Load env/config.env
	_ = godotenv.Load("env/config.env")

	return Config{
		DBHost:     get("DB_HOST", "localhost"),
		DBPort:     get("DB_PORT", "5432"),
		DBUser:     get("DB_USER", "postgres"),
		DBPassword: get("DB_PASSWORD", "Root"),
		DBName:     get("DB_NAME", "ecommerce_db"),
		DBSSLMode:  get("DB_SSLMODE", "disable"),
	}
}

func (c Config) DatabaseURL() string {

	return "postgres://" + c.DBUser + ":" + c.DBPassword +
		"@" + c.DBHost + ":" + c.DBPort + "/" +
		c.DBName + "?sslmode=" + c.DBSSLMode
}

func get(key string, defaultValue string) string {

	value := os.Getenv(key)

	if value == "" {
		return defaultValue
	}

	return value
}
