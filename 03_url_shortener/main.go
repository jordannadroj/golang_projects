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

	// Note that store initialization happens here
	store := store.InitializeStore(store.Config{})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the URL Shortener API",
		})
	})

	// NOTE
	/*
		In more complex apps, endpoints should be in separate files.
	*/

	r.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c, store)
	})

	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c, store)
	})

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}

}
