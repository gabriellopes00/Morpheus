package cache

import "time"

type CacheRepository interface {
	Set(key string, value string, exp time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
}
