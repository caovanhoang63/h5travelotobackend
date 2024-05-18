package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"log"
	"net/http"
	"time"
)

func RateLimitingById(appCtx appContext.AppContext, rateLimit int64, expiration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		api := c.FullPath()
		id := requester.GetUserId()
		limitKey := common.GenRateLimitKeyById(id, api)
		err := incRequestCount(c.Request.Context(), appCtx.GetRedisClient(), limitKey, rateLimit, expiration)
		if err != nil {
			if errors.Is(err, common.RateLimited) {
				err = banUser(c.Request.Context(),
					appCtx.GetRedisClient(),
					common.GenBanUserIdKey(id),
					30*time.Minute)
				if err != nil {
					log.Printf("Error banning user id: %v", err)
				}
				c.AbortWithStatusJSON(http.StatusTooManyRequests, common.ErrTooManyRequest(api, err))
			}
		}
		c.Next()
	}
}

func RateLimitingByIp(appCtx appContext.AppContext, threshHold int64, expiration time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIp := c.ClientIP()
		api := c.FullPath()
		key := common.GenRateLimitKeyByIp(clientIp, api)
		err := incRequestCount(c.Request.Context(), appCtx.GetRedisClient(), key, threshHold, expiration)
		if err != nil {
			if errors.Is(err, common.RateLimited) {
				err = banUser(c.Request.Context(), appCtx.GetRedisClient(), common.GenBanUserIpKey(clientIp), 30*time.Second)
				if err != nil {
					log.Printf("Error banning user ip: %v", err)
				}
				c.AbortWithStatusJSON(http.StatusTooManyRequests, common.ErrTooManyRequest(api, common.RateLimited))
				return
			}
		}
		c.Next()
	}
}

func banUser(ctx context.Context, rdb *redis.Client, key string, expiration time.Duration) error {
	err := rdb.SetNX(ctx, key, 1, expiration).Err()
	if err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func incRequestCount(ctx context.Context,
	rdb *redis.Client,
	key string,
	threshHold int64,
	expiration time.Duration) error {
	err := rdb.Watch(ctx, func(tx *redis.Tx) error {
		_ = tx.SetNX(ctx, key, 0, expiration)
		count, err := tx.Incr(ctx, key).Result()
		if err != nil {
			return common.ErrDb(err)
		}
		if count > threshHold {
			return common.RateLimited
		}
		return nil
	}, key)

	return err
}
