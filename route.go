package main

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/chat/module/message/transport/ginchatmessage"
	"h5travelotobackend/chat/module/room/transport/ginchatroom"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/middleware"
	"h5travelotobackend/module/bookingdetails/transport/ginbookingdetail"
	"h5travelotobackend/module/bookings/transport/ginbooking"
	"h5travelotobackend/module/bookingtracking/transport/ginbookingtracking"
	"h5travelotobackend/module/deals/transport/gindeal"
	gindistrict "h5travelotobackend/module/districts/transport/gindistricts"
	"h5travelotobackend/module/hoteldetails/transport/ginhoteldetail"
	"h5travelotobackend/module/hotelfacilities/transport/ginhotelfacilities"
	"h5travelotobackend/module/hotels/transport/ginhotel"
	"h5travelotobackend/module/hotelsave/transport/ginhotelsave"
	"h5travelotobackend/module/hoteltypes/transport/ginhoteltype"
	"h5travelotobackend/module/htcollection/transport/ginhtcollection"
	"h5travelotobackend/module/provinces/transport/ginprovinces"
	"h5travelotobackend/module/review/transport/ginreview"
	ginroomfacilities "h5travelotobackend/module/roomfacilities/transport/ginhotelfacilities"
	"h5travelotobackend/module/rooms/transport/ginroom"
	"h5travelotobackend/module/roomtypes/transport/ginroomtype"
	"h5travelotobackend/module/upload/transport/ginupload"
	"h5travelotobackend/module/users/transport/ginuser"
	"h5travelotobackend/module/ward/transport/ginward"
	"h5travelotobackend/module/worker/transport/ginworker"
	"h5travelotobackend/payment/module/payin/transport/ginpayin"
	"h5travelotobackend/search/module/hotel/transport/ginhotelsearch"
	"h5travelotobackend/search/module/roomtype/transport/ginrtsearch"
	"h5travelotobackend/search/module/suggest/transport/ginsuggestion"
	"time"
)

func SetUpRoute(appCtx appContext.AppContext, v1 *gin.RouterGroup) {
	v1.Use(middleware.CheckBannedToRequest(appCtx), middleware.RateLimitingByIp(appCtx, 500, time.Minute))

	v1.POST("/upload", ginupload.UploadImage(appCtx))
	v1.POST("/register", ginuser.RegisterUser(appCtx))
	v1.POST("/authenticate", ginuser.UserLogin(appCtx))
	v1.POST("/forgot-password/:email", ginuser.ConfirmResetPassword(appCtx))
	v1.PATCH("/change-password-forgot/:email", ginuser.ChangePasswordForgot(appCtx))
	v1.POST("/renew-token", ginuser.RenewToken(appCtx))
	v1.GET("users/exists", ginuser.CheckExistedEmail(appCtx))
	v1.GET("/users/check-pin/:email", ginuser.CheckPinPassword(appCtx))

	// ===================== User =====================
	usersRead := v1.Group("/users", middleware.RequireAuth(appCtx), middleware.CacheMiddleware(appCtx))
	usersRead.GET("/profile", ginuser.GetProfile(appCtx))

	usersWrite := v1.Group("/users", middleware.RequireAuth(appCtx))
	usersWrite.PATCH("/profile", ginuser.Update(appCtx))
	usersWrite.PATCH("/change-password", ginuser.ChangePassword(appCtx))
	// ===================== User =====================

	// ===================== Room Type =====================
	roomTypes := v1.Group("/")
	roomTypesWrite := roomTypes.Group("hotels/:hotel-id/room-types",
		middleware.RequireAuth(appCtx))
	roomTypesWrite.POST("", ginroomtype.CreateRoomType(appCtx))
	roomTypesWrite.DELETE("/:room-type-id", ginroomtype.DeleteRoomType(appCtx))
	roomTypesWrite.PATCH("/:room-type-id", ginroomtype.UpdateRoomType(appCtx))

	roomTypesRead := roomTypes.Group("room-types")
	roomTypesRead.GET("/:room-type-id", middleware.CacheMiddleware(appCtx),
		ginroomtype.GetRoomTypeById(appCtx))
	roomTypesRead.GET("", ginroomtype.ListRoomType(appCtx))
	// ===================== Room Type =====================

	// ===================== Hotel =====================
	hotels := v1.Group("hotels", middleware.RequireAuth(appCtx))
	hotelsRead := hotels.Group("")
	hotelsRead.GET("/current", middleware.IsWorker(appCtx), ginhotel.GetHotelByUser(appCtx))
	hotelsRead.GET("/:hotel-id", middleware.CacheMiddleware(appCtx),
		ginhotel.GetHotelById(appCtx))
	hotelsRead.GET("/list", ginhotel.ListHotel(appCtx))

	hotelsWrite := hotels.Group("",
		middleware.DelCacheMiddleware(appCtx))
	hotelsWrite.POST("/",
		middleware.RoleRequired(appCtx, common.RoleOwner),
		ginhotel.CreateHotel(appCtx))
	hotelsWrite.DELETE("/:hotel-id",
		middleware.RoleRequired(appCtx, common.RoleOwner), middleware.IsHotelWorker(appCtx),
		ginhotel.DeleteHotel(appCtx))
	hotelsWrite.PATCH("/:hotel-id",
		middleware.RoleRequired(appCtx, common.RoleOwner, common.RoleManager),
		middleware.IsHotelWorker(appCtx),
		ginhotel.UpdateHotel(appCtx))
	hotelsWrite.POST("/:hotel-id/click", ginhotel.UserClick(appCtx))
	// ===================== Hotel =====================

	// ===================== Hotel Detail =====================
	hotelDetails := v1.Group("/hotels/:hotel-id/detail")
	hotelDetails.GET("/", ginhoteldetail.GetHotelDetailById(appCtx))
	// ===================== Hotel Detail =====================

	// ===================== Hotel Facilities =====================
	hotelFacilities := v1.Group("/hotels/facilities")
	hotelFacilities.GET("/", ginhotelfacilities.ListAllHotelFacilities(appCtx))
	hotelFacilitiesWithId := v1.Group("/hotels/:hotel-id/facilities")
	hotelFacilitiesWithId.GET("/", ginhotelfacilities.GetFacilitiesOfAHotel(appCtx))
	// ===================== Hotel Facilities =====================

	// ===================== Room Facilities =====================
	roomFacilities := v1.Group("/rooms/facilities")
	roomFacilities.GET("/", ginroomfacilities.ListAllRoomFacilities(appCtx))

	v1.GET("/room-types/:room-type-id/facilities", ginroomfacilities.GetFacilitiesOfRoomType(appCtx))
	// ===================== Room Facilities =====================

	// ===================== Room =====================
	rooms := v1.Group("hotels/:hotel-id/")
	rooms.Use(middleware.RequireAuth(appCtx))
	rooms.Use(middleware.RoleRequired(appCtx, common.RoleAdmin, common.RoleOwner, common.RoleManager, common.RoleStaff))
	rooms.Use(middleware.IsHotelWorker(appCtx))

	rooms.PATCH("rooms/:room-id", ginroom.UpdateRoom(appCtx))
	rooms.DELETE("rooms/:room-id", ginroom.DeleteRoom(appCtx))
	rooms.POST("rooms/", ginroom.CreateRoom(appCtx))
	rooms.GET("rooms/:room-id", ginroom.GetRoomById(appCtx))
	rooms.GET("rooms", ginroom.ListRoomWithCondition(appCtx))

	rooms.GET("/available-rooms", ginroom.GetAvailableRoomByDate(appCtx))
	rooms.GET("rooms/bookings/:booking-id", ginroom.ListRoomOfBooking(appCtx))
	rooms.GET("rooms/bookings/:booking-id/available", ginroom.GetAvailableRoomForBooking(appCtx))
	// ===================== Room =====================

	// ===================== Vietnamese Units =====================
	v1.GET("/provinces",
		middleware.CacheMiddleware(appCtx),
		ginprovinces.ListAllProvinces(appCtx))
	v1.GET("/provinces/:province-code/districts",
		middleware.CacheMiddleware(appCtx),
		gindistrict.ListDistrictsByProvinceCode(appCtx))
	v1.GET("/districts/:district-code/wards",
		middleware.CacheMiddleware(appCtx),
		ginward.ListWardsByDistrictCode(appCtx))
	// ===================== Vietnamese Units =====================

	// ===================== Booking =====================
	booking := v1.Group("bookings/", middleware.RequireAuth(appCtx))
	booking.POST("/checkdeal/:deal-id", ginbooking.CheckDeal(appCtx))
	booking.POST("/", ginbooking.CreateBooking(appCtx))
	v1.POST("hotels/:hotel-id/bookings/front-desk",
		middleware.RequireAuth(appCtx),
		middleware.IsHotelWorker(appCtx),
		ginbooking.CreateFrontDeskBooking(appCtx))
	booking.GET("/:booking-id", ginbooking.GetBookingById(appCtx))
	booking.DELETE("/:booking-id", ginbooking.DeleteBookingById(appCtx))
	v1.GET("/users/:user-id/bookings", middleware.RequireAuth(appCtx),
		ginbooking.ListBookingByUserId(appCtx))

	bookingHotel := v1.Group("/hotels/:hotel-id/bookings", middleware.RequireAuth(appCtx),
		middleware.RoleRequired(appCtx,
			common.RoleAdmin,
			common.RoleOwner,
			common.RoleManager,
			common.RoleStaff),
		middleware.IsHotelWorker(appCtx))
	bookingHotel.GET("/", ginbooking.ListBookingHotelId(appCtx))
	bookingHotel.PATCH("/:booking-id/check-in", ginbooking.CheckInBooking(appCtx))
	bookingHotel.PATCH("/:booking-id/check-out", ginbooking.CheckOutBooking(appCtx))
	bookingHotel.PATCH("/:booking-id/cancel", ginbooking.CancelBooking(appCtx))
	bookingHotel.GET("/overview", ginbooking.OverviewByDate(appCtx))

	// ===================== Booking =====================

	// ===================== Worker =====================
	worker := v1.Group("hotels/:hotel-id/workers", middleware.RequireAuth(appCtx))
	worker.Use(middleware.RoleRequired(appCtx, common.RoleAdmin, common.RoleOwner, common.RoleManager))
	worker.Use(middleware.IsHotelWorker(appCtx))
	worker.POST("/", ginworker.CreateWorker(appCtx))
	worker.DELETE("/:user-id", ginworker.DeleteWorker(appCtx))
	worker.GET("/", ginworker.ListHotelWorker(appCtx))
	// ===================== Worker =====================

	// ===================== Booking Tracking =====================
	trackings := v1.Group("hotels/:hotel-id/bookings/:booking-id/tracking", middleware.RequireAuth(appCtx))
	trackings.GET("/", ginbookingtracking.GetStatesOfBooking(appCtx))
	// ===================== Booking Tracking =====================

	// ===================== Hotel Type =====================
	hoteltypes := v1.Group("/hotel-types")
	hoteltypes.Use(middleware.RequireAuth(appCtx))
	hoteltypes.POST("/", middleware.RoleRequired(appCtx, common.RoleAdmin), ginhoteltype.CreateHotelType(appCtx))
	hoteltypes.DELETE("/:hotel-type", middleware.RoleRequired(appCtx, common.RoleAdmin), ginhoteltype.DeleteHotelType(appCtx))
	hoteltypes.PATCH("/:hotel-type", middleware.RoleRequired(appCtx, common.RoleAdmin), ginhoteltype.UpdateHotelType(appCtx))

	hotelTypesRead := v1.Group("/hotel-types")
	hotelTypesRead.GET("/:hotel-type", ginhoteltype.FindHotelTypeById(appCtx))
	hotelTypesRead.GET("/", ginhoteltype.ListAllHotelTypes(appCtx))
	// ===================== Hotel Type =====================

	// ===================== Booking Detail =====================
	bookingdetail := v1.Group("hotels/:hotel-id/bookings/:booking-id")
	bookingdetail.Use(middleware.RequireAuth(appCtx))
	bookingdetail.Use(middleware.RoleRequired(appCtx, common.RoleAdmin, common.RoleOwner, common.RoleManager, common.RoleStaff))
	bookingdetail.Use(middleware.IsHotelWorker(appCtx))
	bookingdetail.POST("/details", ginbookingdetail.CreateBookingDetails(appCtx))
	bookingdetail.GET("rooms", ginbookingdetail.ListRoomsOfBooking(appCtx))

	// ===================== Booking Detail =====================

	// ===================== Review =====================
	reviews := v1.Group("/reviews")
	reviews.Use(middleware.RequireAuth(appCtx))
	reviews.POST("/", ginreview.CreateReview(appCtx))
	reviews.GET("", ginreview.ListReviews(appCtx))
	reviews.DELETE("/:id", ginreview.DeleteReviewById(appCtx))
	// ===================== Review =====================

	// ===================== Deal =====================
	dealWrite := v1.Group("hotels/:hotel-id/deals")
	dealWrite.Use(middleware.RequireAuth(appCtx),
		middleware.RoleRequired(appCtx, common.RoleAdmin, common.RoleOwner, common.RoleManager),
		middleware.IsHotelWorker(appCtx))
	dealWrite.POST("/", gindeal.CreateDeal(appCtx))
	dealWrite.DELETE("/:deal-id", gindeal.DeleteDealById(appCtx))
	dealWrite.PATCH("/:deal-id", gindeal.UpdateDeal(appCtx))

	dealRead := v1.Group("deals")
	dealRead.GET("/:deal-id", gindeal.GetDealById(appCtx))
	dealRead.GET("/", gindeal.ListDeal(appCtx))
	// ===================== Deal =====================

	// ===================== Chat =====================
	useChat := v1.Group("/chat")
	useChat.Use(middleware.RequireAuth(appCtx))
	useChat.GET("/hotels/:hotel-id", ginchatroom.GetChatRoom(appCtx))
	useChat.GET("/messages/:message-id", ginchatmessage.GetMessageById(appCtx))
	useChat.GET("/rooms/:room-id/messages", ginchatmessage.ListMessagesByRoomId(appCtx))
	hotelChat := v1.Group("/hotels/:hotel-id/chat")
	hotelChat.Use(middleware.RequireAuth(appCtx))
	hotelChat.Use(middleware.RoleRequired(appCtx, common.RoleAdmin, common.RoleOwner, common.RoleManager, common.RoleStaff))
	hotelChat.Use(middleware.IsHotelWorker(appCtx))
	hotelChat.GET("/", ginchatroom.ListChatRoomByHotelId(appCtx))
	v1.GET("/users/chat-rooms", middleware.RequireAuth(appCtx), ginchatroom.ListChatRoomByUser(appCtx))
	// ===================== Chat =====================

	// ===================== Search =====================
	search := v1.Group("/search")
	search.GET("/suggestions", ginsuggestion.ListSuggestion(appCtx))
	search.GET("/hotels", ginhotelsearch.ListHotel(appCtx))
	search.GET("/hotels/prominent", ginhotelsearch.ListProminentHotels(appCtx))
	search.GET("hotels/viewed", middleware.RequireAuth(appCtx), ginhotelsearch.ListRecentlyViewed(appCtx))

	rtSearch := search.Group("/room-types")
	rtSearch.GET("/", ginrtsearch.ListAvailableRoomType(appCtx))
	// ===================== Search =====================

	// ===================== Save Hotel =====================
	saveHotel := v1.Group("hotels")
	saveHotel.Use(middleware.RequireAuth(appCtx))
	saveHotel.POST("/:hotel-id/save", ginhotelsave.UserSaveHotel(appCtx))
	saveHotel.DELETE("/:hotel-id/unsave", ginhotelsave.UserUnSaveHotel(appCtx))
	saveHotel.GET("/:hotel-id/saved", ginhotelsave.IsHotelSaved(appCtx))
	saveHotel.GET("/saved", ginhotelsave.ListHotelSavedByUser(appCtx))
	// ===================== Save Hotel =====================

	// ===================== Hotel Collection =====================
	collection := v1.Group("collections")
	collection.Use(middleware.RequireAuth(appCtx))
	collection.POST("/", ginhtcollection.CreateCollection(appCtx))
	collection.GET("", ginhtcollection.ListUserCollections(appCtx))
	collection.GET("/:collection-id", ginhtcollection.FindCollectionById(appCtx))
	collection.PATCH("/:collection-id", ginhtcollection.UpdateCollection(appCtx))
	collection.DELETE("/:collection-id", ginhtcollection.DeleteCollection(appCtx))
	collection.POST("/:collection-id/hotels/:hotel-id", ginhtcollection.AddHotelToCollection(appCtx))
	collection.DELETE("/:collection-id/hotels/:hotel-id", ginhtcollection.RemoveHotelFromCollection(appCtx))
	collection.GET("/:collection-id/hotels", ginhtcollection.ListHotelInCollection(appCtx))

	// ===================== Hotel Collection =====================

	// ===================== Payment  =====================
	payment := v1.Group("payment")
	payment.Use(middleware.RequireAuth(appCtx))
	vnPay := payment.Group("vnpay")
	vnPay.GET("/pay-in", ginpayin.PayIn(appCtx))
	payment.POST("/execute/:txn_id", ginpayin.ExecutePayIn(appCtx))
	payment.POST("/cancel/:txn_id", ginpayin.CancelPayIn(appCtx))
	payment.GET("/success/:booking_id", ginpayin.GetPaymentSuccessBooking(appCtx))

	paymentIPN := v1.Group("payment")
	vnPayIPN := paymentIPN.Group("vnpay")
	vnPayIPN.GET("/ipn", ginpayin.VnpIPN(appCtx))

}
