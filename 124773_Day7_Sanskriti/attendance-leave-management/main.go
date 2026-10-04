package main

import (
	"context"
	"fmt"
	"log"

	"attendance_leave_management/config"
	"attendance_leave_management/database"
	"attendance_leave_management/view"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found. Using system environment variables.")
	}

	cfg := config.LoadConfig()

	conn, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("Database connected successfully.")
	view.Start(conn)
}
