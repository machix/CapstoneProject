package database

import (
	"testing"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

// Test database delete polygon for client database
func TestClientDatabaseDeletePolygon(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlmock.NewRows([]string{"id", "name", "polygon"}).AddRow(1, 4.5678, 5.4567)

	mock.ExpectBegin()
	mock.ExpectPrepare(`[DELETE id=? FROM CLIENT_POLYGON]`).ExpectExec()

}

// Test database fetching all current clients in database
func TestClientDatabaseGetClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name"}).AddRow(1, 4.5678, 5.4567)

	mock.ExpectBegin()
	mock.ExpectQuery(`[SELECT * FROM CLIENT]`).WillReturnRows(rows)
}

// Test database adding a new client to database
func TestClientDatabaseAddNewClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
}

// Test databse deleting a client from database
func TestClientDatabaseDeleteClient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlmock.NewRows([]string{"id", "first_name", "last_name"}).AddRow(1, 4.5678, 5.4567)

	mock.ExpectBegin()
	mock.ExpectPrepare(`[DELETE id=? FROM CLIENT]`).ExpectExec()
}
