package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


type EmployeeRequest struct {
	Name   string  `json:"name" binding:"required,min=2"`
	Email  string  `json:"email" binding:"required,email"`
	Age    int     `json:"age" binding:"required,gte=18,lte=65"`
	Salary float64 `json:"salary" binding:"required,gt=0"`
}

func NewRouter(auth *Auth, store *Store) *gin.Engine {
	router := gin.New()
	router.Use(requestID(), logger(), recovery())

	router.POST("/login", func(c *gin.Context) {
		var input LoginRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "send username and password as JSON"})
			return
		}

		token, role, err := auth.Login(input.Username, input.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"token": token, "role": role})
	})

	api := router.Group("/api")
	api.Use(auth.Required())


	api.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"username": c.GetString("username"), "role": c.GetString("role")})
	})


	api.GET("/employees", func(c *gin.Context) {
		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit < 1 {
			limit = 10
		}
		if limit > 100 {
			limit = 100
		}

		employees, total := store.List(page, limit)
		c.JSON(http.StatusOK, gin.H{
			"page":        page,
			"limit":       limit,
			"total":       total,
			"total_pages": (total + limit - 1) / limit,
			"data":        employees,
		})
	})

	api.POST("/employees", auth.RoleRequired("admin"), func(c *gin.Context) {
		var input EmployeeRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		saved := store.Add(Employee{
			Name:   input.Name,
			Email:  input.Email,
			Age:    input.Age,
			Salary: input.Salary,
		})
		c.JSON(http.StatusCreated, saved)
	})

	api.DELETE("/employees/:id", auth.RoleRequired("admin"), func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
			return
		}

		if err := store.Delete(id); err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "employee deleted"})
	})

	return router
}
