package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test get user
func TestGetPosition(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/position", nil)
	w := httptest.NewRecorder()

	Router().ServeHTTP(w, r)
}

// Test create user
func TestDeletePosition(t *testing.T) {
}

// Test remove user
func TestPostPosition(t *testing.T) {
}
