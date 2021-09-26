package cache

import (
	"log"
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
	result := r.client.Set(r.client.Context(), key, value, exp)
	if result.Err() != nil {
		log.Println(result.Err())
		return result.Err()
	}

	return nil
}

func (r *redisCacheRepository) Get(key string) (string, error) {
	result := r.client.Get(r.client.Context(), key)
	if result.Err() != nil {
		log.Println(result.Err())
		return "", result.Err()
	}

	return result.String(), nil
}

func (r *redisCacheRepository) Delete(key string) error {
	result := r.client.Del(r.client.Context(), key)
	if result.Err() != nil {
		log.Println(result.Err())
		return result.Err()
	}

	return nil
}
