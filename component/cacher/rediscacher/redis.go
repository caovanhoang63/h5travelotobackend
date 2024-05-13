package rediscacher

import (
	"golang.org/x/net/context"
	"time"
)

func (r *RedisCacher) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	set := r.redisClient.Set(ctx, key, value, expiration)
	if set.Err() != nil {
		r.logger.Error("Falied to set key: ",
			key, " with value: ",
			value, " and expiration: ",
			expiration, " with error: ",
			set.Err())
	}
	r.logger.Println("Set key: ", key, " with value: ", value, " and expiration: ", expiration)
	return set.Err()
}
func (r *RedisCacher) Get(ctx context.Context, key string) (interface{}, error) {
	get := r.redisClient.Get(ctx, key)
	if get.Err() != nil {
		r.logger.Error("Falied to get key: ", key, " with error: ", get.Err())
	}
	r.logger.Debug("Get key: ", key, " with value: ", get.Val())
	return get.Val(), get.Err()

}
func (r *RedisCacher) Del(ctx context.Context, key string) error {
	del := r.redisClient.Del(ctx, key)
	if del.Err() != nil {
		r.logger.Error("Falied to delete key: ", key, " with error: ", del.Err())
	}
	r.logger.Println("Deleted key: ", key)
	return del.Err()
}
func (r *RedisCacher) Exists(ctx context.Context, key string) (bool, error) {
	exists := r.redisClient.Exists(ctx, key)
	if exists.Err() != nil {
		r.logger.Error("Falied to check key: ", key, " with error: ", exists.Err())
	}
	r.logger.Debug("Check key: ", key, " exists: ", exists.Val())
	return exists.Val() == 1, exists.Err()
}
func (r *RedisCacher) Keys(ctx context.Context) ([]string, error) {
	keys := r.redisClient.Keys(ctx, r.keyPrefix+"*")
	if keys.Err() != nil {
		r.logger.Error("Falied to get keys with error: ", keys.Err())
	}
	r.logger.Debug("Get keys: ", keys.Val())
	return keys.Val(), keys.Err()
}

func (r *RedisCacher) Flush(ctx context.Context) error {
	keys := r.redisClient.Keys(ctx, r.keyPrefix+"*")
	if keys.Err() != nil {
		r.logger.Error("Falied to get keys with error: ", keys.Err())
	}
	r.logger.Debug("Get keys: ", keys.Val())
	for _, key := range keys.Val() {
		del := r.redisClient.Del(ctx, key)
		if del.Err() != nil {
			r.logger.Error("Falied to delete key: ", key, " with error: ", del.Err())
		}
		r.logger.Println("Deleted key: ", key)
	}
	return nil
}
