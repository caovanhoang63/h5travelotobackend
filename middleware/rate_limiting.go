package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"net/http"
	"time"
)

func RateLimiting(appCtx appContext.AppContext, rateLimit, second int64) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		api := c.FullPath()
		key := common.GenRateLimitKey(clientIp, api)

		err := incRequestCount(c.Request.Context(), appCtx.GetRedisClient(), key, rateLimit, second)
		if err != nil {
			if errors.Is(err, common.RateLimited) {
				c.AbortWithStatusJSON(http.StatusTooManyRequests, common.ErrTooManyRequest(clientIp, api, err))

			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInternal(err))
		}

		c.Next()
	}

}

//func banUserIp(ctx context.Context, rdb *redis.Client, clientIp string, second int64) error {
//	//key := common.GenBanUserKey(clientIp)
//	err := rdb.Set(ctx, key, 1, time.Duration(second)*time.Second).Err()
//	if err != nil {
//		return common.ErrDb(err)
//	}
//	return nil
//}

func incRequestCount(ctx context.Context,
	rdb *redis.Client,
	key string,
	rateLimit int64,
	second int64) error {
	err := rdb.Watch(ctx, func(tx *redis.Tx) error {
		_ = tx.SetNX(ctx, key, 0, time.Duration(second)*time.Second)
		count, err := tx.Incr(ctx, key).Result()
		if err != nil {
			return common.ErrDb(err)
		}
		if count > rateLimit {
			err = common.RateLimited
		}
		return nil
	}, key)

	return err
}
