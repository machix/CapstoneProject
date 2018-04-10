package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

// Interface implementation for mockDB testing
func (mdb *mockDB) GetPolygons(p *model.PolygonSummary) error {
	p = &model.PolygonSummary{"Test"}
	return nil
}

// Test get Polygon endpoing
func TestGetPolygons(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getPolygons", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.GetPosition).ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "{\"UserSummary\":[{\"Id\":1,\"Latitude\":1.23,\"Longitude\":1.23}]}"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

// Interface implementation for mockDB testing
func (mdb *mockDB) SavePolygon(p *model.Polygon, c *model.Client) error {
	//p = &model.Polygon{Id: "1", Name: "polygon1", Coordinates: cordTestArray}
	c = &model.Client{ID: 1, FirstName: "Test", LastName: "Testy"}
	return nil
}

// Test saving polygon endpoint
func TestSavePolygon(t *testing.T) {
	var b bytes.Buffer
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/savePolygon", &b)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.PostPosition).ServeHTTP(rec, req)

	expected := "{\"Id\":0,\"Latitude\":0,\"Longitude\":0}"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

// Interface implementation for mockDB testing
func (mdb *mockDB) DeletePolygon(p *model.Polygon, c *model.Client) error {
	return nil
}

// Test the delete polygon endpoint
func TestDeletePolygon(t *testing.T) {

}
