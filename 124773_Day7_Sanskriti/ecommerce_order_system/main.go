package main

import (
    "fmt"
    "log"

    "ecommerce_order_system/config"
    "ecommerce_order_system/database"
    "ecommerce_order_system/view"

    "github.com/joho/godotenv"
)

func main() {

    err := godotenv.Load()

    if err != nil {
        fmt.Println(".env file not found. Using system environment variables.")
    }

    cfg := config.LoadConfig()

    conn, err := database.ConnectDB(cfg)

    if err != nil {
        log.Fatal("Database connection failed:", err)
    }

    defer conn.Close(nil)

    fmt.Println("Database connected successfully.")

    view.Start(conn)
}
