package main

import (
	"fmt"
	"log"

	"employee-api/config"
	"employee-api/database"
)

func main() {
	// Step 1: Load normal configuration.
	cfg := config.Load()

	// Step 2: Validate normal configuration.
	if err := cfg.Validate(); err != nil {
		log.Fatal("Configuration validation failed: ", err)
	}

	// Step 3: Load secrets separately.
	secrets, err := config.LoadSecrets()
	if err != nil {
		log.Fatal("Secret loading failed: ", err)
	}

	// Step 4: Connect to PostgreSQL.
	db, err := database.Connect(cfg, secrets)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	// Step 5: Close the connection pool when main exits.
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL database: ", err)
	}
	defer sqlDB.Close()

	// Step 6: Print safe, non-secret configuration details.
	fmt.Println("========== CONFIG ==========")
	fmt.Println("Environment :", cfg.AppEnv)
	fmt.Println("App Port    :", cfg.AppPort)
	fmt.Println("DB Host     :", cfg.DBHost)
	fmt.Println("DB Port     :", cfg.DBPort)
	fmt.Println("DB User     :", cfg.DBUser)
	fmt.Println("DB Name     :", cfg.DBName)
	fmt.Println("DB SSL Mode :", cfg.DBSSLMode)

	fmt.Println("\nPostgreSQL connected successfully!")
}


/*

# ==========================================
# TASK 08: SEPARATE SECRETS FROM CONFIGURATION
# Project: employee-api
# ==========================================

# Step 1: Navigate to your project folder
cd path\to\employee-api

# Step 2: Select the development environment
$env:APP_ENV = "development"

# Step 3: Set your actual PostgreSQL password
# Replace the placeholder with your local database password
$env:DB_PASSWORD = "your_local_postgres_password"

# Step 4: Download and tidy Go dependencies
go mod tidy

# Step 5: Compile the application and check for errors
go build .

# Step 6: Run the application
go run .

# Expected result:
# Environment : development
# App Port    : 8080
# DB Host     : localhost
# DB Port     : 5432
# DB User     : postgres
# DB Name     : employee_db
# DB SSL Mode : disable
# PostgreSQL connected successfully!

# ==========================================
# STEP 7: TEST SECRET VALIDATION
# ==========================================

# Remove the password from the current terminal session
Remove-Item Env:DB_PASSWORD -ErrorAction SilentlyContinue

# Run the application without the password
# Expected: DB_PASSWORD is required
go run .

# ==========================================
# STEP 8: RESTORE THE PASSWORD
# ==========================================

# Set your actual PostgreSQL password again
$env:DB_PASSWORD = "your_local_postgres_password"

# Run the application again
go run .

# ==========================================
# STEP 9: TEST CONFIGURATION
# ==========================================

# Switch to the test environment
$env:APP_ENV = "test"

# Run using the test database configuration
# Ensure employee_test_db exists in PostgreSQL
go run .

# ==========================================
# TASK 08 COMPLETE
# ==========================================

*/