package ginreview

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	reviewbiz "h5travelotobackend/module/review/biz"
	reviewstorage "h5travelotobackend/module/review/storage/mongo"
	"net/http"
)

func DeleteReviewById(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var objId primitive.ObjectID
		if err := objId.UnmarshalText([]byte(id)); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		store := reviewstorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := reviewbiz.NewDeleteReviewBiz(store, appCtx.GetPubSub())
		if err := biz.DeleteReview(c.Request.Context(), requester, objId); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
