package main

import (
	"os"
	"time"

	"github.com/go-redis/redis"
)

//InitRecommendationCache starts cache
func InitRecommendationCache() (Caches, error) {
	var connection = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_DB"),
		Password: "",
		DB:       0,
	})
	err := connection.Ping().Err()
	if err != nil {
		return nil, err
	}
	return &RecommendationCache{connection: connection}, nil
}

//RecommendationCache struct
type RecommendationCache struct {
	//Recommendation
	connection *redis.Client
}

//GetProducts mehod
func (c *RecommendationCache) GetProducts(id string) ([]string, error) {

	test := make([]string, 0)

	product, err := c.connection.Get(id).Result()

	test = append(test, product)

	return test, err
}

//ForceCache method
func (c *RecommendationCache) ForceCache(id string, value string) error {

	//Forces the regeneration of a product Cache
	err := c.connection.Set(id, value, 24*time.Hour).Err()
	if err != nil {
		return err
	}

	return nil

}

func (c *RecommendationCache) cacheExists(id string) (bool, error) {

	exists := true

	_, err := c.connection.Get(id).Result()

	if err != nil {
		exists = false
	}

	return exists, err

}

func (c *RecommendationCache) getCache(id string) (string, error) {

	return c.connection.Get(id).Result()

}

func (c *RecommendationCache) updateCache(id string, products []string) error {
	return nil

}
