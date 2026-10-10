package main

import (
	"log"
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
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer db.Close()

	// Inject repository into service, then service into controller.
	repo := repository.NewPostgresStudentRepository(db)
	studentService := service.NewStudentService(repo)
	studentController := controller.NewStudentController(studentService)

	router := gin.Default()
	router.POST("/students", studentController.CreateStudent)
	router.GET("/students", studentController.GetStudents)
	log.Printf("Server running at http://localhost:%s", cfg.Port)
	log.Fatal(router.Run(":" + cfg.Port))
}
