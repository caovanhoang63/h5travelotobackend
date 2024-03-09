package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
)

func RoleRequired(ctx appContext.AppContext, allowRoles ...string) func(c *gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser).(common.Requester)

		for _, role := range allowRoles {
			if u.GetRole() == role {
				c.Set(common.CurrentUser, u)
				c.Next()
				return
			}
		}

		panic(common.ErrNoPermission(errors.New("no permission")))
	}
}
