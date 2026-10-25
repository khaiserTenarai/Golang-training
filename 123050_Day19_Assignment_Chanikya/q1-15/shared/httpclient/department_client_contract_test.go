
package httpclient

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestDepartmentClientContract checks that the Department Service
// follows the response contract expected by the Employee Service.
func TestDepartmentClientContract(t *testing.T) {
	t.Run("valid department response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet {
					t.Errorf("expected GET, got %s", r.Method)
				}

				if r.URL.Path != "/api/departments/7" {
					t.Errorf("unexpected path: %s", r.URL.Path)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				_ = json.NewEncoder(w).Encode(
					DepartmentResponse{
						ID:   7,
						Name: "Engineering",
					},
				)
			},
		))
		defer server.Close()

		client := newContractTestClient(server.URL, server.Client())

		department, err := client.GetByID(
			context.Background(),
			7,
			"",
		)
		if err != nil {
			t.Fatalf("expected success, got error: %v", err)
		}

		if department.ID != 7 {
			t.Errorf("expected department ID 7, got %d", department.ID)
		}

		if department.Name != "Engineering" {
			t.Errorf(
				"expected department name Engineering, got %q",
				department.Name,
			)
		}
	})

	t.Run("correlation ID is forwarded", func(t *testing.T) {
		const expectedID = "contract-test-123"

		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if got := r.Header.Get("X-Correlation-ID"); got != expectedID {
					t.Errorf("expected correlation ID %q, got %q", expectedID, got)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(
					`{"id":7,"name":"Engineering"}`,
				))
			},
		))
		defer server.Close()

		client := newContractTestClient(server.URL, server.Client())

		_, err := client.GetByID(
			context.Background(),
			7,
			expectedID,
		)
		if err != nil {
			t.Fatalf("expected success, got error: %v", err)
		}
	})

	t.Run("missing department returns not found", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				http.NotFound(w, r)
			},
		))
		defer server.Close()

		client := newContractTestClient(server.URL, server.Client())

		_, err := client.GetByID(
			context.Background(),
			999,
			"",
		)

		if !errors.Is(err, ErrDepartmentNotFound) {
			t.Fatalf("expected ErrDepartmentNotFound, got %v", err)
		}
	})

	t.Run("unexpected status returns unavailable error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				http.Error(
					w,
					"internal server error",
					http.StatusInternalServerError,
				)
			},
		))
		defer server.Close()

		client := newContractTestClient(server.URL, server.Client())

		_, err := client.GetByID(
			context.Background(),
			7,
			"",
		)

		if !errors.Is(err, ErrDepartmentUnavailable) {
			t.Fatalf("expected ErrDepartmentUnavailable, got %v", err)
		}
	})

	t.Run("invalid JSON returns unavailable error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`not-valid-json`))
			},
		))
		defer server.Close()

		client := newContractTestClient(server.URL, server.Client())

		_, err := client.GetByID(
			context.Background(),
			7,
			"",
		)

		if !errors.Is(err, ErrDepartmentUnavailable) {
			t.Fatalf("expected ErrDepartmentUnavailable, got %v", err)
		}
	})
}

func newContractTestClient(
	baseURL string,
	httpClient *http.Client,
) *DepartmentClient {
	return &DepartmentClient{
		BaseURL: baseURL,
		Client:  httpClient,
		Breaker: NewCircuitBreaker(3, 30*time.Second),
		Retries: 0,
	}
}
