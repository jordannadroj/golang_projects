package db

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
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
	defer mockDB.sqlDB.Close()

	rows := mock.NewRows([]string{"id", "title"}).
		AddRow(1, "laundry").AddRow(2, "second todo")
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	resp, err := mockDB.ListItems()

	for _, item := range resp {
		fmt.Println(item)
	}

	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 2, len(resp))
}

func TestAddIem(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mockDB := Database{db}
	defer mockDB.sqlDB.Close()

	mock.ExpectExec("INSERT INTO todos").WithArgs("laundry").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := mockDB.AddItem("laundry")

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestAddIemError(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mockDB := Database{db}
	defer mockDB.sqlDB.Close()

	mock.ExpectExec("INSERT INTO todos").WithArgs("laundry").
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := mockDB.AddItem("laundry")

	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestUpdateItem(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mockDB := Database{db}
	defer mockDB.sqlDB.Close()

	_ = mock.NewRows([]string{"id", "title"}).
		AddRow(1, "laundry")

	mock.ExpectExec("UPDATE todos").WithArgs("laundry edit", "1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := mockDB.UpdateItem("1", "laundry edit")

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeleteItem(t *testing.T) {
	db, mock, _ := sqlmock.New()
	mockDB := Database{db}
	defer mockDB.sqlDB.Close()

	_ = mock.NewRows([]string{"id", "title"}).
		AddRow(1, "laundry").AddRow(2, "second todo")

	mock.ExpectExec("DELETE FROM todos").WithArgs("1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := mockDB.DeleteItem("1")

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
