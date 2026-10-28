package main

import (
	"log"

	"employee-api/config"
	"employee-api/database"
)

func main() {
	// Step 1: Load configuration.
	cfg := config.Load()
	log.Println("Application environment:", cfg.AppEnv)

	// Step 2: Connect to PostgreSQL.
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Step 3: Close the database pool when the application exits.
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Get database connection:", err)
	}
	defer sqlDB.Close()

	log.Println("PostgreSQL connected successfully!")
}
