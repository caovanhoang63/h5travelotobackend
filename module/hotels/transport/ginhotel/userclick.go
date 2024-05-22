package ginhotel

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelbiz "h5travelotobackend/module/hotels/biz"
	"h5travelotobackend/module/hotels/storage/hotelrdbstore"
	"log"
)

func UserClick(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		uid, _ := common.FromBase58(c.Param("hotel-id"))
		log.Println("User click hotel", uid.GetLocalID())
		store := hotelrdbstore.NewStore(appCtx.GetRedisClient())

		biz := hotelbiz.NewUserClickHotelBiz(store)

		err := biz.AddToUserRecentlyViewedHotel(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID()))
		if err != nil {
			log.Println(err)
		}

	}
}
