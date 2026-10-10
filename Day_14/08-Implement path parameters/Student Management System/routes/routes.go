package routes

import (
	"student-management/controller"

	"github.com/gin-gonic/gin"
)

// StudentRoutes groups all student-related routes.
func StudentRoutes(router *gin.Engine) {
	// All routes share the /api/students prefix.
	student := router.Group("/api/students")

	student.GET("", controller.GetStudents)
	student.GET("/:id", controller.GetStudent)
	student.POST("", controller.CreateStudent)
	student.PUT("/:id", controller.UpdateStudent)
	student.DELETE("/:id", controller.DeleteStudent)
}
