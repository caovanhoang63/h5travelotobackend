package bookingbiz

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	"h5travelotobackend/module/bookings/bookingmodel"
	"log"
)

type CreateBookingStore interface {
	Create(ctx context.Context, data *bookingmodel.BookingCreate) error
}

type FindRoomTypeStore interface {
	FindDTODataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*common.DTORoomType, error)
}

type createBookingBiz struct {
	store             CreateBookingStore
	findRoomTypeStore FindRoomTypeStore
	pb                pubsub.Pubsub
}

func NewCreateBookingBiz(store CreateBookingStore, typeStore FindRoomTypeStore, pb pubsub.Pubsub) *createBookingBiz {
	return &createBookingBiz{store: store, findRoomTypeStore: typeStore, pb: pb}
}

func (biz *createBookingBiz) Create(
	ctx context.Context,
	data *bookingmodel.BookingCreate,
	requester common.Requester) error {

	roomType, err := biz.findRoomTypeStore.FindDTODataWithCondition(ctx, map[string]interface{}{"id": data.RoomTypeId})
	if err != nil {
		if err == common.RecordNotFound {
			return bookingmodel.ErrInvalidRoomType
		}
		return common.ErrCannotCreateEntity(bookingmodel.EntityName, err)
	}

	if roomType.Status == 0 || roomType.HotelId != data.HotelId {
		return bookingmodel.ErrInvalidRoomType
	}
	err = biz.store.Create(ctx, data)
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
