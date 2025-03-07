package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

// Why initialize WHY NOT INITIALISE??
func InitializeRedis(addr, password string, db int) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

// Get retrieves a value from Redis
func Get(key string) (string, error) {
	return rdb.Get(ctx, key).Result()
}

// Set stores a value in Redis with an expiration time
func Set(key string, value interface{}, expiration time.Duration) error {
	return rdb.Set(ctx, key, value, expiration).Err()
}

// Delete removes a key from Redis
func Delete(key string) error {
	return rdb.Del(ctx, key).Err()
}
