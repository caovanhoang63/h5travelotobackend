package ginhoteldetail

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	hoteldetailbiz "h5travelotobackend/module/hoteldetails/biz"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
	hoteldetailstorage "h5travelotobackend/module/hoteldetails/storage"
	"net/http"
)

func CreateHotelDetail(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		hotelUid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data hoteldetailmodel.HotelDetail
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		data.HotelId = int(hotelUid.GetLocalID())
		store := hoteldetailstorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := hoteldetailbiz.NewCreateHotelDetailBiz(store)
		if err := biz.CreateHotelDetail(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
