// @title Product Service API
// @version 1.0
// @description REST API for managing products.
// @host localhost:8081
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
	"order-management/config"
	"order-management/database"
	"order-management/services/product/controller"
	"order-management/services/product/model"
	"order-management/services/product/repository"
	"order-management/services/product/routes"
	"order-management/services/product/service"
	"order-management/shared/discovery"
	"order-management/shared/logger"
	"order-management/shared/middleware"
	"order-management/shared/observability"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "order-management/docs/product"

	"github.com/gin-gonic/gin"
)

// @title Product Service API
// @version 1.0
// @description REST API for managing products.
// @BasePath /
func main() {
	config.Load()
	logger.Init()

	port, err := strconv.Atoi(config.Get("PRODUCT_SERVICE_PORT", "8081"))
	if err != nil || port < 1 || port > 65535 {
		log.Fatal("Invalid PRODUCT_SERVICE_PORT")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer func() {
		if err := database.Close(db); err != nil {
			logger.Get().Error("Failed to close database", "error", err)
		}
	}()

	if err := db.AutoMigrate(&model.Product{}, &model.StockOperation{}); err != nil {
		log.Fatalf("Product database migration failed: %v", err)
	}

	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productController := controller.NewProductController(productService)

	router := gin.New()
	router.Use(middleware.RequestIDMiddleware(), observability.Middleware(), gin.Recovery())

	observability.Register(router, "product-service", func() error {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		return sqlDB.PingContext(ctx)
	})
	if err := routes.Register(router, productController, productService); err != nil {
		log.Fatalf("Product routes registration failed: %v", err)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server := &http.Server{
		Addr:              ":" + strconv.Itoa(port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	consulClient, consulErr := discovery.NewConsulClient()
	if consulErr != nil {
		logger.Get().Warn("Consul unavailable; Product Service will continue without registration", "error", consulErr)
		consulClient = nil
	}

	serviceName := config.Get("PRODUCT_SERVICE_NAME", "product-service")
	serviceID := serviceName + "-" + strconv.Itoa(port)
	serviceAddress := config.Get("SERVICE_ADVERTISE_ADDRESS", config.Get("PRODUCT_SERVICE_ADDRESS", "127.0.0.1"))

	interval := config.Get("HEALTH_CHECK_INTERVAL", "10s")
	timeout := config.Get("HEALTH_CHECK_TIMEOUT", "5s")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		logger.Get().Info(
			"Starting Product Service",
			"address", server.Addr,
		)

		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			logger.Get().Error("Product Service failed", "error", err)
			stop <- syscall.SIGTERM
		}
	}()

	if consulClient != nil {
		if err := discovery.Register(consulClient, serviceName, serviceID, serviceAddress, port, "/health/ready", interval, timeout); err != nil {
			logger.Get().Warn("Consul registration failed; service will continue", "error", err)
		} else {
			logger.Get().Info("Registered with Consul", "service", serviceName)
		}
	}

	<-stop

	ctx, cancel := signalContext()
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Get().Error("Graceful shutdown failed", "error", err)
	}

	if consulClient != nil {
		if err := discovery.Deregister(consulClient, serviceID); err != nil {
			logger.Get().Warn("Consul deregistration failed", "error", err)
		}
	}

	logger.Get().Info("Product Service stopped")
}

// signalContext returns a context with a shutdown timeout.
func signalContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}
