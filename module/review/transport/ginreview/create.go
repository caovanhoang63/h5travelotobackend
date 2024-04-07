package ginreview

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	reviewbiz "h5travelotobackend/module/review/biz"
	reviewmodel "h5travelotobackend/module/review/model"
	reviewstorage "h5travelotobackend/module/review/storage"
)

// url path: /v1/reviews

func CreateReview(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data reviewmodel.Review

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := reviewstorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := reviewbiz.NewCreateReviewBiz(store, appCtx.GetPubSub())
		requester := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = requester.GetUserId()

		data.UnMask()

		if err := biz.CreateReview(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.JSON(200, common.SimpleSuccessResponse(data.ID))
	}
}
