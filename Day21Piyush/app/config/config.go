// Task 10: Externalize AWS Configuration
// All AWS settings come from environment variables — nothing hardcoded.

package config

import "os"

type Config struct {
	Port           string
	DatabaseURL    string
	UseDB          bool
	AWSRegion      string
	S3Bucket       string
	SecretName     string
	CloudWatchGroup string
	Environment    string
}

func Load() Config {
	dbURL := getEnv("DATABASE_URL", "")
	return Config{
		Port:            getEnv("PORT", "8080"),
		DatabaseURL:     dbURL,
		UseDB:           dbURL != "",
		AWSRegion:       getEnv("AWS_REGION", "ap-south-1"),
		S3Bucket:        getEnv("S3_BUCKET", "day21-piyush-employee-data"),
		SecretName:      getEnv("SECRET_NAME", "day21/db-credentials"),
		CloudWatchGroup: getEnv("CW_LOG_GROUP", "/employee-api/app"),
		Environment:     getEnv("APP_ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
