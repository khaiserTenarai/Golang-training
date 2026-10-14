package main

import (
	"day15-secure-gin-api/controller"
	"day15-secure-gin-api/middleware"
	"day15-secure-gin-api/repository"
	"day15-secure-gin-api/service"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.New()
	r.Use(middleware.RequestID(), middleware.Recovery())
	r.Use(cors.New(cors.Config{AllowOrigins: []string{"http://localhost:3000"}, AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, AllowHeaders: []string{"Origin", "Content-Type", "Authorization"}}))
	r.POST("/login", controller.Login)
	repo := repository.NewEmployeeRepository()
	svc := service.NewEmployeeService(repo)
	ctrl := controller.NewEmployeeController(svc)
	p := r.Group("/api")
	p.Use(middleware.JWTAuth())
	{
		p.GET("/employees", ctrl.GetAll)
		p.GET("/employees/:id", ctrl.GetByID)
		p.POST("/employees", ctrl.Create)
		p.PUT("/employees/:id", ctrl.Update)
		p.DELETE("/employees/:id", middleware.RequireRole("admin"), ctrl.Delete)
	}
	r.NoRoute(func(c *gin.Context) { c.JSON(404, gin.H{"error": "Route not found"}) })
	r.NoMethod(func(c *gin.Context) { c.JSON(405, gin.H{"error": "Method not allowed"}) })
	return r
}
func main() {
	setupRouter().Run(":8080")
	_ = http.StatusOK
}
