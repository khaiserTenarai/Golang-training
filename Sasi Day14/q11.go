package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func main() {
	r := gin.Default()

	r.GET("/item/map", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success via gin.H",
		})
	})

	r.GET("/item/struct", func(c *gin.Context) {
		item := Item{ID: 1, Name: "Monitor", Price: 299.99}
		c.JSON(http.StatusOK, item)
	})

	r.Run(":8080")
}
