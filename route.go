package main

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/middleware"
	"h5travelotobackend/module/bookingdetails/transport/ginbookingdetail"
	"h5travelotobackend/module/bookings/transport/ginbooking"
	"h5travelotobackend/module/bookingtracking/transport/ginbookingtracking"
	gindistrict "h5travelotobackend/module/districts/transport/gindistricts"
	"h5travelotobackend/module/hoteldetails/transport/ginhoteldetail"
	"h5travelotobackend/module/hotelfacilities/transport/ginhotelfacilities"
	"h5travelotobackend/module/hotels/transport/ginhotel"
	"h5travelotobackend/module/hoteltypes/transport/ginhoteltype"
	"h5travelotobackend/module/provinces/transport/ginprovinces"
	"h5travelotobackend/module/review/transport/ginreview"
	ginroomfacilities "h5travelotobackend/module/roomfacilities/transport/ginhotelfacilities"
	"h5travelotobackend/module/rooms/transport/ginroom"
	"h5travelotobackend/module/roomtypeabout/transport/ginroomtypeabout"
	"h5travelotobackend/module/roomtypes/transport/ginroomtype"
	"h5travelotobackend/module/upload/transport/ginupload"
	"h5travelotobackend/module/users/transport/ginuser"
	"h5travelotobackend/module/ward/transport/ginward"
	"h5travelotobackend/module/worker/transport/ginworker"
)

func SetUpRoute(appCtx appContext.AppContext, v1 *gin.RouterGroup) {
	v1.POST("/upload", ginupload.UploadImage(appCtx))
	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/authenticate", ginuser.UserLogin(appCtx))
	v1.POST("/renew-token", ginuser.RenewToken(appCtx))
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

	// hotel detail api
	hotelDetails := v1.Group("/hotels/:hotel-id/detail")
	hotelDetails.GET("/", ginhoteldetail.GetHotelDetailById(appCtx))

	// hotel facilities api
	hotelFacilities := v1.Group("/hotels/facilities")
	hotelFacilities.GET("/", ginhotelfacilities.ListAllHotelFacilities(appCtx))

	// room facilities api
	roomFacilities := v1.Group("/rooms/facilities")
	roomFacilities.GET("/", ginroomfacilities.ListAllRoomFacilities(appCtx))

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
	booking.DELETE("/:booking-id", ginbooking.DeleteBookingById(appCtx))
	v1.GET("/users/:user-id/bookings", middleware.RequireAuth(appCtx),
		ginbooking.ListBookingByUserId(appCtx))
	v1.GET("/hotels/:hotel-id/bookings", middleware.RequireAuth(appCtx),
		middleware.RoleRequired(appCtx, common.RoleAdmin, common.RoleOwner, common.RoleManager), middleware.IsHotelWorker(appCtx),
		ginbooking.ListBookingHotelId(appCtx))

	// worker api
	worker := v1.Group("hotels/:hotel-id/workers", middleware.RequireAuth(appCtx))
	worker.Use(middleware.RoleRequired(appCtx, common.RoleAdmin, common.RoleOwner, common.RoleManager))
	worker.Use(middleware.IsHotelWorker(appCtx))
	worker.POST("/", ginworker.CreateWorker(appCtx))
	worker.DELETE("/:user-id", ginworker.DeleteWorker(appCtx))
	worker.GET("/", ginworker.ListHotelWorker(appCtx))

	// tracking api
	trackings := v1.Group("hotels/:hotel-id/bookings/:booking-id/tracking", middleware.RequireAuth(appCtx))
	trackings.GET("/", ginbookingtracking.GetBookingTrackingState(appCtx))
	trackings.PATCH("/", ginbookingtracking.UpdateTrackingState(appCtx))

	// hotel type api
	hoteltypes := v1.Group("/hotel-types")
	hoteltypes.Use(middleware.RequireAuth(appCtx))
	hoteltypes.POST("/", middleware.RoleRequired(appCtx, common.RoleAdmin), ginhoteltype.CreateHotelType(appCtx))
	hoteltypes.DELETE("/:hotel-type", middleware.RoleRequired(appCtx, common.RoleAdmin), ginhoteltype.DeleteHotelType(appCtx))
	hoteltypes.PATCH("/:hotel-type", middleware.RoleRequired(appCtx, common.RoleAdmin), ginhoteltype.UpdateHotelType(appCtx))
	hoteltypes.GET("/:hotel-type", ginhoteltype.FindHotelTypeById(appCtx))
	hoteltypes.GET("/", ginhoteltype.ListAllHotelTypes(appCtx))

	// booking detail api
	bookingdetail := v1.Group("hotels/:hotel-id/bookings/:booking-id")
	bookingdetail.Use(middleware.RequireAuth(appCtx))
	bookingdetail.Use(middleware.RoleRequired(appCtx, common.RoleAdmin, common.RoleOwner, common.RoleManager, common.RoleStaff))
	bookingdetail.Use(middleware.IsHotelWorker(appCtx))

	bookingdetail.POST("/details", ginbookingdetail.CreateBookingDetails(appCtx))
	bookingdetail.GET("/available-rooms", ginroom.GetAvailableRoom(appCtx))

	// review api
	reviews := v1.Group("/reviews")
	reviews.Use(middleware.RequireAuth(appCtx))
	reviews.POST("/", ginreview.CreateReview(appCtx))
	reviews.GET("", ginreview.ListReviews(appCtx))
}
