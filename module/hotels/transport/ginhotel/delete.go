package ginhotel

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelbiz "h5travelotobackend/module/hotels/biz"
	hotelstorage "h5travelotobackend/module/hotels/storage"
)

func DeleteHotel(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := hotelstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hotelbiz.NewDeleteHotelBiz(store, appCtx.GetPubSub())
		if err := biz.DeleteHotel(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(uid))

	}
}
