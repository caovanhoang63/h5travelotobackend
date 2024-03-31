package ginreview

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	reviewbiz "h5travelotobackend/module/review/biz"
	reviewmodel "h5travelotobackend/module/review/model"
	reviewstorage "h5travelotobackend/module/review/storage"
	"net/http"
)

func ListReviews(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter reviewmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		filter.UnMask()
		paging.FullFill()

		store := reviewstorage.NewMongoStore(appCtx.GetMongoConnection())
		biz := reviewbiz.NewListReviewsBiz(store)
		data, err := biz.ListReviews(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}

}
