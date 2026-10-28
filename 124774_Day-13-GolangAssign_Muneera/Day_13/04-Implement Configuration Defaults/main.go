package main

import (
	"fmt"
	"log"

	"employee-api/config"
	"employee-api/database"
)

func main() {
	// Step 1: Load configuration from .env or system environment.
	cfg := config.Load()

	// Step 2: Validate configuration.
	if err := cfg.Validate(); err != nil {
		log.Fatal("Configuration validation failed: ", err)
	}

	fmt.Println("Configuration validated successfully!")

	// Step 3: Connect to PostgreSQL using GORM.
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	// Step 4: Get the underlying SQL connection pool.
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL database: ", err)
	}

	// Step 5: Close the database pool when the application exits.
	defer sqlDB.Close()

	// Step 6: Print application configuration.
	fmt.Println("\n========== APP CONFIG ==========")
	fmt.Println("App Environment :", cfg.AppEnv)
	fmt.Println("App Port         :", cfg.AppPort)

	// Step 7: Print database configuration (never print password).
	fmt.Println("\n========== DATABASE CONFIG ==========")
	fmt.Println("DB Host          :", cfg.DBHost)
	fmt.Println("DB Port          :", cfg.DBPort)
	fmt.Println("DB User          :", cfg.DBUser)
	fmt.Println("DB Name          :", cfg.DBName)
	fmt.Println("DB SSL Mode      :", cfg.DBSSLMode)

	// Step 8: Confirm successful connection.
	fmt.Println("\nPostgreSQL connected successfully!")
}
