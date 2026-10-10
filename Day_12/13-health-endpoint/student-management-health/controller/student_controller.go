package controller

import (
 "net/http"
 "strconv"

 "student-management/model"
 "student-management/service"
 "student-management/utility"

 "github.com/gin-gonic/gin"
)

type StudentController struct { service service.StudentService }

func NewStudentController(s service.StudentService) *StudentController {
 return &StudentController{service: s}
}

func (c *StudentController) Create(ctx *gin.Context) {
 var student model.Student
 if err := ctx.ShouldBindJSON(&student); err != nil {
  ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON body"})
  return
 }
 if err := utility.ValidateStudent(student); err != nil {
  ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  return
 }
 if err := c.service.CreateStudent(ctx.Request.Context(), student); err != nil {
  ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create student"})
  return
 }
 ctx.JSON(http.StatusCreated, gin.H{"message": "student created"})
}

func (c *StudentController) List(ctx *gin.Context) {
 page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
 if err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error": "page must be a number"}); return }
 limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
 if err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error": "limit must be a number"}); return }

 students, total, err := c.service.ListStudents(ctx.Request.Context(), ctx.Query("name"), ctx.Query("grade"), ctx.DefaultQuery("sortBy", "id"), ctx.DefaultQuery("order", "asc"), page, limit)
 if err != nil { ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
 ctx.JSON(http.StatusOK, gin.H{"data": students, "total": total, "page": page, "limit": limit, "sortBy": ctx.DefaultQuery("sortBy", "id"), "order": ctx.DefaultQuery("order", "asc")})
}
