package auth

import (
	"context"
	"testing"
	"time"
)

func TestTokenRoundTrip(t *testing.T) {
	ts := NewTokenService("secret", time.Minute)
	token, err := ts.Generate(7, "admin", RoleAdmin)
	if err != nil {
		t.Fatalf("generate failed: %v", err)
	}
	claims, err := ts.Parse(token)
	if err != nil {
		t.Fatalf("parse failed: %v", err)
	}
	if claims.UserID != 7 || claims.Username != "admin" || claims.Role != RoleAdmin {
		t.Fatalf("wrong claims: %+v", claims)
	}
}

func TestTokenWithWrongSecretIsRejected(t *testing.T) {
	token, _ := NewTokenService("secret", time.Minute).Generate(1, "admin", RoleAdmin)
	if _, err := NewTokenService("other-secret", time.Minute).Parse(token); err == nil {
		t.Fatal("token signed with another secret must be rejected")
	}
}

func TestExpiredTokenIsRejected(t *testing.T) {
	ts := NewTokenService("secret", -time.Minute) // already expired
	token, _ := ts.Generate(1, "admin", RoleAdmin)
	if _, err := ts.Parse(token); err == nil {
		t.Fatal("expired token must be rejected")
	}
}

func TestPasswordHashing(t *testing.T) {
	hash, err := HashPassword("admin123")
	if err != nil {
		t.Fatalf("hash failed: %v", err)
	}
	if hash == "admin123" {
		t.Fatal("password must not be stored as plain text")
	}
	if !CheckPassword(hash, "admin123") {
		t.Fatal("correct password should match")
	}
	if CheckPassword(hash, "wrong") {
		t.Fatal("wrong password should not match")
	}
}

func TestRequireUserAndRole(t *testing.T) {
	if _, err := RequireUser(context.Background()); err != ErrUnauthenticated {
		t.Fatalf("expected ErrUnauthenticated, got %v", err)
	}
	userCtx := WithUser(context.Background(), &Claims{Username: "u", Role: RoleUser})
	if _, err := RequireUser(userCtx); err != nil {
		t.Fatalf("logged-in user should pass, got %v", err)
	}
	if _, err := RequireRole(userCtx, RoleAdmin); err != ErrForbidden {
		t.Fatalf("expected ErrForbidden, got %v", err)
	}
	adminCtx := WithUser(context.Background(), &Claims{Username: "a", Role: RoleAdmin})
	if _, err := RequireRole(adminCtx, RoleAdmin); err != nil {
		t.Fatalf("admin should pass, got %v", err)
	}
}
