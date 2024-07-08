package middleware

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	workersqlstorage "h5travelotobackend/module/worker/storage/sqlstorage"
)

func IsWorker(appCtx appContext.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)
		if user.GetRole() == common.RoleAdmin {
			c.Set(common.CurrentUser, user)
			c.Next()
			return
		}

		userId := user.GetUserId()
		db := appCtx.GetGormDbConnection()
		store := workersqlstorage.NewSqlStore(db)
		data, err := store.FindWithCondition(c.Request.Context(), map[string]interface{}{"user_id": userId})
		if err != nil {
			panic(common.ErrNoPermission(err))
		}
		c.Set(common.CurrentUser, user)
		c.Set(common.CurrentWorker, data)
		c.Next()

	}

}
