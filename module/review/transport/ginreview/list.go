package ginreview

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	reviewbiz "h5travelotobackend/module/review/biz"
	reviewmodel "h5travelotobackend/module/review/model"
	reviewrepo "h5travelotobackend/module/review/repo"
	reviewstorage "h5travelotobackend/module/review/storage/mongo"
	userstorage "h5travelotobackend/module/users/storage"
	"net/http"
)

func ListReviews(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Called")
		var filter reviewmodel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := filter.UnMask(); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()

		store := reviewstorage.NewMongoStore(appCtx.GetMongoConnection())
		userStore := userstorage.NewSqlStore(appCtx.GetGormDbConnection())
		repo := reviewrepo.NewListReviewsRepo(store, userStore)
		biz := reviewbiz.NewListReviewsBiz(repo)
		data, err := biz.ListReviews(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
			data[i].FixTime()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(data, paging, filter))
	}
}
