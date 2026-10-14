package main

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name" binding:"required"`
	Salary float64 `json:"salary" binding:"required,gt=0"`
}

func NewRouter() *gin.Engine {
	var mu sync.Mutex
	employees := map[int]Employee{}
	nextID := 1

	router := gin.New()

	router.POST("/employees", func(c *gin.Context) {
		var input Employee
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		mu.Lock()
		input.ID = nextID
		nextID++
		employees[input.ID] = input
		mu.Unlock()

		c.JSON(http.StatusCreated, input)
	})

	router.GET("/employees/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id must be a number"})
			return
		}

		mu.Lock()
		emp, found := employees[id]
		mu.Unlock()

		if !found {
			c.JSON(http.StatusNotFound, gin.H{"error": "employee not found"})
			return
		}
		c.JSON(http.StatusOK, emp)
	})

	return router
}
