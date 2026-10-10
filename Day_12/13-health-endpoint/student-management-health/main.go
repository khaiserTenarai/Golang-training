package main

import (
 "log"
 "net/http"
 "time"

 "student-management/config"
 "student-management/controller"
 "student-management/database"
 "student-management/repository"
 "student-management/service"

 "github.com/gin-gonic/gin"
)

func main() {
 cfg, err := config.Load()
 if err != nil { log.Fatal(err) }

 db, err := database.Connect(cfg.DatabaseURL)
 if err != nil { log.Fatalf("database connection failed: %v", err) }
 defer db.Close()

 repo := repository.NewPostgresStudentRepository(db)
 studentService := service.NewStudentService(repo)
 studentController := controller.NewStudentController(studentService)

 router := gin.New()
 router.Use(gin.Recovery())
 router.Use(loggingMiddleware())

 // Health endpoint confirms the API process is responding.
 router.GET("/health", func(ctx *gin.Context) {
  ctx.JSON(http.StatusOK, gin.H{"status": "UP", "message": "Student Management API is running"})
 })

 router.POST("/students", studentController.Create)
 router.GET("/students", studentController.List)

 server := &http.Server{Addr: ":" + cfg.ServerPort, Handler: router, ReadHeaderTimeout: 5 * time.Second}
 log.Printf("server starting on port %s", cfg.ServerPort)
 log.Fatal(server.ListenAndServe())
}

func loggingMiddleware() gin.HandlerFunc {
 return func(ctx *gin.Context) {
  start := time.Now()
  ctx.Next()
  log.Printf("method=%s path=%s status=%d duration=%s", ctx.Request.Method, ctx.Request.URL.Path, ctx.Writer.Status(), time.Since(start))
 }
}
