package database

import (
	"testing"
    "fmt"

	"github.com/NaturalFractals/CapstoneProject/backend/model"
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

	rows := sqlmock.NewRows([]string{"Id", "Latitude", "Longitude"}).AddRow(1, 4.5678, 5.4567)

	mock.ExpectBegin()
	mock.ExpectQuery(`[SELECT * FROM USER_LOCATION]`).WillReturnRows(rows)

	tempSummary := model.Summary{}

	// Test QueryPosition function
	if err = QueryPosition(&tempSummary, db); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	// Ensure the Expectations match
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

// Test database delete query
func TestDatabaseDeleteQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlmock.NewRows([]string{"Id", "Latitude", "Longitude"}).AddRow(1, 4.5678, 5.4567)

	mock.ExpectBegin()
	mock.ExpectPrepare(`[DELETE Id=? FROM USER_LOCATION]`).ExpectExec()

}

// Test database update query
func TestDatabaseUpdateQuery(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlmock.NewRows([]string{"Id", "Latitude", "Longitude"}).AddRow(1, 4.5678, 5.4567)

	mock.ExpectBegin()
	mock.ExpectPrepare("UPDATE USER_LOCATION SET Latitude=?")
	mock.ExpectExec("UPDATE USER_LOCATION SET Latitude=?").WillReturnResult(sqlmock.NewResult(1, 1))

	tempUser := model.User{}
    
	err = UpdatePosition(&tempUser, db)
	if err != nil {
		t.Errorf("Error was not expected while updating stats: %s", err)
	}

    fmt.Printf("%s, %g, %g", tempUser.Id, tempUser.Latitude, tempUser.Longitude) 
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("Expectations were not met: %s", err)
	}
}
