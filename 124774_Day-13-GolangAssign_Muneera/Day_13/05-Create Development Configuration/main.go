package main

import (
	"fmt"
	"log"

	"employee-api/config"
	"employee-api/database"
)

func main() {
	// Step 1: Load configuration.
	cfg := config.Load()

	// Step 2: Validate configuration before connecting.
	if err := cfg.Validate(); err != nil {
		log.Fatal("Configuration validation failed: ", err)
	}

	fmt.Println("Configuration validated successfully!")

	// Step 3: Connect to PostgreSQL.
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	// Step 4: Ensure the connection pool is closed on exit.
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL database: ", err)
	}
	defer sqlDB.Close()

	// Step 5: Print application configuration.
	fmt.Println("\n========== APP CONFIG ==========")
	fmt.Println("App Environment :", cfg.AppEnv)
	fmt.Println("App Port        :", cfg.AppPort)

	// Step 6: Print database configuration without secrets.
	fmt.Println("\n========== DATABASE CONFIG ==========")
	fmt.Println("DB Host         :", cfg.DBHost)
	fmt.Println("DB Port         :", cfg.DBPort)
	fmt.Println("DB User         :", cfg.DBUser)
	fmt.Println("DB Name         :", cfg.DBName)
	fmt.Println("DB SSL Mode     :", cfg.DBSSLMode)

	// Step 7: Report successful startup.
	fmt.Println("\nPostgreSQL connected successfully!")
	fmt.Println("Development configuration loaded successfully!")
}
