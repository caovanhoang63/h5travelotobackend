package middleware

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
)

func DelCacheMiddleware(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		key := c.GetString(common.CacheKey)
		if key == "" {
			return
		}
		cacher := appCtx.GetCacher()
		if cacher == nil {
			c.Next()
			return
		}
		ok, err := cacher.Exists(c.Request.Context(), key)
		if !ok || err != nil {
			c.Next()
			return
		}
		err = cacher.Del(c.Request.Context(), key)
	}
}
