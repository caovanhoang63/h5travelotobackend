package subcriber

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	"h5travelotobackend/email"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	hoteldetailsqlstorage "h5travelotobackend/module/hoteldetails/storage"
	roomtypesqlstorage "h5travelotobackend/module/roomtypes/storage/sqlstorage"
	userstorage "h5travelotobackend/module/users/storage"
	hotelstorage "h5travelotobackend/search/module/hotel/storage/esstore"
	"log"
	"strconv"
)

func SendConfirmMail(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "SendConfirmMail",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var data common.PaymentBooking
			err := json.Unmarshal(message.Data, &data)
			if err != nil {
				return err
			}

			bkStore := bookingsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
			booking, err := bkStore.FindWithCondition(ctx, map[string]interface{}{"id": data.BookingId})
			if err != nil {
				return err
			}
			booking.Mask(false)

			hStore := hotelstorage.NewESStore(appCtx.GetElasticSearchClient())
			hotel, err := hStore.FindHotelById(ctx, booking.HotelId)
			if err != nil {
				return err
			}

			hdStore := hoteldetailsqlstorage.NewSqlStore(appCtx.GetGormDbConnection())
			detail, err := hdStore.FindWithCondition(ctx, map[string]interface{}{"hotel_id": hotel.Id})
			if err != nil {
				return err
			}

			uStore := userstorage.NewSqlStore(appCtx.GetGormDbConnection())
			user, err := uStore.FindUser(ctx, map[string]interface{}{"id": booking.UserId})
			if err != nil {
				return err
			}

			rStore := roomtypesqlstorage.NewSqlStore(appCtx.GetGormDbConnection())

			roomType, err := rStore.FindDataWithCondition(ctx, map[string]interface{}{"id": booking.RoomTypeId})
			if err != nil {
				return err
			}
			log.Println(roomType.Name)

			var confirmE = common.ConfirmBooking{
				HotelName:    hotel.Name,
				RoomQuantity: booking.RoomQuantity,
				RoomName:     roomType.Name,
				StartDate:    booking.StartDate.String(),
				EndDate:      booking.EndDate.String(),
				CheckInTime:  detail.CheckInTime,
				Phone:        hotel.Hotline,
				CustomerName: fmt.Sprintf("%s %s", user.Firstname, user.LastName),
				BookingId:    booking.FakeId.String(),
				Amount:       strconv.FormatFloat(data.Amount, 'f', -1, 64),
				Addr:         fmt.Sprintf("%s, %s, %s, %s", hotel.Address, hotel.Ward.FullName, hotel.District.FullName, hotel.Province.Name),
				PaymentType:  "(Đã thanh toán)",
			}
			mail := email.NewConfirmEmail(user.Email, confirmE)

			e := appCtx.GetSendMailEngine()
			e.Send(*mail)
			return nil
		},
	}
}
