// Package controller contains HTTP request handlers.
package controller

import (
	"employee-management/auth"
	"encoding/json"
	"net/http"
)

type AuthController struct{ auth *auth.AuthService }

// NewAuthController creates the authentication controller.
func NewAuthController(a *auth.AuthService) *AuthController { return &AuthController{auth: a} }

// Login authenticates a user and returns a JWT.
func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if json.NewDecoder(r.Body).Decode(&req) != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}
	token, err := c.auth.Login(req.Username, req.Password)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"token": token, "tokenType": "Bearer"})
}
