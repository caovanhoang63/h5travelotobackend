package ginhoteltype

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hoteltypebiz "h5travelotobackend/module/hoteltypes/biz"
	hoteltypestorage "h5travelotobackend/module/hoteltypes/storage"
	"net/http"
)

func FindHotelTypeById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		uid, err := common.FromBase58(c.Param("hotel-type"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := hoteltypestorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hoteltypebiz.NewFindHotelTypeBiz(store)

		data, err := biz.FindHotelTypeById(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
