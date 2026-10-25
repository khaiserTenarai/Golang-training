package auth

import (
	"strings"
	"testing"
)

func TestLoginAndParseToken(t *testing.T) {
	a, err := New(
		strings.Repeat("x", 40), 60,
		"admin", "AdminPassword123!",
		"user", "UserPassword123!",
	)
	if err != nil {
		t.Fatal(err)
	}

	token, err := a.Login("admin", "AdminPassword123!")
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	claims, err := a.ParseAuthorization("Bearer " + token)
	if err != nil {
		t.Fatalf("ParseAuthorization failed: %v", err)
	}

	if claims.Subject != "admin" || claims.Role != "admin" {
		t.Fatalf("unexpected claims: %+v", claims)
	}
}

func TestLoginRejectsIncorrectPassword(t *testing.T) {
	a, err := New(
		strings.Repeat("x", 40), 60,
		"admin", "AdminPassword123!",
		"user", "UserPassword123!",
	)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.Login("admin", "wrong"); err == nil {
		t.Fatal("expected invalid credentials to be rejected")
	}
}

func TestParseRejectsMissingBearerToken(t *testing.T) {
	a, err := New(
		strings.Repeat("x", 40), 60,
		"admin", "AdminPassword123!",
		"user", "UserPassword123!",
	)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := a.ParseAuthorization(""); err == nil {
		t.Fatal("expected missing token to be rejected")
	}
}
