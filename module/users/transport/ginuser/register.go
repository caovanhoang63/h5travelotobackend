package ginuser

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/hasher"
	userbiz "h5travelotobackend/module/users/business"
	usermodel "h5travelotobackend/module/users/model"
	userstorage "h5travelotobackend/module/users/storage"
)

func RegisterUser(appctx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//a, _ := c.GetRawData()
		//fmt.Println(string(a))
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appctx.GetGormDbConnection()

		store := userstorage.NewSqlStore(db)

		sha256Hasher := hasher.NewSha256Hash()

		biz := userbiz.NewRegisterBiz(store, sha256Hasher)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(200, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
