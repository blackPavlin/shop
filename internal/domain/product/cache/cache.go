package cache

import "github.com/redis/go-redis/v9"

// Cache
type Cache struct {
	client *redis.Client
}

// NewCache
func NewCache(client *redis.Client) *Cache {
	return &Cache{client: client}
}
