package cache

import (
	"accounts/config/env"
	"context"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     env.REDIS_URL,
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
