package handler

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type todo struct {
	Item string
}

func IndexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var todos []string

	rows, err := db.Query("SELECT * FROM todos") // return all rows that are returned by query
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() { // read each row
		rows.Scan(&res)            // Scan() copies the row into a dedicated pointer variable
		todos = append(todos, res) // append each row to the todos array
	}
	return c.Render("index", fiber.Map{
		"Todos": todos, // the key "Todos" is what we will use in the html as our access to the todos array
	})
}

// there is a form in the html file with action POST
func PostHandler(c *fiber.Ctx, db *sql.DB) error {
	//	add a new todo to the list in the db
	//	render the todos again with the list
	newTodo := todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured %v", err)
		return c.SendString(err.Error())
	}

	if newTodo.Item != "" {
		_, err := db.Exec("INSERT INTO todos VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Printf("An error occured while executing query: %v", err)
		}
		log.Printf("Row added to database 'todos' with value %q", newTodo.Item)
	}

	return c.Redirect("/")
}

func PutHandler(c *fiber.Ctx, db *sql.DB) error {
	olditem := c.Query("olditem")
	newitem := c.Query("newitem")
	_, err := db.Exec("UPDATE todos SET item=$1 WHERE item=$2", newitem, olditem)
	if err != nil {
		log.Errorf("An error occured while executing query: %v", err)
		return c.SendString(err.Error())
	}
	log.Infof("Item %q updated to %q", olditem, newitem)
	return nil
}

func DeleteHandler(c *fiber.Ctx, db *sql.DB) error {
	deleteItem := c.Query("item")
	_, err := db.Exec("DELETE FROM todos WHERE item=$1", deleteItem)
	if err != nil {
		log.Errorf("An error occured while executing query: %v", err)
		return c.SendString(err.Error())
	}
	log.Infof("%v deleted", deleteItem)
	return nil
}
