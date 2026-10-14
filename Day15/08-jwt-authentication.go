package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var reader = bufio.NewReader(os.Stdin)

func ask(question string) string {
	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	secret := ask("Enter a secret key used to sign tokens: ")
	username := ask("Create a username: ")
	password := ask("Create a password: ")
	if secret == "" || username == "" || password == "" {
		fmt.Println("Secret, username and password are all needed.")
		return
	}


	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Could not hash the password:", err)
		return
	}

	router := gin.New()

	router.POST("/login", func(c *gin.Context) {
		var input LoginRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "send username and password as JSON"})
			return
		}

		wrongPassword := bcrypt.CompareHashAndPassword(hash, []byte(input.Password)) != nil
		if input.Username != username || wrongPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": input.Username,
			"exp":      time.Now().Add(time.Hour).Unix(),
		})
		signed, err := token.SignedString([]byte(secret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": signed})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
