// @title Order Service API
// @version 1.0
// @description REST and GraphQL API for managing orders.
// @host localhost:8082
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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

	"order-management/config"
	"order-management/database"
	"order-management/services/order/client"
	"order-management/services/order/controller"
	"order-management/services/order/model"
	"order-management/services/order/repository"
	"order-management/services/order/routes"
	"order-management/services/order/service"
	"order-management/shared/discovery"
	"order-management/shared/logger"
	"order-management/shared/middleware"
	"order-management/shared/observability"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "order-management/docs/order"
)

func main() {
	config.Load()
	logger.Init()
	port, err := strconv.Atoi(config.Get("ORDER_SERVICE_PORT", "8082"))
	if err != nil || port < 1 || port > 65535 {
		log.Fatal("Invalid ORDER_SERVICE_PORT")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			logger.Get().Error("Database close error", "error", err)
		}
	}()
	if err := db.AutoMigrate(&model.Order{}); err != nil {
		log.Fatalf("Order migration failed: %v", err)
	}

	orderRepo := repository.NewOrderRepository(db)
	productClient := client.NewProductClient()
	orderSvc := service.NewOrderService(orderRepo, productClient)
	orderCtrl := controller.NewOrderController(orderSvc)

	router := gin.New()
	router.Use(middleware.RequestIDMiddleware(), observability.Middleware(), gin.Recovery())
	observability.Register(router, "order-service", func() error {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		return sqlDB.PingContext(ctx)
	})
	routes.Register(router, orderCtrl, orderSvc)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: router, ReadHeaderTimeout: 5 * time.Second}

	consulClient, consulErr := discovery.NewConsulClient()
	serviceName := config.Get("ORDER_SERVICE_NAME", "order-service")
	serviceID := serviceName + "-" + strconv.Itoa(port)
	if consulErr != nil {
		logger.Get().Warn("Consul unavailable; Order Service will use configured Product Service fallback", "error", consulErr)
	} else if err := discovery.Register(consulClient, serviceName, serviceID, config.Get("SERVICE_ADVERTISE_ADDRESS", "127.0.0.1"), port, "/health/ready", config.Get("HEALTH_CHECK_INTERVAL", "10s"), config.Get("HEALTH_CHECK_TIMEOUT", "3s")); err != nil {
		logger.Get().Warn("Consul registration failed", "error", err)
	} else {
		logger.Get().Info("Registered with Consul", "service", serviceName, "port", port)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		logger.Get().Info("Order Service starting", "port", port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Get().Error("Order Service failed", "error", err)
			stop <- syscall.SIGTERM
		}
	}()
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Get().Error("Graceful shutdown failed", "error", err)
	}
	if consulClient != nil {
		if err := discovery.Deregister(consulClient, serviceID); err != nil {
			logger.Get().Warn("Consul deregistration failed", "error", err)
		}
	}
}
