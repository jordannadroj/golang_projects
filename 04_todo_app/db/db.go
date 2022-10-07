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
	sqlDB *sql.DB
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

	return &Database{sqlDB: db}
}

func (db *Database) CloseDB() {
	db.sqlDB.Close()
}

// Queries all rows of the database and returns a list of items
func (db *Database) ListItems() ([]Todo, error) {
	var todos []Todo
	rows, err := db.sqlDB.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		return todos, errors.New("error retrieving items from DB")
	}
	for rows.Next() { // read each row
		todo := Todo{}
		rows.Scan(&todo.ID, &todo.Item) // Scan() copies the row into a dedicated pointer variable
		todos = append(todos, todo)     // append each row to the todos array
	}
	return todos, nil
}

func (db *Database) AddItem(item string) error {
	res, err := db.sqlDB.Exec("INSERT INTO todos(id,item) VALUES (DEFAULT,$1)", item)
	rows, err2 := res.RowsAffected()
	if err != nil {
		return err
	} else if err2 != nil {
		return err2
	} else if rows == 0 {
		return errors.New("error when adding item to database")
	}
	return nil
}

func (db *Database) UpdateItem(oldItem, newItem string) error {
	res, err := db.sqlDB.Exec("UPDATE todos SET item=$1 WHERE id=$2", newItem, oldItem)
	rows, err2 := res.RowsAffected()
	if err != nil {
		return err
	} else if err2 != nil {
		return err2
	} else if rows == 0 {
		return errors.New("error when updating item oin database")
	}
	return nil
}

func (db *Database) DeleteItem(itemID string) error {
	res, err := db.sqlDB.Exec("DELETE FROM todos WHERE id=$1", itemID)
	rows, err2 := res.RowsAffected()
	if err != nil {
		return err
	} else if err2 != nil {
		return err2
	} else if rows == 0 {
		return errors.New("error when deleting item from database")
	}
	return nil
}
