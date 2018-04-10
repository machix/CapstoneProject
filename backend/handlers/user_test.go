package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

type mockDB struct{}

// Interface implementation for mockDB
func (mdb *mockDB) QueryPosition(us *model.Summary) error {
	us.UserSummary = append(us.UserSummary, model.User{Id: 1, Latitude: 1.23, Longitude: 1.23})
	return nil
}

// Test get user position
func TestGetPosition(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/getPosition", nil)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.GetPosition).ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check what is returned by handler from mockDB
	expected := "{\"UserSummary\":[{\"Id\":1,\"Latitude\":1.23,\"Longitude\":1.23}]}"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

// Interface implementation for mockDB
func (mdb *mockDB) PostPosition(u *model.User) error {
	u = &model.User{Id: 1, Latitude: 1.34, Longitude: 1.34}
	return nil
}

// Test post user position
func TestPostPosition(t *testing.T) {
	var b bytes.Buffer
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/postPosition", &b)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.PostPosition).ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check what is returned by handler from mockDB
	expected := "{\"Id\":0,\"Latitude\":0,\"Longitude\":0}"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

// Interface implementation for mockDB
func (mdb *mockDB) DeletePosition(u *model.User) error {
	u = &model.User{Id: 1, Latitude: 1.45, Longitude: 1.45}
	return nil
}

//Test delete user position
func TestDeletePosition(t *testing.T) {
	var b bytes.Buffer
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/deletePosition", &b)

	env := Env{db: &mockDB{}}
	http.HandlerFunc(env.DeletePosition).ServeHTTP(rec, req)

	// Check the status code is what we expect.
	if status := rec.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check what is returned by handler from mockDB
	expected := "{\"Id\":0,\"Latitude\":0,\"Longitude\":0}"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}
