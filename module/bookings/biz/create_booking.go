package bookingbiz

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	"h5travelotobackend/module/bookings/model"
	"log"
	"time"
)

type CreateBookingStore interface {
	Create(ctx context.Context, data *bookingmodel.BookingCreate) error
	CountBookedRoom(ctx context.Context, rtId int,
		startDate *common.CivilDate, endDate *common.CivilDate) (*int, error)
}

type FindRoomTypeStore interface {
	FindDTODataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*common.DTORoomType, error)
}

type createBookingBiz struct {
	bookingStore  CreateBookingStore
	roomTypeStore FindRoomTypeStore
	pb            pubsub.Pubsub
}

func NewCreateBookingBiz(store CreateBookingStore, typeStore FindRoomTypeStore, pb pubsub.Pubsub) *createBookingBiz {
	return &createBookingBiz{bookingStore: store, roomTypeStore: typeStore, pb: pb}
}

func (biz *createBookingBiz) Create(
	ctx context.Context,
	data *bookingmodel.BookingCreate,
	requester common.Requester) error {

	if err := ValidateBooking(data); err != nil {
		return common.ErrInvalidRequest(err)
	}

	roomType, err := biz.roomTypeStore.FindDTODataWithCondition(ctx, map[string]interface{}{"id": data.RoomTypeId})
	if err != nil {
		if err == common.RecordNotFound {
			return bookingmodel.ErrInvalidRoomType
		}
		return common.ErrCannotCreateEntity(bookingmodel.EntityName, err)
	}

	if roomType.Status == 0 || roomType.HotelId != data.HotelId {
		return bookingmodel.ErrInvalidRoomType
	}

	booked, err := biz.bookingStore.CountBookedRoom(ctx, data.RoomTypeId, data.StartDate, data.EndDate)
	if err != nil {
		return common.ErrInternal(err)
	}

	if booked == nil {
		booked = new(int)
		*booked = 0
	}

	if *booked+data.RoomQuantity > roomType.TotalRoom {
		return bookingmodel.ErrRoomNotAvailable
	}

	data.TotalAmount = roomType.Price * float64(data.RoomQuantity) * float64(data.EndDate.DateDiff(*data.StartDate))
	data.DiscountAmount = 0
	data.Currency = common.VND

	data.FinalAmount = data.TotalAmount - data.DiscountAmount

	err = biz.bookingStore.Create(ctx, data)
	if err != nil {
		return common.ErrCannotCreateEntity(bookingmodel.EntityName, err)
	}
	var dtoBooking common.DTOBooking

	dtoBooking.Id = data.Id

	err = biz.pb.Publish(ctx, common.TopicCreateBooking, pubsub.NewMessage(dtoBooking))
	if err != nil {
		log.Println(common.ErrCannotCreateEntity(bookingmodel.EntityName, err))
	}
	return nil

}

func ValidateBooking(f *bookingmodel.BookingCreate) error {
	if f.RoomQuantity <= 0 {
		return bookingmodel.ErrRoomQuantityIsZero
	}
	if f.Adults+f.Children == 0 {
		return bookingmodel.ErrOccupancyEmpty
	}
	if f.StartDate == nil {
		return bookingmodel.ErrStartIsEmpty
	}
	if f.EndDate == nil {
		return bookingmodel.ErrEndIsEmpty
	}

	if f.StartDate.After(*f.EndDate) {
		return bookingmodel.ErrStartDateAfterEndDate
	}

	now := time.Now()
	if !(f.StartDate.After(common.CivilDate(now)) || (f.StartDate.IsEqual(common.CivilDate(now)))) {
		return bookingmodel.ErrStartInPass
	}

	return nil
}
