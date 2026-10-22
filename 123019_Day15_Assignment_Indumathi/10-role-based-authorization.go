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

type User struct {
	Username string
	Hash     []byte
	Role     string
}

var users = map[string]User{}

func addUser(username, password, role string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	users[username] = User{Username: username, Hash: hash, Role: role}
	return nil
}

func authRequired(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		token, err := jwt.Parse(strings.TrimPrefix(header, "Bearer "), func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		}, jwt.WithValidMethods([]string{"HS256"}))
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired token"})
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Set("username", claims["username"])
		c.Set("role", claims["role"])
		c.Next()
	}
}


func roleRequired(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("role") != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission for this"})
			return
		}
		c.Next()
	}
}

func main() {
	port := ask("Enter port number (press Enter for 8080): ")
	if port == "" {
		port = "8080"
	}

	secret := ask("Enter a secret key used to sign tokens: ")
	adminName := ask("Create an admin username: ")
	adminPass := ask("Create an admin password: ")
	userName := ask("Create a normal user username: ")
	userPass := ask("Create a normal user password: ")
	if secret == "" || adminName == "" || adminPass == "" || userName == "" || userPass == "" {
		fmt.Println("Every value is needed.")
		return
	}
	if adminName == userName {
		fmt.Println("The two usernames must be different.")
		return
	}

	if err := addUser(adminName, adminPass, "admin"); err != nil {
		fmt.Println("Could not create the admin:", err)
		return
	}
	if err := addUser(userName, userPass, "user"); err != nil {
		fmt.Println("Could not create the user:", err)
		return
	}

	router := gin.New()

	router.POST("/login", func(c *gin.Context) {
		var input LoginRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "send username and password as JSON"})
			return
		}

		user, found := users[input.Username]
		if !found || bcrypt.CompareHashAndPassword(user.Hash, []byte(input.Password)) != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"role":     user.Role,
			"exp":      time.Now().Add(time.Hour).Unix(),
		})
		signed, _ := token.SignedString([]byte(secret))
		c.JSON(http.StatusOK, gin.H{"token": signed, "role": user.Role})
	})

	api := router.Group("/api")
	api.Use(authRequired(secret))

	api.GET("/profile", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"username": c.GetString("username"), "role": c.GetString("role")})
	})

	api.GET("/admin/report", roleRequired("admin"), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "secret admin report"})
	})

	if err := router.Run(":" + port); err != nil {
		fmt.Println("Could not start the server:", err)
	}
}
