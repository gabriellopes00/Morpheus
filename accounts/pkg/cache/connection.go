package cache

import (
	"accounts/config/env"
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", env.REDIS_HOST, env.REDIS_PORT),
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
