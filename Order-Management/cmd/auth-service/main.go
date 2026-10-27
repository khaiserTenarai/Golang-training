package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"order-management/config"
	authcontroller "order-management/services/auth/controller"
	authroutes "order-management/services/auth/routes"
	authservice "order-management/services/auth/service"
	"order-management/shared/discovery"
	"order-management/shared/logger"
	"order-management/shared/middleware"
	"order-management/shared/observability"
)

func main() {
	config.Load()
	logger.Init()
	port, err := strconv.Atoi(config.Get("AUTH_SERVICE_PORT", "8084"))
	if err != nil || port < 1 || port > 65535 {
		log.Fatal("Invalid AUTH_SERVICE_PORT")
	}
	authSvc := authservice.NewAuthService()
	authCtrl := authcontroller.NewAuthController(authSvc)
	router := gin.New()
	router.Use(middleware.RequestIDMiddleware(), observability.Middleware(), gin.Recovery())
	observability.Register(router, "auth-service", nil)
	authroutes.Register(router, authCtrl)
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: router, ReadHeaderTimeout: 5 * time.Second}
	consulClient, consulErr := discovery.NewConsulClient()
	serviceName := config.Get("AUTH_SERVICE_NAME", "auth-service")
	serviceID := serviceName + "-" + strconv.Itoa(port)
	if consulErr != nil {
		logger.Get().Warn("Consul unavailable; Auth Service will use fallback URL", "error", consulErr)
	} else if err := discovery.Register(consulClient, serviceName, serviceID, config.Get("SERVICE_ADVERTISE_ADDRESS", "127.0.0.1"), port, "/health/ready", config.Get("HEALTH_CHECK_INTERVAL", "10s"), config.Get("HEALTH_CHECK_TIMEOUT", "3s")); err != nil {
		logger.Get().Warn("Consul registration failed", "error", err)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		logger.Get().Info("Auth Service starting", "port", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Get().Error("Auth Service failed", "error", err)
			stop <- syscall.SIGTERM
		}
	}()
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_ = server.Shutdown(ctx)
	if consulClient != nil {
		_ = discovery.Deregister(consulClient, serviceID)
	}
}
