package database

import (
	"testing"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// Test connection to the AWS Postgres database
func TestUserDatabaseConnect(t *testing.T) {
	// Add Test to test database connection
}

type Env struct {
	db model.Datastore
}

// Test get request to database for user
func TestDatabaseSelectQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	env := &DB{db}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"Id", "Latitude", "Longitude"}).AddRow(1, 4.5678, 5.4567)

	mock.ExpectBegin()
	mock.ExpectQuery(`[SELECT * FROM USER_LOCATION]`).WillReturnRows(rows)

	tempSummary := model.Summary{}

	// Test QueryPosition function
	if err = env.QueryPosition(&tempSummary); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// Ensure the Expectations match
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Test database insert query for user
func TestDatabaseInsertQuery(t *testing.T) {

}

// Test database delete query for user
func (env *Env) TestDatabaseDeleteQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	sqlmock.NewRows([]string{"Id", "Latitude", "Longitude"}).AddRow(1, 4.5678, 5.4567)

	mock.ExpectBegin()
	mock.ExpectPrepare(`[DELETE Id=? FROM USER_LOCATION]`).ExpectExec()

}
