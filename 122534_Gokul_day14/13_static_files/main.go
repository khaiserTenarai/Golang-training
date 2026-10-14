/*
Question 13: Serve static files.

Run it from inside this folder (the static files are read from ./static):
    cd 13_static_files
    go run .
Then open http://localhost:8080/ in a browser.
*/

package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func readPort() string {
	fmt.Print("Enter the port (press Enter for 8080): ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	text = strings.TrimSpace(text)
	if text == "" {
		return "8080"
	}
	return text
}

func main() {
	router := gin.Default()

	router.Static("/static", "./static")
	router.StaticFile("/", "./static/index.html")

	router.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from the Gin API"})
	})

	router.Run(":" + readPort())
}
