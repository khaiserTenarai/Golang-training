package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"secure-banking-api/config"
	"secure-banking-api/controller"
	"secure-banking-api/database"
	"secure-banking-api/middleware"
	"secure-banking-api/model"
	"secure-banking-api/repository"
	"secure-banking-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	if cfg.DBURL == "" || cfg.JWTSecret == "" {
		log.Fatal("DB_URL and JWT_SECRET must be set")
	}

	db, err := database.Connect(cfg.DBURL)
	if err != nil {
		log.Fatalf("database: %v", err)
	}
	defer db.Close()

	// repository -> service -> controller
	userRepo := repository.NewUserRepository(db)
	accRepo := repository.NewAccountRepository(db)
	authCtl := controller.NewAuthController(service.NewAuthService(userRepo, cfg))
	accCtl := controller.NewAccountController(service.NewAccountService(accRepo), cfg.UploadDir)

	r := gin.New()
	r.MaxMultipartMemory = 8 << 20
	// Global middleware order: RequestID -> Logger -> Recovery -> CORS
	r.Use(middleware.RequestID(), middleware.Logger(), middleware.Recovery(), middleware.CORS(cfg.AllowedOrigins))

	v1 := r.Group("/api/v1")

	// Demo: watch the logs to see A/B before and after order.
	v1.GET("/ping", middleware.Trace("A"), middleware.Trace("B"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Public
	v1.POST("/auth/register", authCtl.Register)
	v1.POST("/auth/login", authCtl.Login)

	// Any logged-in user
	accounts := v1.Group("/accounts", middleware.JWTAuth(cfg.JWTSecret))
	{
		accounts.POST("", accCtl.Create)
		accounts.GET("", accCtl.ListMine)
		accounts.GET("/:id", accCtl.Get)
		accounts.POST("/:id/document", accCtl.UploadDocument)
	}

	// Admin only (JWT first, then role check)
	admin := v1.Group("/admin", middleware.JWTAuth(cfg.JWTSecret), middleware.RequireRole(model.RoleAdmin))
	{
		admin.GET("/accounts", accCtl.ListAll)
	}

	srv := &http.Server{Addr: ":" + cfg.Port, Handler: r, ReadHeaderTimeout: 5 * time.Second}

	go func() {
		log.Printf("listening on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("forced shutdown: %v", err)
	}
}
