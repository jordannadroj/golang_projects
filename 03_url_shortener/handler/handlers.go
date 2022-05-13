package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url_shortener/pkg/shortener"
	"url_shortener/store"
)

// Request model definition
// the body in the POST request should follow this format. CreateShortUrl will extract the body values based on these keys.
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"` // binding is specific to gin
}

func CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	// handle error
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl)

	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func HandleShortUrlRedirect(c *gin.Context) {
	// c.Param will extract the param with the given key
	// ex. /:shortUrl -> colon indicates param
	shortUrl := c.Param("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	// retrieve the original URL via the short URL key
	// redirects to the path of the original url
	c.Redirect(302, initialUrl)
}
