package ginhtcollection

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	htcollectionbiz "h5travelotobackend/module/htcollection/biz"
	htcollection "h5travelotobackend/module/htcollection/model"
	htcollectionstore "h5travelotobackend/module/htcollection/store"
	"net/http"
)

// CreateCollection

func CreateCollection(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data htcollection.HotelCollectionCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := htcollectionstore.NewStore(appCtx.GetGormDbConnection())
		biz := htcollectionbiz.NewCreateCollectionBiz(store)
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()

		if err := biz.CreateCollection(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}

}
