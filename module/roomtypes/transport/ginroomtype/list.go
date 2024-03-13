package ginroomtype

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	roomtypebiz "h5travelotobackend/module/roomtypes/biz"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
)

func ListRoomType(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter roomtypemodel.Filter
		var paging common.Paging
		var data []roomtypemodel.RoomType

		if err := c.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		paging.FullFill()

		if err := c.ShouldBindQuery(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if c.Param("max-price") == "" && c.Param("min-price") == "" {
			filter.SetDefault()
		}

		store := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := roomtypebiz.NewListRoomTypeBiz(store)
		data, err := biz.ListRoomTypeWithCondition(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range data {
			data[i].Mask(false)
		}

		c.JSON(200, common.NewSuccessResponse(data, paging, filter))
	}
}
