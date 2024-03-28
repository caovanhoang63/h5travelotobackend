package ginbookingdetail

import (
	"github.com/gin-gonic/gin"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingdetailbiz "h5travelotobackend/module/bookingdetails/biz"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
	bookingdetailrepo "h5travelotobackend/module/bookingdetails/repo"
	bookingdetailstorage "h5travelotobackend/module/bookingdetails/storage"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	roomstorage "h5travelotobackend/module/rooms/storage"
	"log"
	"net/http"
)

func CreateBookingDetails(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("booking-id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		bookingStore := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
		roomStore := roomstorage.NewSqlStore(appCtx.GetGormDbConnection())
		bookingDetailStore := bookingdetailstorage.NewSqlStore(appCtx.GetGormDbConnection())
		repo := bookingdetailrepo.NewGetRoomBookedRepo(bookingDetailStore)
		ps := appCtx.GetPubSub()
		biz := bookingdetailbiz.NewCreateBookingDetailBiz(bookingDetailStore, bookingStore, roomStore, repo, ps)

		var data bookingdetailmodel.BookingDetailRequest

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		log.Println("data: ", data.RoomFakeIds)

		data.BookingId = int(uid.GetLocalID())

		data.UnMask()
		if err := biz.CreateBookingDetail(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
