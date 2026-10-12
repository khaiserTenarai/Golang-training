package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.Run(":8080")
}

// Request URL: http://localhost:8080/login
// Method: POST

// Input (Form Data):
// username = Raj
// password = 12345

// Output:
// {
//   "username": "Raj",
//   "password": "12345"
// }
// Status: 200 OK
