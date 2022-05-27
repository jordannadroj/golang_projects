package store

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"time"
)

type Config struct {
	Host     string `envconfig:"HOST" default:"localhost:6379"`
	Password string `envconfig:"PASSWORD" default:""`
	DB       int    `envconfig:"DB" default:"0"`
}

/*
We will start by setting up our wrappers around Redis, the wrappers will be used as interface for persisting and retrieving our application data mapping.
*/

// Define the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Note that in a real world usage, the cache duration shouldn't have
// an expiration time, an LRU policy config should be set where the
// values that are retrieved less often are purged automatically from
// the cache and stored back in RDBMS whenever the cache is full

const CacheDuration = 6 * time.Hour

// Initializing the store service and return a store pointer
func InitializeStore(cfg Config) *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     cfg.Host,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	//fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	log.Infof("Redis started successfully: pong message = {%s}", pong)
	return &StorageService{redisClient: redisClient}
}

/* We want to be able to save the mapping between the originalUrl
and the generated shortUrl url
*/

func (s *StorageService) SaveUrlMapping(shortUrl string, originalUrl string) error {
	//redis.SET store the url in cache memory
	// we set the short url as the key and the original url as the value. This allows us to query for the original url using the short url
	/*
			To access the redis memory
		 	cd /usr/local/bin/redis-cli
			KEYS *
			get <short_url>
	*/
	if shortUrl == "" {
		return errors.New("short URL cannot be empty string")
	}

	err := s.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		return fmt.Errorf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl)
	}
	return nil
}

/*
We should be able to retrieve the initial long URL once the short
is provided. This is when users will be calling the shortlink in the
url, so what we need to do here is to retrieve the long url and
think about redirect.
*/

func (s *StorageService) RetrieveInitialUrl(shortUrl string) (string, error) {
	result, err := s.redisClient.Get(shortUrl).Result()
	if err != nil {
		return "", fmt.Errorf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl)
	}
	return result, nil
}
