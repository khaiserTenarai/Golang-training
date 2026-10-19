package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/graphql-go/handler"
)

func TestGraphQLQueryAndMutation(t *testing.T) {
	schema, err := BuildGraphQLSchema()
	if err != nil {
		t.Fatalf("Failed to build schema: %v", err)
	}

	h := AuthMiddleware(handler.New(&handler.Config{
		Schema: &schema,
	}))

	server := httptest.NewServer(h)
	defer server.Close()

	// 1. Test Query
	query := `{"query": "{ employees(department: \"Engineering\") { name department { name } } }"}`
	resp, err := http.Post(server.URL, "application/json", bytes.NewBufferString(query))
	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fatalf("Query failed: %v", err)
	}

	var qResult map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&qResult)
	if qResult["data"] == nil {
		t.Fatalf("Query data returned nil")
	}

	// 2. Test Mutation Authorization Failure (Non-Admin)
	mutation := `{"query": "mutation { createEmployee(name: \"Test\", email: \"test@test.com\", departmentId: 101, salary: 50000) { id } }"}`
	respErr, _ := http.Post(server.URL, "application/json", bytes.NewBufferString(mutation))
	var mResult map[string]interface{}
	json.NewDecoder(respErr.Body).Decode(&mResult)

	if mResult["errors"] == nil {
		t.Fatalf("Expected authorization error for non-admin mutation, got none")
	}
}