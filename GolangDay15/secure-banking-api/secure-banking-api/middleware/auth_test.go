package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"secure-banking-api/util"

	"github.com/gin-gonic/gin"
)

const testSecret = "test-secret"

func newRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(RequestID(), Recovery())
	r.GET("/me", JWTAuth(testSecret), func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	r.GET("/admin", JWTAuth(testSecret), RequireRole("admin"), func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	return r
}

func call(r http.Handler, path, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodGet, path, nil)
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestAuthAndRBAC(t *testing.T) {
	r := newRouter()
	userTok, _ := util.GenerateToken(testSecret, 1, "user", time.Minute)
	adminTok, _ := util.GenerateToken(testSecret, 2, "admin", time.Minute)
	expired, _ := util.GenerateToken(testSecret, 1, "user", -time.Minute)
	wrongKey, _ := util.GenerateToken("other-secret", 1, "user", time.Minute)

	tests := []struct {
		name  string
		path  string
		token string
		want  int
	}{
		{"no token", "/me", "", http.StatusUnauthorized},
		{"garbage token", "/me", "abc.def.ghi", http.StatusUnauthorized},
		{"expired token", "/me", expired, http.StatusUnauthorized},
		{"wrong signing key", "/me", wrongKey, http.StatusUnauthorized},
		{"valid user", "/me", userTok, http.StatusOK},
		{"user on admin route", "/admin", userTok, http.StatusForbidden},
		{"admin on admin route", "/admin", adminTok, http.StatusOK},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := call(r, tc.path, tc.token).Code; got != tc.want {
				t.Fatalf("want %d, got %d", tc.want, got)
			}
		})
	}
}

func TestRecoveryReturns500(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(Recovery())
	r.GET("/boom", func(c *gin.Context) { panic("boom") })

	if got := call(r, "/boom", "").Code; got != http.StatusInternalServerError {
		t.Fatalf("want 500, got %d", got)
	}
}
