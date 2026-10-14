package main

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Hash     []byte
	Role     string
}

type Auth struct {
	secret []byte
	users  map[string]User
}

func NewAuth(secret string) *Auth {
	return &Auth{secret: []byte(secret), users: map[string]User{}}
}


func (a *Auth) AddUser(username, password, role string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.users[username] = User{Username: username, Hash: hash, Role: role}
	return nil
}


func (a *Auth) Login(username, password string) (string, string, error) {
	user, found := a.users[username]
	if !found || bcrypt.CompareHashAndPassword(user.Hash, []byte(password)) != nil {
		return "", "", errors.New("invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})
	signed, err := token.SignedString(a.secret)
	if err != nil {
		return "", "", err
	}
	return signed, user.Role, nil
}


func (a *Auth) Required() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		token, err := jwt.Parse(strings.TrimPrefix(header, "Bearer "), func(t *jwt.Token) (interface{}, error) {
			return a.secret, nil
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


func (a *Auth) RoleRequired(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("role") != role {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "you do not have permission for this"})
			return
		}
		c.Next()
	}
}
