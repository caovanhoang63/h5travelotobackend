package ginuser

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	userbiz "h5travelotobackend/module/users/business"
	usermodel "h5travelotobackend/module/users/model"
)

func CheckPinPassword(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Pin struct {
			Pin string `json:"pin" form:"pin" binding:"required"`
		}
		var pin Pin
		err := c.ShouldBind(&pin)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		email := c.Param("email")
		if email == "" {
			panic(common.ErrInvalidRequest(usermodel.ErrInvalidEmail))
		}

		biz := userbiz.NewCheckPinPasswordBiz(appCtx.GetCacher())
		err = biz.CheckPinPassword(c.Request.Context(), email, pin.Pin)
		if err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
