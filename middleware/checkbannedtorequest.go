package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
)

func CheckBannedToRequest(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		var requester common.Requester
		var id *int
		if cur, ok := c.Get(common.CurrentUser); ok {
			requester = cur.(common.Requester)
			*id = requester.GetUserId()
		} else {
			id = nil
		}
		if ok, _ := checkBanned(c.Request.Context(), appCtx.GetRedisClient(), id, ip); ok {
			c.AbortWithStatusJSON(429, common.ErrTooManyRequest("server", common.RateLimited))
			return
		}
		c.Next()
	}
}

func checkBanned(ctx context.Context, rdb *redis.Client, id *int, ip string) (bool, error) {
	var val int64
	var err error
	if id == nil {
		cmd := rdb.Exists(ctx, common.GenBanUserIpKey(ip))
		val, err = cmd.Result()
	} else {
		cmd := rdb.Exists(ctx, common.GenBanUserIdKey(*id), common.GenBanUserIpKey(ip))
		val, err = cmd.Result()
	}

	if err != nil {
		return false, common.ErrDb(err)
	}
	return val > 0, nil

}
