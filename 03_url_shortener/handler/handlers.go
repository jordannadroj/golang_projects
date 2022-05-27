package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"url_shortener/pkg/shortener"
	"url_shortener/store"
)

// Request model definition
// the body in the POST request should follow this format. CreateShortUrl will extract the body values based on these keys.
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"` // binding is specific to gin
}

type HttpHandler struct {
	storageService *store.StorageService
}

func NewHttpHandler(storage *store.StorageService) *HttpHandler {
	return &HttpHandler{
		storageService: storage,
	}
}

func (h *HttpHandler) CreateShortUrl(c *gin.Context) {
	var creationRequest UrlCreationRequest
	// handle error
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Infof("error binding JSON: %q", err.Error())
		// adding a return will stop the function from continuing
		return
	}

	shortUrl, err := shortener.GenerateShortLink(creationRequest.LongUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Infof("error generating short URL: %q", err.Error())
		return
	}
	err = h.storageService.SaveUrlMapping(shortUrl, creationRequest.LongUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Infof("error saving URL to database: %q", err.Error())
		return
	}

	host := "http://localhost:9808/"
	c.JSON(http.StatusOK, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortUrl,
	})
}

func (h *HttpHandler) HandleShortUrlRedirect(c *gin.Context) {
	// c.Param will extract the param with the given key
	// ex. /:shortUrl -> colon indicates param
	shortUrl := c.Param("shortUrl")
	initialUrl, err := h.storageService.RetrieveInitialUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Infof("error retrieving URL from database: %q", err.Error())
		return
	}
	// retrieve the original URL via the short URL key
	// redirects to the path of the original url
	c.Redirect(http.StatusFound, initialUrl)
}
