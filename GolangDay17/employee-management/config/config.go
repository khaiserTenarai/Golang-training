package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
	JWTSecret  string        
	JWTExpiry  time.Duration 
	DebugSQL   bool          
}

func Load() Config {
	_ = godotenv.Load()
	minutes, err := strconv.Atoi(os.Getenv("JWT_EXPIRY_MINUTES"))
	if err != nil || minutes <= 0 {
		minutes = 60
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "change-me-in-env-file"
	}
	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: os.Getenv("SERVER_PORT"),
		JWTSecret:  secret,
		JWTExpiry:  time.Duration(minutes) * time.Minute,
		DebugSQL:   os.Getenv("DEBUG_SQL") == "true",
	}
}
func (c Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort)
}
