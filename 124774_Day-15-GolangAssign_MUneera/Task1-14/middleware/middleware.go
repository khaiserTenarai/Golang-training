package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"time"

	"employee-management/logging"

	"github.com/gin-gonic/gin"
)

// Question 1: Custom middleware.
func CustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Custom middleware: BEFORE")

		c.Next()

		fmt.Println("Custom middleware: AFTER")
	}
}

// Question 2 and 3: Middleware order and Next().
func SecondMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Second middleware: BEFORE")

		c.Next()

		fmt.Println("Second middleware: AFTER")
	}
}

// Question 4: Abort().
// Question 4: Abort()
func AbortDemo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"error": "request stopped by Abort middleware",
		})
	}
}

// Question 5: Request ID middleware.
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := fmt.Sprintf("%d", time.Now().UnixNano())

		c.Set("requestID", id)
		c.Header("X-Request-ID", id)

		c.Next()
	}
}

// Question 6: Recovery middleware.
func Recovery(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if recovered := recover(); recovered != nil {
				logger.Error(fmt.Sprintf(
					"panic recovered: %v\n%s",
					recovered,
					debug.Stack(),
				))

				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					gin.H{"error": "internal server error"},
				)
			}
		}()

		c.Next()
	}
}

// Question 7: CORS middleware.
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, DELETE, OPTIONS",
		)
		c.Header(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization, X-Request-ID",
		)

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// Request logging middleware.
func Logging(logger *logging.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		logger.Info(fmt.Sprintf(
			"%s %s %d %s",
			c.Request.Method,
			c.Request.URL.RequestURI(),
			c.Writer.Status(),
			time.Since(start),
		))
	}
}

// JWT claims stored inside the token.
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	Expires  int64  `json:"exp"`
}

// GenerateToken creates a signed HS256 JWT.
func GenerateToken(
	username string,
	role string,
	secret string,
	expiryMinutes int,
) (string, error) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	claims := Claims{
		Username: username,
		Role:     role,
		Expires: time.Now().Add(
			time.Duration(expiryMinutes) * time.Minute,
		).Unix(),
	}

	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	encodedHeader := base64.RawURLEncoding.EncodeToString(headerJSON)
	encodedClaims := base64.RawURLEncoding.EncodeToString(claimsJSON)

	unsignedToken := encodedHeader + "." + encodedClaims

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(unsignedToken))

	signature := base64.RawURLEncoding.EncodeToString(mac.Sum(nil))

	return unsignedToken + "." + signature, nil
}

// verifyToken checks the signature and expiration.
func verifyToken(token string, secret string) (*Claims, error) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	unsignedToken := parts[0] + "." + parts[1]

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(unsignedToken))

	expectedSignature := mac.Sum(nil)

	receivedSignature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, fmt.Errorf("invalid token signature")
	}

	if !hmac.Equal(expectedSignature, receivedSignature) {
		return nil, fmt.Errorf("invalid token signature")
	}

	headerJSON, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid token header")
	}

	var header map[string]string
	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return nil, fmt.Errorf("invalid token header")
	}

	if header["alg"] != "HS256" {
		return nil, fmt.Errorf("unsupported token algorithm")
	}

	claimsJSON, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid token claims")
	}

	var claims Claims
	if err := json.Unmarshal(claimsJSON, &claims); err != nil {
		return nil, fmt.Errorf("invalid token claims")
	}

	if claims.Expires <= time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}

	return &claims, nil
}

// Question 8 and 9: JWT authentication and protected routes.
func Auth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if !strings.HasPrefix(header, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing bearer token",
			})
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(header, "Bearer "))

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing bearer token",
			})
			return
		}

		claims, err := verifyToken(token, secret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		c.Set("username", claims.Username)
		c.Set("role", claims.Role)

		c.Next()
	}
}

// Question 10: Role-based authorization.
func RequireRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		actualRole, exists := c.Get("role")

		if !exists || actualRole != requiredRole {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "forbidden: insufficient role",
			})
			return
		}

		c.Next()
	}
}
