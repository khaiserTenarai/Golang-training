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

// checks the "Authorization: Bearer <token>" header before letting the
// request reach a protected route
func authRequired(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		tokenText := strings.TrimPrefix(header, "Bearer ")
		token, err := jwt.Parse(tokenText, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		}, jwt.WithValidMethods([]string{"HS256"}))
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("username", claims["username"])
		c.Next()
	}
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

	router.GET("/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "anyone can open this"})
	})

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
		signed, _ := token.SignedString([]byte(secret))
		c.JSON(http.StatusOK, gin.H{"token": signed})
	})

	protected := router.Group("/api")
	protected.Use(authRequired(secret))
	protected.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "welcome back", "username": c.GetString("username")})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
