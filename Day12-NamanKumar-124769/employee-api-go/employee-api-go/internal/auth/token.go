package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type Claims struct {
	Subject string `json:"sub"`
	Role    string `json:"role"`
	Expires int64  `json:"exp"`
}

type TokenManager struct {
	secret []byte
	ttl    time.Duration
}

func NewTokenManager(secret []byte, ttl time.Duration) *TokenManager {
	return &TokenManager{secret: secret, ttl: ttl}
}

func (m *TokenManager) Issue(subject, role string) (string, error) {
	claims := Claims{Subject: subject, Role: role, Expires: time.Now().Add(m.ttl).Unix()}
	payload, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}
	encoded := base64.RawURLEncoding.EncodeToString(payload)
	signature := m.sign(encoded)
	return encoded + "." + signature, nil
}

func (m *TokenManager) Verify(token string) (Claims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 2 {
		return Claims{}, errors.New("invalid token")
	}
	expected := m.sign(parts[0])
	if !hmac.Equal([]byte(expected), []byte(parts[1])) {
		return Claims{}, errors.New("invalid token signature")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return Claims{}, errors.New("invalid token payload")
	}
	var claims Claims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return Claims{}, errors.New("invalid token claims")
	}
	if claims.Subject == "" || claims.Role == "" || time.Now().Unix() >= claims.Expires {
		return Claims{}, errors.New("token expired or invalid")
	}
	return claims, nil
}

func (m *TokenManager) sign(value string) string {
	mac := hmac.New(sha256.New, m.secret)
	mac.Write([]byte(value))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}
