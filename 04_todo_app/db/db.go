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
func (db *Database) ListItems(response string, list []string) ([]string, error) {
	rows, err := db.SqlDB.Query("SELECT * FROM todos")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		return list, errors.New("error retrieving items from DB")
	}
	for rows.Next() { // read each row
		rows.Scan(&response)          // Scan() copies the row into a dedicated pointer variable
		list = append(list, response) // append each row to the todos array
	}
	return list, nil
}

func (db *Database) AddItem(item string) error {
	_, err := db.SqlDB.Exec("INSERT INTO todos VALUES ($1)", item)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateItem(oldItem, newItem string) error {
	_, err := db.SqlDB.Exec("UPDATE todos SET item=$1 WHERE item=$2", newItem, oldItem)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) DeleteItem(item string) error {
	_, err := db.SqlDB.Exec("DELETE FROM todos WHERE item=$1", item)
	if err != nil {
		return err
	}
	return nil
}
