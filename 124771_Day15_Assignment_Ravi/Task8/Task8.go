package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret []byte

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	Username     string
	PasswordHash []byte
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Demo user store. Replace with a database in production.
var users map[string]User

func main() {
	secret := os.Getenv("JWT_SECRET")
	if len(secret) < 32 {
		log.Fatal("JWT_SECRET must contain at least 32 characters")
	}
	jwtSecret = []byte(secret)

	// Demo credentials: admin / password123
	hash, err := bcrypt.GenerateFromPassword(
		[]byte("password123"), bcrypt.DefaultCost,
	)
	if err != nil {
		log.Fatal(err)
	}
	users = map[string]User{
		"admin": {Username: "admin", PasswordHash: hash},
	}

	r := gin.Default()

	r.POST("/login", loginHandler)

	protected := r.Group("/api")
	protected.Use(authMiddleware())
	{
		protected.GET("/profile", profileHandler)
		protected.GET("/dashboard", dashboardHandler)
	}

	log.Println("Server listening on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func loginHandler(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	user, exists := users[req.Username]
	if !exists ||
		bcrypt.CompareHashAndPassword(
			user.PasswordHash, []byte(req.Password),
		) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	now := time.Now()

	claims := Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "gin-jwt-demo",
			Subject:   user.Username,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(15 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not generate token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenString,
		"token_type":   "Bearer",
		"expires_in":   900,
	})
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		parts := strings.Fields(header)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Missing or invalid Authorization header",
			})
			return
		}

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(
			parts[1],
			claims,
			func(token *jwt.Token) (interface{}, error) {
				// Prevent accepting a token signed with another algorithm.
				if token.Method != jwt.SigningMethodHS256 {
					return nil, jwt.ErrSignatureInvalid
				}
				return jwtSecret, nil
			},
			jwt.WithValidMethods([]string{"HS256"}),
			jwt.WithIssuer("gin-jwt-demo"),
			jwt.WithExpirationRequired(),
		)

		if err != nil || token == nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}

func profileHandler(c *gin.Context) {
	username, _ := c.Get("username")

	c.JSON(http.StatusOK, gin.H{
		"message":  "Authenticated successfully",
		"username": username,
	})
}

func dashboardHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the protected dashboard",
	})
}
