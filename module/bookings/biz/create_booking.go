package bookingbiz

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/bookingmodel"
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
}

func NewCreateBookingBiz(store CreateBookingStore, typeStore FindRoomTypeStore) *createBookingBiz {
	return &createBookingBiz{store: store, findRoomTypeStore: typeStore}
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

	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(bookingmodel.EntityName, err)
	}
	return nil
}
