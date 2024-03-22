package ginhoteltype

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hoteltypebiz "h5travelotobackend/module/hoteltypes/biz"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
	hoteltypestorage "h5travelotobackend/module/hoteltypes/storage"
	"net/http"
)

func UpdateHotelType(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data hoteltypemodel.HotelTypeUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		uid, err := common.FromBase58(c.Param("hotel-type"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := hoteltypestorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hoteltypebiz.NewUpdateHotelTypeBiz(store)

		if err := biz.UpdateHotelType(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
