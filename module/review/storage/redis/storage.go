package reviewstorage

import "github.com/redis/go-redis/v9"

type store struct {
	redis *redis.Client
}

func NewRedisStore(redis *redis.Client) *store {
	return &store{
		redis: redis,
	}
}
