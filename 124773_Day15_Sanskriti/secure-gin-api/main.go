package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"secure-gin-api/config"
	"secure-gin-api/controller"
	"secure-gin-api/database"
	"secure-gin-api/repository"
	"secure-gin-api/routes"
	"secure-gin-api/service"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := database.Connect(ctx, cfg)
	if err != nil {
		log.Fatalf("PostgreSQL connection failed: %v", err)
	}
	defer db.Close()
	repo := repository.NewEmployeeRepository(db)
	svc := service.NewEmployeeService(repo)
	ctrl := controller.NewEmployeeController(svc, cfg.JWTSecret, cfg.JWTExpiryMinutes)
	router := routes.Setup(ctrl, cfg.JWTSecret)
	srv := &http.Server{Addr: ":" + cfg.ServerPort, Handler: router, ReadHeaderTimeout: 5 * time.Second}
	go func() {
		log.Printf("server listening on :%s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed: %v", err)
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown error: %v", err)
	}
}
