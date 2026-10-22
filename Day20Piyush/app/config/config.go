// Configuration — reads from environment variables.
// Task 12: Connect Go application to RDS (via DATABASE_URL)
// Task 9/10: S3 config for uploads

package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
	UseDB       bool
	AWSRegion   string
	S3Bucket    string
}

func Load() Config {
	dbURL := getEnv("DATABASE_URL", "")
	return Config{
		Port:        getEnv("PORT", "8080"),
		DatabaseURL: dbURL,
		UseDB:       dbURL != "",
		AWSRegion:   getEnv("AWS_REGION", "ap-south-1"),
		S3Bucket:    getEnv("S3_BUCKET", "day20-piyush-employee-reports"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
