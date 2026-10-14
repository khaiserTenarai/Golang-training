package main

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	initJWT()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create an admin account from environment variables.
	adminEmail := strings.ToLower(
		strings.TrimSpace(os.Getenv("ADMIN_EMAIL")),
	)
	adminPassword := os.Getenv("ADMIN_PASSWORD")

	if adminEmail != "" && adminPassword != "" {
		if !validEmail(adminEmail) || !validPassword(adminPassword) {
			log.Fatal("invalid admin email or password configuration")
		}

		hash, err := hashPassword(adminPassword)
		if err != nil {
			log.Fatal("could not hash admin password")
		}

		if _, exists := users[adminEmail]; !exists {
			users[adminEmail] = User{
				ID:       adminEmail,
				Email:    adminEmail,
				Password: hash,
				Role:     "admin",
			}
		}
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := SetupRouter()

	log.Printf("Server starting on port %s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
