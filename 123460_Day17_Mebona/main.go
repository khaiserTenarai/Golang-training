package main

import (
	"employee-management/auth"
	"employee-management/config"
	"employee-management/controller"
	gqlserver "employee-management/graphql"
	"employee-management/middleware"
	"employee-management/model"
	"employee-management/repository"
	"employee-management/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	cfg := config.Load()

	// 1. Database
	// TranslateError makes GORM return gorm.ErrDuplicatedKey for unique-constraint errors.
	gormCfg := &gorm.Config{TranslateError: true}
	if cfg.DebugSQL {
		gormCfg.Logger = logger.Default.LogMode(logger.Info) // print every SQL query
	}
	db, err := gorm.Open(postgres.Open(cfg.DSN()), gormCfg)
	if err != nil {
		log.Fatal("database connection failed:", err)
	}
	if err := db.AutoMigrate(&model.User{}, &model.Department{}, &model.Employee{}); err != nil {
		log.Fatal("migration failed:", err)
	}

	// 2. Build the layers from the bottom up (dependency injection)
	tokens := auth.NewTokenService(cfg.JWTSecret, cfg.JWTExpiry)

	userRepo := repository.NewUserRepository(db)
	deptRepo := repository.NewDepartmentRepository(db)
	empRepo := repository.NewEmployeeRepository(db)

	authSvc := service.NewAuthService(userRepo, tokens)
	deptSvc := service.NewDepartmentService(deptRepo)
	empSvc := service.NewEmployeeService(empRepo, deptRepo)

	if err := authSvc.SeedDefaultUsers(); err != nil {
		log.Fatal("could not create default users:", err)
	}

	authCtrl := controller.NewAuthController(authSvc)
	deptCtrl := controller.NewDepartmentController(deptSvc)
	empCtrl := controller.NewEmployeeController(empSvc)

	// the SAME services are used by REST (controllers) and GraphQL (resolvers)
	gqlHandler, err := gqlserver.NewHTTPHandler(empSvc, deptSvc, authSvc)
	if err != nil {
		log.Fatal("GraphQL schema failed:", err)
	}

	// 3. Routes
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logging())

	// public
	r.POST("/login", authCtrl.Login)

	// REST: must be logged in (any role)
	api := r.Group("", middleware.Auth(tokens))
	api.GET("/employees", empCtrl.List)
	api.GET("/employees/:id", empCtrl.GetByID)
	api.GET("/departments", deptCtrl.List)

	// REST: must be logged in AND be an ADMIN
	admin := api.Group("", middleware.RequireRole(auth.RoleAdmin))
	admin.POST("/employees", empCtrl.Create)
	admin.PUT("/employees/:id", empCtrl.Update)
	admin.DELETE("/employees/:id", empCtrl.Delete)
	admin.POST("/departments", deptCtrl.Create)

	// GraphQL: the token is read here, each resolver decides who may do what
	optional := middleware.OptionalAuth(tokens)
	r.Any("/graphql", optional, gin.WrapH(gqlHandler))
	r.Any("/playground", optional, gin.WrapH(gqlHandler))

	log.Println("REST:    http://localhost:" + cfg.ServerPort + "/employees")
	log.Println("GraphQL: http://localhost:" + cfg.ServerPort + "/graphql")
	log.Println("UI:      http://localhost:" + cfg.ServerPort + "/playground")
	log.Println("Demo logins -> admin / admin123 (ADMIN)   user / user123 (USER)")
	log.Fatal(r.Run(":" + cfg.ServerPort))
}
