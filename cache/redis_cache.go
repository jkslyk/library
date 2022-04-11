package cache

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/jkslyk/library/internal/domain"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) BookCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client	 {
	return redis.NewClient(&redis.Options{
		Addr: cache.host,
		Password: "",
		DB: cache.db,
	})
}	

func (cache *redisCache) Set(key string, value *domain.Book) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if(err != nil) {
		panic(err)
	}
	client.Set(key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *domain.Book {
	client := cache.getClient()
	val, err := client.Get(key).Result()
	if(err != nil) {
		return nil
	} 
	book := domain.Book{}
	err = json.Unmarshal([]byte(val), &book)
	if(err != nil) {
		panic(err)
	}
	return &book
}
