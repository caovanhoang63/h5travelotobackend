package rediscacher

import (
	"github.com/redis/go-redis/v9"
	"h5travelotobackend/component/logger"
	"time"
)

type RedisCacher struct {
	redisClient *redis.Client
	keyPrefix   string
	marshal     func(value interface{}) ([]byte, error)
	unmarshal   func(b []byte, value interface{}) error
	ttl         time.Duration
	logger      logger.Logger
}

func NewRedisCacher(redisClient *redis.Client,
	keyPrefix string,
	marshal func(value interface{}) ([]byte, error),
	unmarshal func(b []byte, value interface{}) error,
	ttl time.Duration, logger logger.Logger) *RedisCacher {
	return &RedisCacher{
		redisClient: redisClient,
		keyPrefix:   keyPrefix,
		marshal:     marshal,
		unmarshal:   unmarshal,
		ttl:         ttl,
		logger:      logger,
	}
}
