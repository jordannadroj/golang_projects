package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"url_shortener/handler"
	"url_shortener/store"
)

func main() {
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

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
