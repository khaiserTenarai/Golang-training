// Package auth implements authentication and JWT token creation/validation.

package auth

import (
	// crypto/sha256 is used to convert a password into a hash.
	"crypto/sha256"

	// encoding/hex converts the hash into a hexadecimal string.
	"encoding/hex"

	// errors is used to create error messages.
	"errors"

	// strings is used to process the Authorization header.
	"strings"

	// time is used to calculate JWT expiration time.
	"time"

	// JWT library is used to create and validate JWT tokens.
	"github.com/golang-jwt/jwt/v5"
)

// User represents one application user.
type User struct {

	// Username stores the user's login username.
	Username string

	// PasswordHash stores the hashed password.
	// We store the hash instead of the actual password.
	PasswordHash string

	// Role tells us what type of user this is.
	// Example: admin or user.
	Role string
}

// AuthService contains everything required for authentication.
type AuthService struct {

	// Secret is used to sign JWT tokens
	// and later verify the token.
	Secret []byte

	// Expiry stores how long a JWT should be valid.
	Expiry time.Duration

	// Admin stores the configured admin user.
	Admin User

	// Normal stores the configured normal user.
	Normal User
}

// New creates and returns an AuthService.
//
// It receives:
// 1. JWT secret
// 2. JWT expiry time
// 3. Admin username and password hash
// 4. Normal user username and password hash
func New(secret string, expiryMinutes int, adminUsername, adminHash, userUsername, userHash string) *AuthService {

	// Create an AuthService object and return its address.
	//
	// []byte(secret):
	// Converts the JWT secret from string to bytes.
	//
	// time.Duration(expiryMinutes) * time.Minute:
	// Converts minutes into a Go time duration.
	//
	// Admin:
	// Creates a user with the "admin" role.
	//
	// Normal:
	// Creates a user with the "user" role.
	return &AuthService{
		Secret: []byte(secret),
		Expiry: time.Duration(expiryMinutes) * time.Minute,
		Admin:  User{adminUsername, adminHash, "admin"},
		Normal: User{userUsername, userHash, "user"},
	}
}

// Login validates username/password and returns a JWT.
func (a *AuthService) Login(username, password string) (string, error) {

	// Search for the user using the username.
	u := a.findUser(username)

	// Check whether:
	// 1. The user exists.
	// 2. The given password matches the stored password hash.
	//
	// If either condition is false, login fails.
	if u == nil || !compareHash(password, u.PasswordHash) {

		// Return an error instead of creating a JWT.
		return "", errors.New("invalid username or password")
	}

	// Create the information that will be stored inside the JWT.
	//
	// sub = username
	// role = user's role
	// exp = token expiration time
	claims := jwt.MapClaims{
		"sub":  u.Username,
		"role": u.Role,
		"exp":  time.Now().Add(a.Expiry).Unix(),
	}

	// Create a new JWT using the claims.
	//
	// HS256 is the signing algorithm used for this token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using our secret.
	//
	// This returns the final JWT string.
	return token.SignedString(a.Secret)
}

// Parse validates a bearer token and returns username and role.
func (a *AuthService) Parse(header string) (string, string, error) {

	// Split the Authorization header into separate words.
	//
	// Example:
	// "Bearer abc123"
	//
	// becomes:
	// ["Bearer", "abc123"]
	parts := strings.Fields(header)

	// Check whether the Authorization header has the
	// correct Bearer format.
	//
	// Expected format:
	// Bearer <token>
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {

		// Header format is incorrect.
		return "", "", errors.New("invalid authorization header")
	}

	// Parse and validate the JWT token.
	//
	// The second parameter is a function that gives the JWT
	// library the secret needed to verify the token.
	token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {

		// Make sure the token uses the signing algorithm
		// that our application expects.
		if t.Method != jwt.SigningMethodHS256 {

			// Reject the token if another algorithm is used.
			return nil, errors.New("unexpected signing method")
		}

		// Return our secret so the JWT library can
		// verify the token signature.
		return a.Secret, nil
	})

	// Check whether:
	// 1. Parsing produced an error.
	// 2. The token is not valid.
	//
	// This can include an expired or invalid JWT.
	if err != nil || !token.Valid {

		// Reject the request.
		return "", "", errors.New("invalid or expired token")
	}

	// Get the claims stored inside the JWT.
	//
	// MapClaims contains values such as:
	//
	// sub  -> username
	// role -> user role
	// exp  -> expiry time
	claims, ok := token.Claims.(jwt.MapClaims)

	// Make sure the claims have the expected type.
	if !ok {

		// Claims are not in the expected format.
		return "", "", errors.New("invalid token claims")
	}

	// Read the username from the "sub" claim.
	user, _ := claims["sub"].(string)

	// Read the role from the "role" claim.
	role, _ := claims["role"].(string)

	// Return the username, role and no error.
	//
	// Example:
	// user = "ganesh"
	// role = "admin"
	return user, role, nil
}

// findUser searches for a configured user using the username.
func (a *AuthService) findUser(username string) *User {

	// Check whether the username belongs to the admin user.
	if username == a.Admin.Username {

		// Return the admin user.
		return &a.Admin
	}

	// Check whether the username belongs to the normal user.
	if username == a.Normal.Username {

		// Return the normal user.
		return &a.Normal
	}

	// Username was not found.
	return nil
}

// compareHash compares the given password with the stored password hash.
func compareHash(password, expected string) bool {

	// Convert the given password into a SHA-256 hash.
	//
	// Example:
	//
	// password
	//    ↓
	// SHA-256
	//    ↓
	// generated hash
	sum := sha256.Sum256([]byte(password))

	// Convert the generated hash into a hexadecimal string.
	//
	// Then compare it with the stored hash.
	//
	// true  = password is correct
	// false = password is incorrect
	return hex.EncodeToString(sum[:]) == expected
}
