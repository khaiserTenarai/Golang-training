package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"employee-api/config"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type ClaimsKeyType string

const ClaimsKey ClaimsKeyType = "claims"

type AuthService struct {
	Secret            []byte
	Expiry            time.Duration
	AdminUsername     string
	AdminPasswordHash string
	UserUsername      string
	UserPasswordHash  string
}

func NewAuth(cfg config.Config) *AuthService {
	return &AuthService{
		Secret:            []byte(cfg.JWTSecret),
		Expiry:            time.Duration(cfg.JWTExpiryMinutes) * time.Minute,
		AdminUsername:     cfg.AdminUsername,
		AdminPasswordHash: cfg.AdminPasswordHash,
		UserUsername:      cfg.UserUsername,
		UserPasswordHash:  cfg.UserPasswordHash,
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthService) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var role string
	var passwordHash string

	switch input.Username {
	case a.AdminUsername:
		role = "admin"
		passwordHash = a.AdminPasswordHash

	case a.UserUsername:
		role = "user"
		passwordHash = a.UserPasswordHash

	default:
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Reject login if the password hash is missing or the password is incorrect.
	if passwordHash == "" ||
		bcrypt.CompareHashAndPassword(
			[]byte(passwordHash),
			[]byte(input.Password),
		) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	now := time.Now()

	claims := Claims{
		Username: input.Username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(a.Expiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(a.Secret)
	if err != nil {
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_ = json.NewEncoder(w).Encode(map[string]string{
		"token": signedToken,
		"role":  role,
	})
}

func RequireAuth(secret []byte) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			parts := strings.SplitN(authHeader, " ", 2)

			if len(parts) != 2 ||
				!strings.EqualFold(parts[0], "Bearer") ||
				parts[1] == "" {
				http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
				return
			}

			claims := &Claims{}

			token, err := jwt.ParseWithClaims(
				parts[1],
				claims,
				func(token *jwt.Token) (any, error) {
					if token.Method != jwt.SigningMethodHS256 {
						return nil, jwt.ErrSignatureInvalid
					}
					return secret, nil
				},
			)

			if err != nil || !token.Valid {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), ClaimsKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireRole(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := r.Context().Value(ClaimsKey).(*Claims)

			if !ok || claims.Role != role {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
