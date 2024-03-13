package main

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/middleware"
	"h5travelotobackend/module/hotels/transport/ginhotel"
	"h5travelotobackend/module/roomtypes/transport/ginroomtype"
	"h5travelotobackend/module/upload/transport/ginupload"
	"h5travelotobackend/module/users/transport/ginuser"
)

func SetUpRoute(appCtx appContext.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/upload", ginupload.UploadImage(appCtx))
	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/authenticate", ginuser.UserLogin(appCtx))

	users := v1.Group("/users", middleware.RequireAuth(appCtx))

	users.GET("/profile", ginuser.GetProfile(appCtx))
	users.PATCH("/profile", ginuser.Update(appCtx))
	users.PATCH("/change-password", ginuser.ChangePassword(appCtx))

	hotels := v1.Group("hotels", middleware.RequireAuth(appCtx))
	hotels.POST("/", ginhotel.CreateHotel(appCtx))
	hotels.GET("/:id", ginhotel.GetHotelById(appCtx))
	hotels.GET("/:id/additional", ginhotel.GetHotelAdditionalInfoById(appCtx))
	hotels.GET("/list", ginhotel.ListHotel(appCtx))
	hotels.PATCH("/:id", ginhotel.UpdateHotel(appCtx))
	hotels.PATCH("/:id/additional", ginhotel.UpdateHotelAdditionalInfo(appCtx))

	hotelRoomTypes := hotels.Group("/:hotel-id")
	hotelRoomTypes.DELETE("/", ginhotel.DeleteHotel(appCtx))

	hotelRoomTypes.POST("/room-types", middleware.CheckWorkerRole(appCtx, common.RoleManager, common.RoleOwner), ginroomtype.CreateRoomType(appCtx))
	hotelRoomTypes.DELETE("/room-types/:room-type-id", middleware.CheckWorkerRole(appCtx, common.RoleManager, common.RoleOwner), ginroomtype.DeleteRoomType(appCtx))

	roomTypes := v1.Group("room-types")
	roomTypes.GET("/:id", ginroomtype.GetRoomTypeById(appCtx))
}
