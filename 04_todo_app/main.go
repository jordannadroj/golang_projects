package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"github.com/joho/godotenv"
	database "github.com/jordannadroj/52_in_52/04_todo_app/db"
	"github.com/jordannadroj/52_in_52/04_todo_app/handler"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
)

func StartApp(handler handler.HttpHandler) *fiber.App {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/public", "./public")

	app.Get("/", handler.IndexHandler)

	app.Post("/api/todo", handler.PostHandler)

	app.Put("/api/todo", handler.PutHandler)

	app.Delete("/api/todo", handler.DeleteHandler)

	return app
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	db := database.ConnectToDB(&database.Config{})
	defer db.CloseDB()

	httpHandler := handler.NewHttpHandler(db)

	app := StartApp(*httpHandler)

	// log.Fatalln will log the output in case of any errors.
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
