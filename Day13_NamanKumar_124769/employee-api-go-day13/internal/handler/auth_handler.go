package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"employee-api-go/internal/auth"
)

type AuthHandler struct {
	tokens *auth.TokenManager
}

func NewAuthHandler(tokens *auth.TokenManager) *AuthHandler {
	return &AuthHandler{tokens: tokens}
}

type tokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *AuthHandler) Token(w http.ResponseWriter, r *http.Request) {
	var input tokenRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20)).Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid JSON body"})
		return
	}

	role := ""
	switch {
	case input.Username == "admin" && input.Password == "admin123":
		role = "admin"
	case input.Username == "user" && input.Password == "user123":
		role = "user"
	default:
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
		return
	}

	token, err := h.tokens.Issue(input.Username, role)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not issue token"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{
		"access_token": token,
		"token_type":   "Bearer",
		"expires_in":   int((8 * time.Hour).Seconds()),
		"role":         strings.ToLower(role),
	})
}
