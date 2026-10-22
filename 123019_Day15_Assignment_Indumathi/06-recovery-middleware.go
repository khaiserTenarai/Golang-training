package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var reader = bufio.NewReader(os.Stdin)

func ask(question string) string {
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func myRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if problem := recover(); problem != nil {
				fmt.Println("Recovered from panic:", problem)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "something went wrong on the server"})
			}
		}()
		c.Next()
	}
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(myRecovery())


	router.GET("/divide", func(c *gin.Context) {
		a, errA := strconv.Atoi(c.Query("a"))
		b, errB := strconv.Atoi(c.Query("b"))
		if errA != nil || errB != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "send two numbers like /divide?a=10&b=2"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"result": a / b})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
