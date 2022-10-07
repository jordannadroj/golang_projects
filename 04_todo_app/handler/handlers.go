package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jordannadroj/52_in_52/04_todo_app/db"
	log "github.com/sirupsen/logrus"
)

type todo struct {
	Item string
}

type updateTodo struct {
	OldItem string `json:"olditem"`
	NewItem string `json:"newitem"`
}

type HttpHandler struct {
	database db.TodoDatabase
}

// NewHttpHandler takes the TodoDatabase interface to create a handler with any supported db takes the TodoDatabse interface to create a handler with any supported db
func NewHttpHandler(database db.TodoDatabase) *HttpHandler {
	return &HttpHandler{
		database: database,
	}
}

func (h *HttpHandler) IndexHandler(c *fiber.Ctx) error {
	todos, err := h.database.ListItems()
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
	updateItem := updateTodo{}
	fmt.Println(string(c.Body()))
	if err := json.Unmarshal(c.Body(), &updateItem); err != nil {
		log.Printf("An error occured %v", err)
		return c.SendString(err.Error())
	}
	err := h.database.UpdateItem(updateItem.OldItem, updateItem.NewItem)
	if err != nil {
		log.Errorf("An error occured while executing query: %v", err)
		return c.SendString(err.Error())
	}
	log.Infof("Item %q updated to %q", updateItem.OldItem, updateItem.NewItem)
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
