package db

import (
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

	//mockDB.AddItem("test_item")
	//got, _ := db.Exec("SELECT item FROM todos WHERE ITEM='test_item'")
	//assert.Equal(t, "test_item", got)
	mock.ExpectQuery("SELECT * FROM todos").WillReturnRows(sqlmock.NewRows([]string{"item"}))

	var res string
	var list []string

	resp, err := mockDB.ListItems(res, list)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp))
}
