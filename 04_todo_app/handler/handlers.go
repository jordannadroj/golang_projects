package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jordannadroj/52_in_52/04_todo_app/db"
	log "github.com/sirupsen/logrus"
)

type todo struct {
	Item string
}

type HttpHandler struct {
	database *db.Database
}

func NewHttpHandler(database *db.Database) *HttpHandler {
	return &HttpHandler{
		database: database,
	}
}

func (h *HttpHandler) IndexHandler(c *fiber.Ctx) error {
	var res string
	var list []string

	todos, err := h.database.ListItems(res, list)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.Render("index", fiber.Map{
		"Todos": todos, // the key "Todos" is what we will use in the html as our access to the todos array
	})
}

// there is a form in the html file with action POST
func (h *HttpHandler) PostHandler(c *fiber.Ctx) error {
	newTodo := todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured %v", err)
		return c.SendString(err.Error())
	}

	if newTodo.Item != "" {
		err := h.database.AddItem(newTodo.Item)
		if err != nil {
			log.Printf("An error occured while executing query: %v", err)
		}
		log.Printf("Row added to database 'todos' with value %q", newTodo.Item)
	}

	return c.Redirect("/")
}

func (h *HttpHandler) PutHandler(c *fiber.Ctx) error {
	oldItem := c.Query("olditem")
	newItem := c.Query("newitem")
	err := h.database.UpdateItem(oldItem, newItem)
	if err != nil {
		log.Errorf("An error occured while executing query: %v", err)
		return c.SendString(err.Error())
	}
	log.Infof("Item %q updated to %q", oldItem, newItem)
	return nil
}

func (h *HttpHandler) DeleteHandler(c *fiber.Ctx) error {
	deleteItem := c.Query("item")
	err := h.database.DeleteItem(deleteItem)
	if err != nil {
		log.Errorf("An error occured while executing query: %v", err)
		return c.SendString(err.Error())
	}
	log.Infof("%v deleted", deleteItem)
	return nil
}
