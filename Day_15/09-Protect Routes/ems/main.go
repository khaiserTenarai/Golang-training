package main

import (
	"log"

	"ems/auth"
	"ems/config"
	"ems/controller"
	"ems/database"
	"ems/middleware"
	"ems/model"
	"ems/repository"
	"ems/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Load configuration.
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// 2. Connect to PostgreSQL.
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	// 3. Create the employee table if necessary.
	if err := db.AutoMigrate(&model.Employee{}); err != nil {
		log.Fatal("Migration failed: ", err)
	}

	// 4. Initialize JWT authentication.
	authService, err := auth.New(
		cfg.JWTSecret,
		cfg.JWTExpiryMinutes,
		cfg.AdminUsername,
		cfg.AdminPassword,
		cfg.UserUsername,
		cfg.UserPassword,
	)
	if err != nil {
		log.Fatal(err)
	}

	// 5. Initialize employee layers.
	repo := repository.NewEmployeeRepository(db)
	employeeService := service.NewEmployeeService(repo)
	employeeController := controller.NewEmployeeController(employeeService)
	authController := controller.NewAuthController(authService)

	// 6. Create the Gin router.
	router := gin.New()

	// Global middleware executes for every request.
	router.Use(middleware.RequestID())
	router.Use(middleware.Logging())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORSMiddleware())

	// Public routes.
	router.POST("/login", authController.Login)

	router.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(503, gin.H{"status": "database unavailable"})
			return
		}
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// Day 15: Next() demonstration.
	router.GET("/middleware-next", middleware.NextDemo(),
		func(c *gin.Context) {
			log.Println("Handler: Executing route handler")
			c.JSON(200, gin.H{
				"message": "Next() demonstration completed",
			})
		},
	)

	// Day 15: Abort() demonstration.
	router.GET("/middleware-abort", middleware.AbortDemo(),
		func(c *gin.Context) {
			log.Println("Handler: This should not execute")
			c.JSON(200, gin.H{"message": "Handler executed"})
		},
	)

	// Day 15: Recovery demonstration.
	router.GET("/panic", func(c *gin.Context) {
		panic("testing recovery middleware")
	})

	// All routes in this group require a valid JWT.
	protected := router.Group("/employees")
	protected.Use(middleware.RequireAuth(authService))

	// Admin and user can read employee data.
	protected.GET("", employeeController.GetAll)
	protected.GET("/:id", employeeController.GetByID)

	// Only admins can create, update and delete.
	admin := protected.Group("")
	admin.Use(middleware.RequireRole("admin"))

	admin.POST("", employeeController.Create)
	admin.PUT("/:id", employeeController.Update)
	admin.DELETE("/:id", employeeController.Delete)

	// 7. Start the server.
	log.Printf("Server running at http://localhost:%s", cfg.ServerPort)

	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}
