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
	gatewaycontroller "order-management/services/gateway/controller"
	gatewayroutes "order-management/services/gateway/routes"
	"order-management/shared/discovery"
	"order-management/shared/logger"
	"order-management/shared/middleware"
	"order-management/shared/observability"
)

func main() {
	config.Load()
	logger.Init()
	port, err := strconv.Atoi(config.Get("API_GATEWAY_PORT", "8080"))
	if err != nil || port < 1 || port > 65535 {
		log.Fatal("Invalid API_GATEWAY_PORT")
	}
	gatewayController, err := gatewaycontroller.NewGatewayController()
	if err != nil {
		log.Fatalf("Create gateway controller: %v", err)
	}
	router := gin.New()
	router.Use(middleware.RequestIDMiddleware(), observability.Middleware(), gin.Recovery())
	observability.Register(router, "api-gateway", nil)
	gatewayroutes.Register(router, gatewayController)
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: router, ReadHeaderTimeout: 5 * time.Second}
	consulClient, consulErr := discovery.NewConsulClient()
	serviceName := config.Get("API_GATEWAY_NAME", "api-gateway")
	serviceID := serviceName + "-" + strconv.Itoa(port)
	if consulErr != nil {
		logger.Get().Warn("Consul unavailable; Gateway remains available with fallback URLs", "error", consulErr)
	} else if err := discovery.Register(consulClient, serviceName, serviceID, config.Get("SERVICE_ADVERTISE_ADDRESS", "127.0.0.1"), port, "/health/ready", config.Get("HEALTH_CHECK_INTERVAL", "10s"), config.Get("HEALTH_CHECK_TIMEOUT", "3s")); err != nil {
		logger.Get().Warn("Gateway Consul registration failed", "error", err)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		logger.Get().Info("API Gateway starting", "port", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Get().Error("API Gateway failed", "error", err)
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
