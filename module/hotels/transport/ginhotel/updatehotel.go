package ginhotel

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelbiz "h5travelotobackend/module/hotels/biz"
	hotelmodel "h5travelotobackend/module/hotels/model"
	hotelstorage "h5travelotobackend/module/hotels/storage"
)

func UpdateHotel(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data hotelmodel.HotelUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := hotelstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hotelbiz.NewUpdateHotelBiz(store, appCtx.GetPubSub())
		if err := biz.UpdateHotel(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(uid))

	}
}
