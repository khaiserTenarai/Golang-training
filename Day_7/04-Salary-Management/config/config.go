package config
import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	AppPort string
	AppEnv  string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
}

func Load() Config {

	err := godotenv.Load("env/config.env")

	if err != nil {
		fmt.Println("Error loading config.env:", err)
	}

	return Config{
		AppName: get("APP_NAME"),
		AppPort: get("APP_PORT"),
		AppEnv:  get("APP_ENV"),

		DBHost:     get("DB_HOST"),
		DBPort:     get("DB_PORT"),
		DBUser:     get("DB_USER"),
		DBPassword: get("DB_PASSWORD"),
		DBName:     get("DB_NAME"),
		DBSSLMode:  get("DB_SSLMODE"),
	}
}

func (c Config) DatabaseURL() string {

	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
}

func get(key string) string {

	return os.Getenv(key)
}
