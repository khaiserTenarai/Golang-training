package main

import (
	"fmt"

	"employee-api/config"
)

func main() {
	// Step 1: Load configuration.
	cfg := config.Load()

	// Step 2: Print application configuration.
	fmt.Println("========== APP CONFIG ==========")
	fmt.Println("App Environment :", cfg.AppEnv)
	fmt.Println("App Port        :", cfg.AppPort)

	// Step 3: Print database configuration except password.
	fmt.Println("\n========== DATABASE CONFIG ==========")
	fmt.Println("DB Host         :", cfg.DBHost)
	fmt.Println("DB Port         :", cfg.DBPort)
	fmt.Println("DB User         :", cfg.DBUser)
	fmt.Println("DB Name         :", cfg.DBName)
	fmt.Println("DB SSL Mode     :", cfg.DBSSLMode)

	// DBPassword is intentionally not printed.
}
