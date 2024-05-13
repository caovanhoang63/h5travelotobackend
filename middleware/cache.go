package middleware

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"time"
)

func CacheMiddleware(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := appCtx.GetLogger()
		key := common.GetApiCacheKey(c.Request.URL.String())
		cacher := appCtx.GetCacher()
		if cacher == nil {
			c.Next()
			return
		}
		log.Println("key", key)
		data, err := cacher.Get(c.Request.Context(), key)
		if err == nil {
			c.Data(200, "application/json", []byte(data))
			c.Abort()
			return
		}
		c.Next()

		response, ok := c.Get("response")
		if ok {
			if err = cacher.Set(c.Request.Context(), key, response, 10*time.Minute); err != nil {
			}
		}
	}
}
