// Task 12: Correlation ID — tracks a request across microservices.

package middleware

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"time"
)

type ctxKey string

const CorrelationKey ctxKey = "correlationID"

func CorrelationID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := r.Header.Get("X-Correlation-ID")
		if cid == "" {
			b := make([]byte, 8)
			rand.Read(b)
			cid = fmt.Sprintf("cid-%x", b)
		}
		w.Header().Set("X-Correlation-ID", cid)
		ctx := context.WithValue(r.Context(), CorrelationKey, cid)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetCorrelationID(ctx context.Context) string {
	if cid, ok := ctx.Value(CorrelationKey).(string); ok {
		return cid
	}
	return ""
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lw := &statusWriter{ResponseWriter: w, code: 200}
		next.ServeHTTP(lw, r)
		cid := GetCorrelationID(r.Context())
		log.Printf("[%s] %s %s → %d (%v) [%s]",
			r.Method, r.URL.Path, r.RemoteAddr, lw.code,
			time.Since(start).Round(time.Microsecond), cid)
	})
}

type statusWriter struct {
	http.ResponseWriter
	code int
}

func (w *statusWriter) WriteHeader(code int) {
	w.code = code
	w.ResponseWriter.WriteHeader(code)
}
