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

func UpdateCollection(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("collection-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var update htcollection.HotelCollectionUpdate

		if err = c.ShouldBind(&update); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := htcollectionstore.NewStore(appCtx.GetGormDbConnection())
		biz := htcollectionbiz.NewUpdateCollectionBiz(store)
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err = biz.UpdateCollection(c.Request.Context(), int(uid.GetLocalID()), &update, requester); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
