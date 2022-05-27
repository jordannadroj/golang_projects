package store

import (
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const longUrl = "https://github.com/jordannadroj/52_in_52/tree/main/03_url_shortener"
const shortUrl = "Jsz4k57oAX"

func TestStoreInit(t *testing.T) {
	s := miniredis.RunT(t)
	assert.True(t, s != nil)
}

func TestSaveUrlMapping(t *testing.T) {

	tests := []struct {
		name        string
		initialLink string
		shortURL    string
		expect      string
		wantErr     bool
	}{
		{
			name:        "correctly saves the url mapping",
			initialLink: longUrl,
			shortURL:    shortUrl,
			expect:      longUrl,
			wantErr:     false,
		},
		{
			name:        "returns error when an invalid short url is tried to save",
			initialLink: longUrl,
			shortURL:    "",
			expect:      "",
			wantErr:     true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			miniRedis := miniredis.RunT(t)
			redisClient := redis.NewClient(&redis.Options{
				Addr:     miniRedis.Addr(),
				Password: "",
				DB:       0,
			})
			storageService := StorageService{redisClient: redisClient}

			// Persist data mapping
			err := storageService.SaveUrlMapping(tc.shortURL, tc.initialLink)

			got, errRedis := miniRedis.Get(tc.shortURL)
			assert.Equal(t, tc.expect, got)

			if tc.wantErr {
				require.Error(t, errRedis)
				require.Error(t, err)
			} else {
				require.NoError(t, errRedis)
				require.NoError(t, err)
			}
		})
	}
}

func TestRetrieveInitialUrl(t *testing.T) {
	t.Helper()
	tests := []struct {
		name        string
		initialLink string
		shortURL    string
		expect      string
		wantErr     bool
	}{
		{
			name:        "works",
			initialLink: longUrl,
			shortURL:    shortUrl,
			expect:      longUrl,
			wantErr:     false,
		},
		{
			name:        "error",
			initialLink: longUrl,
			shortURL:    shortUrl,
			expect:      "",
			wantErr:     true,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			miniRedis := miniredis.RunT(t)
			redisClient := redis.NewClient(&redis.Options{
				Addr:     miniRedis.Addr(),
				Password: "",
				DB:       0,
			})
			storageService := StorageService{redisClient: redisClient}

			miniRedis.Set(tc.shortURL, tc.initialLink)

			var got string
			var err error

			if !tc.wantErr {
				got, err = storageService.RetrieveInitialUrl(tc.shortURL)
			} else {
				got, err = storageService.RetrieveInitialUrl("")
			}

			assert.Equal(t, tc.expect, got)

			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

//
//func TestRetrieveInitialUrl(t *testing.T) {
//	initialLink := longUrl
//	shortURL := "Jsz4k57oAX"
//	miniRedis := miniredis.RunT(t)
//	redisClient := redis.NewClient(&redis.Options{
//		Addr:     miniRedis.Addr(),
//		Password: "",
//		DB:       0,
//	})
//	storageService := StorageService{redisClient: redisClient}
//	errRedis := miniRedis.Set(shortURL, initialLink)
//	assert.NoError(t, errRedis)
//
//	retrievedUrl, err := storageService.RetrieveInitialUrl(shortURL)
//
//	assert.Equal(t, initialLink, retrievedUrl)
//	assert.NoError(t, err)
//}
