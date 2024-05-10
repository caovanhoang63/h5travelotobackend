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

func CreateHotelType(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data hoteltypemodel.HotelTypeCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := hoteltypestorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hoteltypebiz.NewCreateHotelTypeBiz(store)

		if err := biz.CreateHotelType(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
