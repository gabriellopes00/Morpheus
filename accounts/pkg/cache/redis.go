package cache

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type redisCacheRepository struct {
	client *redis.Client
}

func NewRedisCacheRepository(client *redis.Client) *redisCacheRepository {
	return &redisCacheRepository{
		client: client,
	}
}

func (r *redisCacheRepository) Set(key string, value string, exp time.Duration) error {
	err := r.client.Set(r.client.Context(), key, value, exp).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisCacheRepository) Get(key string) (string, error) {
	result := r.client.Get(r.client.Context(), key)
	if result.Err() != nil {
		return "", result.Err()
	}

	return result.Val(), nil
}

func (r *redisCacheRepository) Delete(key string) error {
	err := r.client.Del(r.client.Context(), key).Err()
	if err != nil {
		return err
	}

	return nil
}
