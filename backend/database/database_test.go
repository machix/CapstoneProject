package database

import (
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

//Test connection to the AWS Postgres database
func TestDatabaseConnect(t *testing.T) {
	//Add Test to test database connection
}

// Test get request to database
func TestDatabaseSelectQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Test QueryPosition function
	if err = QueryPosition(db, 2, 3); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// Ensure the Expectations match
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Test database delete query
func TestDatabaseDeleteQuery(t *testing.T) {
	//Unit test for deleting from the database
}

// Test database update query
func TestDatabaseUpdateQuery(t *testing.T) {
	//Unit test for updating database
}
