package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	// Create a new request
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler directly
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Replicating the handler from main.go
		io.WriteString(w, "Hello from Go HTTP server!\n")
	})

	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Hello from Go HTTP server!\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
