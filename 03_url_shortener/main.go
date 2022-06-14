package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jordannadroj/52_in_52/03_url_shortener/handler"
	"github.com/jordannadroj/52_in_52/03_url_shortener/store"
	"net/http"
	"os"
)

func main() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		panic("Error loading .env file")
	}

	r := gin.Default()

	// Store initialization
	storage := store.InitializeStore(store.Config{})

	// now the httphandler has the redis instance inside it
	httpHandler := handler.NewHttpHandler(storage)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	// NOTE
	/*
		In more complex apps, endpoints should be in separate files.
	*/

	r.POST("/create-short-url", httpHandler.CreateShortUrl)

	r.GET("/:shortUrl", httpHandler.HandleShortUrlRedirect)

	err = r.Run(os.Getenv("APP_HOST"))
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
