package ginuser

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	userbiz "h5travelotobackend/module/users/business"
	usermodel "h5travelotobackend/module/users/model"
	userstorage "h5travelotobackend/module/users/storage"
)

func ChangePassword(appCtx appContext.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		requester := c.MustGet(common.CurrentUser).(common.Requester)

		var data usermodel.UserChangePassword

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSqlStore(appCtx.GetGormDbConnection())
		biz := userbiz.NewUserChangePasswordBiz(store)

		if err := biz.ChangePassword(c.Request.Context(), requester.GetUserId(), &data); err != nil {
			panic(err)
		}

		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
