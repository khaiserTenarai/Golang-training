package main

import (
	"log"
	"net/http"

	"student-management/config"
	"student-management/controller"
	"student-management/database"
	"student-management/repository"
	"student-management/service"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inject each layer into the next layer.
	studentRepo := repository.NewPostgresRepository(db)
	studentService := service.NewStudentService(studentRepo)
	studentController := controller.NewStudentController(studentService)

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(requestLogging())

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	router.GET("/students", studentController.GetStudents)
	router.POST("/students", studentController.CreateStudent)

	log.Printf("Server listening on :%s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal(err)
	}
}

// requestLogging logs each request.
func requestLogging() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("Started %s %s", ctx.Request.Method, ctx.Request.URL.Path)
		ctx.Next()
		log.Printf("Finished %s %s status=%d", ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status())
	}
}
