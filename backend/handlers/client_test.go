package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// Interface implementation for mockDB testing
func (mdb *mockDB) GetClients(c *model.ClientSummary) error {
	c.ClientSummary = append(c.ClientSummary, model.Client{ID: 1, FirstName: "Test", LastName: "Testy"})
	return nil
}

// Test the get client endpoint
func TestGetClient(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getClient", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.GetClient).ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check what is returned by handler from mockDB
	expected := "{\"ClientSummary\":[{\"Id\":1,\"FirstName\":\"Test\",\"LastName\":\"Testy\"}]}"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

// Interface implementation for mockDB testing
func (mdb *mockDB) AddNewClient(c *model.Client) error {
	c = &model.Client{ID: 1, FirstName: "Test", LastName: "Testy"}
	return nil
}

// Test the create client endpoint
func TestCreateClient(t *testing.T) {
	var b bytes.Buffer
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/postClient", &b)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.CreateClient).ServeHTTP(rec, req)

	// Check to ensure successful status from handler
	if status := rec.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}
}

// Interface implementation for mockDB testing
func (mdb *mockDB) DeleteClient(c *model.Client) error {
	return nil
}

// Test the delete client endpoint
func TestDeleteClient(t *testing.T) {
	var b bytes.Buffer
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/deleteClient", &b)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.RemoveClient).ServeHTTP(rec, req)

	// Check for successful status code
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
