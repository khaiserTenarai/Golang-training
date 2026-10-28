package main

import (
	"fmt"
	"log"

	"employee-api/config"
	"employee-api/database"
)

func main() {
	// Step 1: Load the selected environment configuration.
	cfg := config.Load()

	fmt.Println("Environment:", cfg.AppEnv)

	// Step 2: Validate the configuration.
	if err := cfg.Validate(); err != nil {
		log.Fatal("Configuration validation failed: ", err)
	}

	// Step 3: Connect to the configured database.
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	// Step 4: Close the database pool on exit.
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL database: ", err)
	}
	defer sqlDB.Close()

	// Step 5: Print settings, excluding the password.
	fmt.Println("\n========== CONFIG ==========")
	fmt.Println("App Environment :", cfg.AppEnv)
	fmt.Println("App Port        :", cfg.AppPort)
	fmt.Println("DB Host         :", cfg.DBHost)
	fmt.Println("DB Port         :", cfg.DBPort)
	fmt.Println("DB User         :", cfg.DBUser)
	fmt.Println("DB Name         :", cfg.DBName)
	fmt.Println("DB SSL Mode     :", cfg.DBSSLMode)

	fmt.Println("\nPostgreSQL connected successfully!")
}
