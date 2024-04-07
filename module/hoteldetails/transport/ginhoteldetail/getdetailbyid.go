package ginhoteldetail

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hoteldetailbiz "h5travelotobackend/module/hoteldetails/biz"
	hoteldetailsqlstorage "h5travelotobackend/module/hoteldetails/storage"
	"net/http"
)

//url : v1/hotels/:hotel-id/detail

func GetHotelDetailById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("hotel-id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		hotelId := int(uid.GetLocalID())

		store := hoteldetailsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := hoteldetailbiz.NewGetHotelDetailByIdBiz(store)

		data, err := biz.GetHotelDetailById(c.Request.Context(), hotelId)
		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))

	}

}
