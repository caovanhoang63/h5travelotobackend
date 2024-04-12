package ginhotelfacilities

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hotelfacilitiesbiz "h5travelotobackend/module/hotelfacilities/biz"
	hotelfacilitysqlstore "h5travelotobackend/module/hotelfacilities/storage/sqlstore"
	"net/http"
)

// url pattern: /hotels/:hotel-id/facilities

func GetFacilitiesOfAHotel(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := hotelfacilitysqlstore.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hotelfacilitiesbiz.NewListFaciOfHotelBiz(store)

		data, err := biz.ListFacilitiesOfHotel(c.Request.Context(), int(uid.GetLocalID()))
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
