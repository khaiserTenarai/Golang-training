package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello Gin!")
	})

	r.Run(":8080")
}

/*

request url and output
http://localhost:8080/
Hello Gin!

*/
