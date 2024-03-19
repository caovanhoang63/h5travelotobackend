package main

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/middleware"
	"h5travelotobackend/module/bookings/transport/ginbooking"
	gindistrict "h5travelotobackend/module/districts/transport/gindistricts"
	"h5travelotobackend/module/hotels/transport/ginhotel"
	"h5travelotobackend/module/provinces/transport/ginprovinces"
	"h5travelotobackend/module/rooms/transport/ginroom"
	"h5travelotobackend/module/roomtypeabout/transport/ginroomtypeabout"
	"h5travelotobackend/module/roomtypes/transport/ginroomtype"
	"h5travelotobackend/module/upload/transport/ginupload"
	"h5travelotobackend/module/users/transport/ginuser"
	"h5travelotobackend/module/ward/transport/ginward"
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

	// room type api
	roomTypes := v1.Group("/")
	roomTypes.POST("hotels/:hotel-id/room-types", ginroomtype.CreateRoomType(appCtx))
	roomTypes.DELETE("hotels/:hotel-id/room-types/:room-type-id", ginroomtype.DeleteRoomType(appCtx))
	roomTypes.PATCH("hotels/:hotel-id/room-types/:room-type-id", ginroomtype.UpdateRoomType(appCtx))
	roomTypes.GET("/room-types/:room-type-id", ginroomtype.GetRoomTypeById(appCtx))
	roomTypes.GET("/room-types", ginroomtype.ListRoomType(appCtx))

	// hotel api
	hotels.POST("/", middleware.RoleRequired(appCtx, common.RoleOwner), ginhotel.CreateHotel(appCtx))
	hotels.GET("/:hotel-id", ginhotel.GetHotelById(appCtx))
	hotels.GET("/list", ginhotel.ListHotel(appCtx))
	hotels.DELETE("/:hotel-id", middleware.RoleRequired(appCtx, common.RoleOwner), middleware.IsHotelWorker(appCtx), ginhotel.DeleteHotel(appCtx))
	hotels.PATCH("/:hotel-id", middleware.RoleRequired(appCtx, common.RoleOwner, common.RoleManager), middleware.IsHotelWorker(appCtx), ginhotel.UpdateHotel(appCtx))
	hotels.GET("/:hotel-id/additional", ginhotel.GetHotelAdditionalInfoById(appCtx))
	hotels.PATCH("/:hotel-id/additional", ginhotel.UpdateHotelAdditionalInfo(appCtx))

	// room api
	rooms := v1.Group("hotels/:hotel-id/rooms")
	rooms.Use(middleware.RequireAuth(appCtx))
	rooms.Use(middleware.RoleRequired(appCtx, common.RoleAdmin, common.RoleOwner, common.RoleManager, common.RoleStaff))
	rooms.Use(middleware.IsHotelWorker(appCtx))

	rooms.PATCH("/:room-id", ginroom.UpdateRoom(appCtx))
	rooms.DELETE("/:room-id", ginroom.DeleteRoom(appCtx))
	rooms.POST("/", ginroom.CreateRoom(appCtx))
	rooms.GET(":room-id", ginroom.GetRoomById(appCtx))
	rooms.GET("", ginroom.ListRoomWithCondition(appCtx))

	// room type about api
	roomTypeAbout := v1.Group("hotels/:hotel-id/room-types/:room-type-id/about")
	roomTypeAbout.Use(middleware.RequireAuth(appCtx))
	roomTypeAbout.Use(middleware.RoleRequired(appCtx, common.RoleAdmin, common.RoleOwner, common.RoleManager))
	roomTypeAbout.Use(middleware.IsHotelWorker(appCtx))

	roomTypeAbout.POST("/", ginroomtypeabout.CreateRoomTypeAbout(appCtx))
	roomTypeAbout.DELETE("/", ginroomtypeabout.DeleteByRoomTypeId(appCtx))
	roomTypeAbout.PATCH("/", ginroomtypeabout.UpdateByRoomTypeId(appCtx))
	v1.GET("/hotels/:hotel-id/room-types/:room-type-id/about", ginroomtypeabout.GetAboutByRoomTypeId(appCtx))

	// VietNam Unit api
	v1.GET("/provinces", ginprovinces.ListAllProvinces(appCtx))
	v1.GET("/provinces/:province-code/districts", gindistrict.ListDistrictsByProvinceCode(appCtx))
	v1.GET("/districts/:district-code/wards", ginward.ListWardsByDistrictCode(appCtx))

	// Booking api
	booking := v1.Group("bookings/", middleware.RequireAuth(appCtx))
	booking.POST("/", ginbooking.CreateBooking(appCtx))
	booking.GET("/:booking-id", ginbooking.GetBookingById(appCtx))

}
