package ginhotel

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelbiz "h5travelotobackend/module/hotels/biz"
	hotelmodel "h5travelotobackend/module/hotels/model"
	hotelstorage "h5travelotobackend/module/hotels/storage"
)

func CreateHotel(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data hotelmodel.HotelCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := hotelstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hotelbiz.NewCreateHotelBiz(store, appCtx.GetPubSub())
		if err := biz.CreateHotel(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(200, common.SimpleSuccessResponse(data.FakeId))

	}
}
