package cacher

import (
	"golang.org/x/net/context"
	"time"
)

type IntCacher interface {
	Incr(ctx context.Context, key string) (int64, error)
	IncrBy(ctx context.Context, key string, value int64) (int64, error)
	Decr(ctx context.Context, key string) (int64, error)
	DecrBy(ctx context.Context, key string, value int64) (int64, error)
}

type ArrayCacher interface {
	LPush(ctx context.Context, key string, value interface{}) error
	RPush(ctx context.Context, key string, value interface{}) error
	LPop(ctx context.Context, key string) (interface{}, error)
	RPop(ctx context.Context, key string) (interface{}, error)
	LRange(ctx context.Context, key string, start, stop int64) ([]interface{}, error)
	LTrim(ctx context.Context, key string, start, stop int64) error
	Append(ctx context.Context, key string, value interface{}) error
}

type StringCacher interface {
	Append(ctx context.Context, key string, value string) error
}

type HashCacher interface {
	HSet(ctx context.Context, key, field string, value interface{}) error
	HGet(ctx context.Context, key, field string) (interface{}, error)
	HDel(ctx context.Context, key, field string) error
	HExists(ctx context.Context, key, field string) (bool, error)
	HKeys(ctx context.Context, key string) ([]string, error)
	HVals(ctx context.Context, key string) ([]interface{}, error)
	HGetAll(ctx context.Context, key string) (map[string]interface{}, error)
	HLen(ctx context.Context, key string) (int64, error)
	HIncrBy(ctx context.Context, key, field string, value int64) (int64, error)
	HScan(ctx context.Context, key string, cursor uint64, match string, count int64) (uint64, []string, error)
}

type Cacher interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
	Keys(ctx context.Context) ([]string, error)
	Flush(ctx context.Context) error
}

var (
	ErrKeyNotFound = "key not found"
)
