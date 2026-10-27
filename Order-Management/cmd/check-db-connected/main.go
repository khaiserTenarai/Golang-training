package main

import (
	"fmt"
	"log"

	"order-management/config"
	"order-management/database"
)

// main loads configuration and verifies the PostgreSQL connection.
func main() {
	config.Load()

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			log.Printf("Database close error: %v", err)
		}
	}()

	fmt.Println("PostgreSQL connected successfully!")
}
