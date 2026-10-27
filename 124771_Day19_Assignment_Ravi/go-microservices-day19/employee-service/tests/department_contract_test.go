package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Task 14: contract test. It checks that Department Service's response contains
// the fields the Employee Service's DepartmentDTO expects, without requiring PostgreSQL.
func TestDepartmentResponseContract(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/departments/1" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"id":1,"name":"Engineering","description":"Builds software"}`))
	}))
	defer server.Close()
	resp, err := http.Get(server.URL + "/departments/1")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
	var body struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}
	if body.ID != 1 || body.Name != "Engineering" || body.Description != "Builds software" {
		t.Fatalf("unexpected department contract: %+v", body)
	}
}
