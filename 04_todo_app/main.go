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

func main() {
	//load .env file
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	db := database.ConnectToDB(&database.Config{})
	defer db.Close()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app.Static("/public", "./public")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return handler.IndexHandler(ctx, db)
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
		return handler.PostHandler(ctx, db)
	})

	app.Put("/update", func(ctx *fiber.Ctx) error {
		return handler.PutHandler(ctx, db)
	})

	app.Delete("/delete", func(ctx *fiber.Ctx) error {
		handler.DeleteHandler(ctx, db)
		return nil
	})

	// log.Fatalln will log the output in case of any errors.
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
