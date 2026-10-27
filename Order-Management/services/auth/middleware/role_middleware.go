package middleware

import (
	"net/http"
	"strings"

	"order-management/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// RequireAuth validates the JWT from the Authorization header.
func RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")
		parts := strings.SplitN(header, " ", 2)

		if len(parts) != 2 ||
			!strings.EqualFold(parts[0], "Bearer") ||
			strings.TrimSpace(parts[1]) == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing or invalid Authorization header",
			})
			return
		}

		secret := []byte(config.Get(
			"JWT_SECRET",
			"replace-with-a-long-random-secret-at-least-32-characters",
		))

		token, err := jwt.ParseWithClaims(
			parts[1],
			&Claims{},
			func(token *jwt.Token) (interface{}, error) {
				if token.Method != jwt.SigningMethodHS256 {
					return nil, jwt.ErrSignatureInvalid
				}
				return secret, nil
			},
		)

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		claims, ok := token.Claims.(*Claims)
		if !ok || claims.Username == "" ||
			(claims.Role != "admin" && claims.Role != "user") {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			return
		}

		ctx.Set("username", claims.Username)
		ctx.Set("role", claims.Role)
		ctx.Next()
	}
}

// RequireRole allows only the specified role to access a route.
func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		role, exists := ctx.Get("role")
		if !exists || role != requiredRole {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "insufficient permissions",
			})
			return
		}

		ctx.Next()
	}
}
