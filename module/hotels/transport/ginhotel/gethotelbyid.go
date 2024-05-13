package ginhotel

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelbiz "h5travelotobackend/module/hotels/biz"
	hotelstorage "h5travelotobackend/module/hotels/storage"
)

func GetHotelById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := hotelstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hotelbiz.NewFindHotelBiz(store, appCtx.GetPubSub())
		data, err := biz.FindWithConditionHotel(c.Request.Context(), map[string]interface{}{"id": int(uid.GetLocalID())})
		if err != nil {
			panic(err)
		}
		data.Mask(false)

		c.JSON(200, common.SimpleSuccessResponse(data))
		c.Set("response", data)
	}
}
