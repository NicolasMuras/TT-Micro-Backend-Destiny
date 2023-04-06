package cache

import (
	"context"
	"fmt"
	"os"
	"time"

	"log"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	client *redis.Client
}

func ConnectRedis() RedisCache {
	sanityCheckRedisVars()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	redisCache := RedisCache{
		client: redisClient,
	}
	log.Println("Connected to Redis")
	return redisCache
}

func sanityCheckRedisVars() {
	envProps := []string{
		"REDIS_HOST",
		"REDIS_PORT",
		"REDIS_PASSWORD",
	}
	for _, k := range envProps {
		if os.Getenv(k) == "" {
			log.Fatal(fmt.Sprintf("Environment variable %s not defined. Terminating application...", k))
		}
	}
}

func (r *RedisCache) Set(ctx context.Context, key string, data string, expiration time.Duration) error {
	return r.client.SetNX(ctx, key, data, expiration).Err()
}

func (r *RedisCache) Get(ctx context.Context, key string) ([]byte, error) {
	result, err := r.client.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return nil, nil
	}
	return result, err
}
