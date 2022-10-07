package db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/vrischmann/envconfig"
)

type Config struct {
	Host     string `envconfig:"HOST"`
	Port     int    `envconfig:"DB_PORT"`
	User     string `envconfig:"USER"`
	Password string `envconfig:"PASSWORD"`
	DB       string `envconfig:"DB_NAME" `
}

type Database struct {
	SqlDB *sql.DB
}

type Todo struct {
	ID   int
	Item string
}

// Establishes a connection to a SQL database
func ConnectToDB(cfg *Config) *Database {
	if err := envconfig.Init(cfg); err != nil {
		log.Fatalln(err)
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DB)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return &Database{SqlDB: db}
}

// Queries all rows of the database and returns a list of items
func (db *Database) ListItems() ([]Todo, error) {
	todo := Todo{}
	var todos []Todo
	rows, err := db.SqlDB.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		return todos, errors.New("error retrieving items from DB")
	}
	for rows.Next() { // read each row
		rows.Scan(&todo.ID, &todo.Item) // Scan() copies the row into a dedicated pointer variable
		todos = append(todos, todo)     // append each row to the todos array
	}
	return todos, nil
}

func (db *Database) AddItem(item string) error {
	result, err := db.SqlDB.Exec("INSERT INTO todos(id,item) VALUES (DEFAULT,$1)", item)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return errors.New("one row must have been affected")
	}
	return nil
}

func (db *Database) UpdateItem(oldItem, newItem string) error {
	_, err := db.SqlDB.Exec("UPDATE todos SET item=$1 WHERE id=$2", newItem, oldItem)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) DeleteItem(itemID string) error {
	_, err := db.SqlDB.Exec("DELETE FROM todos WHERE id=$1", itemID)
	if err != nil {
		return err
	}
	return nil
}
