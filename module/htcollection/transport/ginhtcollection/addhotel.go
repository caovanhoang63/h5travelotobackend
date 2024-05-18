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

func AddHotelToCollection(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var add htcollection.HotelCollectionDetailCreate
		collectionUid, err := common.FromBase58(c.Param("collection-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		hotelUid, err := common.FromBase58(c.Param("hotel-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		add.CollectionId = int(collectionUid.GetLocalID())
		add.HotelId = int(hotelUid.GetLocalID())

		store := htcollectionstore.NewStore(appCtx.GetGormDbConnection())
		biz := htcollectionbiz.NewAddHotelBiz(store)
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		err = biz.AddHotelToCollection(c.Request.Context(), &add, requester)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
