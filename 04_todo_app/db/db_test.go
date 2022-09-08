package db

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestDBInit(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	err = db.Ping()
	assert.NoError(t, err)
}

func TestListItems(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mockDB := Database{db}

	rows := sqlmock.NewRows([]string{"id", "item"}).
		AddRow(1, "laundry").
		AddRow(2, "pay bills")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	todos, err := mockDB.ListItems()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.NotNil(t, todos)
	assert.NoError(t, err)

}

func TestAddItem(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mockDB := Database{db}

	mock.ExpectExec("INSERT INTO todos").
		WithArgs("laundry").
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := mockDB.AddItem("laundry"); err != nil {
		t.Errorf("error not expected while updating DB")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
