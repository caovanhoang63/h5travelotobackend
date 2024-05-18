package ginhtcollection

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	htcollectionbiz "h5travelotobackend/module/htcollection/biz"
	htcollectionstore "h5travelotobackend/module/htcollection/store"
	"net/http"
)

func FindCollectionById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("collection-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := htcollectionstore.NewStore(appCtx.GetGormDbConnection())
		biz := htcollectionbiz.NewFincCollectionBiz(store)
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data, err := biz.FindCollectionById(c.Request.Context(), int(uid.GetLocalID()), requester)
		if err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
