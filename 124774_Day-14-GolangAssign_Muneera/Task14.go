package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.HandleMethodNotAllowed = true

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "route not found",
		})
	})

	r.NoMethod(func(c *gin.Context) {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "method not allowed",
		})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Gin!",
		})
	})

	r.Run(":8080")
}

//Use NoRoute() for unknown routes and NoMethod() for unsupported HTTP methods.
// 14. Custom 404 and 405 Handling

// Request URL: http://localhost:8080/unknown
// Method: GET
// Output:
// {
//   "error": "route not found"
// }
// Status: 404 Not Found

// Request URL: http://localhost:8080/hello
// Method: GET
// Output:
// {
//   "message": "Hello Gin!"
// }
// Status: 200 OK

// Request URL: http://localhost:8080/hello
// Method: POST
// Output:
// {
//   "error": "method not allowed"
// }
// Status: 405 Method Not Allowed
