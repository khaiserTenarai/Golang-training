// Package auth contains everything about security: roles, JWT tokens,
// password hashing and "who is the current user" helpers.
package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

const (
	RoleAdmin = "ADMIN" // can read and change data
	RoleUser  = "USER"  // can only read data
)

var (
	ErrUnauthenticated = errors.New("authentication required")
	ErrForbidden       = errors.New("you do not have permission to do this")
)

// Claims is the information stored INSIDE the login token.
type Claims struct {
	UserID   uint   `json:"userId"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
