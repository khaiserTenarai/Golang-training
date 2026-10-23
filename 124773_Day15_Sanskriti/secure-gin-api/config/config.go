package config

import (
 "fmt"
 "os"
 "strconv"

 "github.com/joho/godotenv"
)

type Config struct { DBHost, DBPort, DBName, DBUser, DBPassword, ServerPort, JWTSecret string; JWTExpiryMinutes int }

func Load() (Config, error) {
 _ = godotenv.Load()
 expiry, err := strconv.Atoi(getEnv("JWT_EXPIRY_MINUTES", "60")); if err != nil || expiry < 1 { return Config{}, fmt.Errorf("JWT_EXPIRY_MINUTES must be a positive integer") }
 c := Config{DBHost:getEnv("DB_HOST","localhost"), DBPort:getEnv("DB_PORT","5432"), DBName:getEnv("DB_NAME","secure_employee_db"), DBUser:getEnv("DB_USER","postgres"), DBPassword:os.Getenv("DB_PASSWORD"), ServerPort:getEnv("SERVER_PORT","8080"), JWTSecret:os.Getenv("JWT_SECRET"), JWTExpiryMinutes:expiry}
 if len(c.JWTSecret) < 32 { return Config{}, fmt.Errorf("JWT_SECRET must be at least 32 characters") }
 return c,nil
}
func getEnv(k, fallback string) string { if v:=os.Getenv(k); v!="" { return v }; return fallback }
