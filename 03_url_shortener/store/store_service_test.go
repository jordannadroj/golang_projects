package store

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreInit(t *testing.T) {
	s := miniredis.RunT(t)
	assert.True(t, s != nil)
}

//TODO: refactor
func TestSaveUrlMapping(t *testing.T) {
	initialLink := "https://github.com/jordannadroj/52_in_52/tree/main/03_url_shortener"
	shortURL := "Jsz4k57oAX"
	miniRedis := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     miniRedis.Addr(),
		Password: "",
		DB:       0,
	})
	storeService := StorageService{redisClient: redisClient}

	// Persist data mapping
	err := SaveUrlMapping(shortURL, initialLink, &storeService)

	// Retrieve initial URL
	got, _ := miniRedis.Get(shortURL)

	assert.Equal(t, initialLink, got)
	assert.NoError(t, err)
}

func TestSaveUrlMapping_error(t *testing.T) {
	initialLink := "https://github.com/jordannadroj/52_in_52/tree/main/03_url_shortener"
	shortURL := ""
	miniRedis := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     miniRedis.Addr(),
		Password: "",
		DB:       0,
	})
	storeService := StorageService{redisClient: redisClient}

	// Persist data mapping
	err := SaveUrlMapping(shortURL, initialLink, &storeService)

	// Retrieve initial URL
	got, _ := miniRedis.Get(shortURL)

	assert.Error(t, err)
	assert.Equal(t, "", got)

}

func TestRetrieveInitialUrl(t *testing.T) {
	initialLink := "https://github.com/jordannadroj/52_in_52/tree/main/03_url_shortener"
	shortURL := "Jsz4k57oAX"
	miniRedis := miniredis.RunT(t)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     miniRedis.Addr(),
		Password: "",
		DB:       0,
	})
	storeService := StorageService{redisClient: redisClient}
	miniRedis.Set(shortURL, initialLink)

	retrievedUrl, _ := RetrieveInitialUrl(shortURL, &storeService)

	assert.Equal(t, initialLink, retrievedUrl)
}
