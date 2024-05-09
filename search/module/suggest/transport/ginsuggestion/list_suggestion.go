package ginsuggestion

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	suggestbiz "h5travelotobackend/search/module/suggest/biz"
	suggestmodel "h5travelotobackend/search/module/suggest/model"
	suggeststorage "h5travelotobackend/search/module/suggest/storage"
)

// url: GET search/suggestions

func ListSuggestion(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req suggestmodel.SuggestRequest

		if err := c.ShouldBind(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if req.Limit <= 0 {
			req.Limit = 7
		}

		store := suggeststorage.NewESStore(appCtx.GetElasticSearchClient())
		biz := suggestbiz.NewListSuggestBiz(store)
		res, err := biz.ListSuggestions(c.Request.Context(), &req)

		if err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(res))
	}
}
