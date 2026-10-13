package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"gin-secure-api-day15/internal/api"
	"gin-secure-api-day15/internal/config"
	"gin-secure-api-day15/internal/store"
)

func main() {
	cfg := config.Load()
	if cfg.JWTSecret == "" {
		log.Fatal("JWT_SECRET must be configured")
	}

	if cfg.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	repo := store.NewMemoryStore()
	router := api.NewRouter(cfg, repo)

	serverAddr := cfg.Address
	if serverAddr == "" {
		serverAddr = ":8080"
	}

	log.Printf("server listening on %s", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		log.Printf("server stopped: %v", err)
		os.Exit(1)
	}
}

