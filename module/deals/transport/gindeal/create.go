package gindeal

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	dealbiz "h5travelotobackend/module/deals/biz"
	dealmodel "h5travelotobackend/module/deals/model"
	dealsqlstorage "h5travelotobackend/module/deals/storage"
	"net/http"
)

func CreateDeal(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var deal dealmodel.DealCreate
		if err := c.ShouldBind(&deal); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		hotelUid, err := common.FromBase58(c.Param("hotel-id"))
		deal.UnMask()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		deal.HotelId = int(hotelUid.GetLocalID())

		store := dealsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := dealbiz.NewCreateDealBiz(store)
		if biz.CreateDeal(c.Request.Context(), &deal); err != nil {
			panic(err)
		}

		deal.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(deal.FakeId))

	}
}
