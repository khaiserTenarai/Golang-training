package routes

import (
	"student-management/controller"

	"github.com/gin-gonic/gin"
)

// StudentRoutes registers student endpoints.
func StudentRoutes(router *gin.Engine) {
	student := router.Group("/api/students")

	student.GET("", controller.GetStudents)
	student.GET("/:id", controller.GetStudent)
	student.POST("", controller.CreateStudent)
	student.POST("/form", controller.CreateStudentForm)
	student.PUT("/:id", controller.UpdateStudent)
	student.DELETE("/:id", controller.DeleteStudent)
}
