package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
)

type mockDB struct{}

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

	expected := "{\"UserSummary\":[{\"Id\":1,\"Latitude\":1.23,\"Longitude\":1.23}]}"
	if expected != rec.Body.String() {
		t.Errorf("\n...expected = %v\n...obtained = %v", expected, rec.Body.String())
	}
}

func (mdb *mockDB) DeletePosition(u *model.User) error {
	return nil
}

//Test delete user position
func TestDeletePosition(t *testing.T) {
}

func (mdb *mockDB) PostPosition(u *model.User) error {
	return nil
}

// Test post user position
func TestPostPosition(t *testing.T) {
}
