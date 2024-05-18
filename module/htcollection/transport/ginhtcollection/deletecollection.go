package ginhtcollection

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	htcollectionbiz "h5travelotobackend/module/htcollection/biz"
	htcollectionstore "h5travelotobackend/module/htcollection/store"
	"net/http"
)

func DeleteCollection(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("collection-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		store := htcollectionstore.NewStore(appCtx.GetGormDbConnection())
		biz := htcollectionbiz.NewDeleteCollectionBiz(store)
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		if err = biz.DeleteCollection(c.Request.Context(), int(uid.GetLocalID()), requester); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}

}
