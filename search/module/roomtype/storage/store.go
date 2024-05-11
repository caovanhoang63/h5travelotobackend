package rtsearchstorage

import (
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
)

type store struct {
	es  *elasticsearch.TypedClient
	rdb *redis.Client
}

func NewStore(es *elasticsearch.TypedClient, rdb *redis.Client) *store {
	return &store{es: es, rdb: rdb}
}
