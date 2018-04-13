package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Test create geofence endpoint
func TestCreateGeofence(t *testing.T) {
	rec := httptest.NewRecorder()
	json := `{"Id": 1, "Name": "polygon1", "Coordinates": "[ {Latitude: 1.23, Longitude: 1.23}, {Latitude: 2.34, Longitude: 2.34},
		{Latitude: 4.56, Longitude: 4.56}, {Latitude: 1.23, Longitude: 1.23}]" }`
	req, _ := http.NewRequest("POST", "/createGeofence", strings.NewReader(json))
	req.Header.Set("Content-Type", "application/json")

	http.HandlerFunc(CreateGeofence).ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// Test point in polygon endpoint
func TestCheckPointInPolygon(t *testing.T) {
	rec := httptest.NewRecorder()
	json := `{"Latitude": 1.23, "Longitude": 1.23}`
	req, _ := http.NewRequest("POST", "checkGeofence", strings.NewReader(json))
	req.Header.Set("Content-Type", "application/json")

	http.HandlerFunc(CheckPointInPolygon).ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// Test polygon overlap handler
func TestPolygonOverlap(t *testing.T) {

}
