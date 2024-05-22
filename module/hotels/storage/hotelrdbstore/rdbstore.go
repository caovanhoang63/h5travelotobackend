package hotelrdbstore

import "github.com/redis/go-redis/v9"

type store struct {
	rdb *redis.Client
}

func NewStore(rdb *redis.Client) *store {
	return &store{
		rdb: rdb,
	}
}
