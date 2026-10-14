/*
Question 12: Handle form parameters.

Open http://localhost:8080/ in a browser and submit the form, or use curl:
curl -X POST localhost:8080/register -d "name=Asha" -d "email=asha@example.com"
curl -X POST localhost:8080/register-struct -d "name=Ravi" -d "email=ravi@example.com"
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

const formPage = `<!DOCTYPE html>
<html>
<body>
  <h2>Register an employee</h2>
  <form action="/register" method="post">
    <p>Name: <input name="name"></p>
    <p>Email: <input name="email"></p>
    <p>Department: <input name="department"></p>
    <button type="submit">Register</button>
  </form>
</body>
</html>`

type Registration struct {
	Name       string `form:"name" binding:"required"`
	Email      string `form:"email" binding:"required,email"`
	Department string `form:"department"`
}

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

	router.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(formPage))
	})

	router.POST("/register", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		department := c.DefaultPostForm("department", "General")

		if name == "" || email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "name and email are required"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"name": name, "email": email, "department": department})
	})

	router.POST("/register-struct", func(c *gin.Context) {
		var reg Registration
		if err := c.ShouldBind(&reg); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, reg)
	})

	router.Run(":" + readPort())
}
