
package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

const (
	RoleAdmin = "ADMIN" 
	RoleUser  = "USER"  
)

var (
	ErrUnauthenticated = errors.New("authentication required")
	ErrForbidden       = errors.New("you do not have permission to do this")
)


type Claims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
