package controller

import (
	"day15-secure-gin-api/model"
	"day15-secure-gin-api/service"
	"day15-secure-gin-api/utility"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var x LoginRequest
	if c.ShouldBindJSON(&x) != nil {
		c.JSON(400, gin.H{"error": "Username and password are required"})
		return
	}
	role := ""
	if x.Username == "admin" && x.Password == "admin123" {
		role = "admin"
	} else if x.Username == "user" && x.Password == "user123" {
		role = "user"
	} else {
		c.JSON(401, gin.H{"error": "Invalid username or password"})
		return
	}
	t, e := utility.GenerateToken(x.Username, role)
	if e != nil {
		c.JSON(500, gin.H{"error": "Token generation failed"})
		return
	}
	c.JSON(200, gin.H{"token": t, "role": role})
}

type EmployeeController struct {
	s service.EmployeeService
}

func NewEmployeeController(s service.EmployeeService) *EmployeeController {
	return &EmployeeController{s}
}
func (c *EmployeeController) GetAll(x *gin.Context) {
	p, _ := strconv.Atoi(x.DefaultQuery("page", "1"))
	l, _ := strconv.Atoi(x.DefaultQuery("limit", "2"))
	x.JSON(200, gin.H{"page": p, "limit": l, "data": c.s.GetAll(p, l)})
}
func (c *EmployeeController) GetByID(x *gin.Context) {
	id, e := strconv.Atoi(x.Param("id"))
	if e != nil {
		x.JSON(400, gin.H{"error": "Invalid employee ID"})
		return
	}
	v, ok := c.s.GetByID(id)
	if !ok {
		x.JSON(404, gin.H{"error": "Employee not found"})
		return
	}
	x.JSON(200, v)
}
func (c *EmployeeController) Create(x *gin.Context) {
	var e model.Employee
	if x.ShouldBindJSON(&e) != nil {
		x.JSON(400, gin.H{"error": "Validation failed"})
		return
	}
	x.JSON(201, c.s.Create(e))
}
func (c *EmployeeController) Update(x *gin.Context) {
	id, e := strconv.Atoi(x.Param("id"))
	if e != nil {
		x.JSON(400, gin.H{"error": "Invalid employee ID"})
		return
	}
	var v model.Employee
	if x.ShouldBindJSON(&v) != nil {
		x.JSON(400, gin.H{"error": "Validation failed"})
		return
	}
	v, ok := c.s.Update(id, v)
	if !ok {
		x.JSON(404, gin.H{"error": "Employee not found"})
		return
	}
	x.JSON(200, v)
}
func (c *EmployeeController) Delete(x *gin.Context) {
	id, e := strconv.Atoi(x.Param("id"))
	if e != nil {
		x.JSON(400, gin.H{"error": "Invalid employee ID"})
		return
	}
	if !c.s.Delete(id) {
		x.JSON(404, gin.H{"error": "Employee not found"})
		return
	}
	x.JSON(200, gin.H{"message": "Employee deleted successfully"})
}
