package middleware

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"net/http"
	"time"
)

func CacheMiddleware(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := common.GenApiCacheKey(c.Request.URL.String())
		cacher := appCtx.GetCacher()
		if cacher == nil {
			c.Next()
			return
		}
		data, err := cacher.Get(c.Request.Context(), key)
		if err == nil {
			c.Data(http.StatusOK, "application/json", []byte(data))
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
