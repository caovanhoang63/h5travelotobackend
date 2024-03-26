package bookingdetailbiz

import (
	"fmt"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
)

type CreateBookingDetailStorage interface {
	Create(ctx context.Context, data []bookingdetailmodel.BookingDetailCreate, oldIds []int) error
	CountRoomOfBooking(ctx context.Context, bookingId int) (int, error)
	ListRoomOfBooking(ctx context.Context, bookingId int) ([]int, error)
}

type FindBookingStore interface {
	FindDTOWithCondition(ctx context.Context,
		condition map[string]interface{}) (*common.DTOBooking, error)
}

type FindRoomStore interface {
	FindRoomsDTOByIds(
		ctx context.Context,
		ids []int,
	) ([]common.DTORoom, error)
}

type createBookingDetailBiz struct {
	store        CreateBookingDetailStorage
	bookingStore FindBookingStore
	roomStore    FindRoomStore
	ps           pubsub.Pubsub
}

func NewCreateBookingDetailBiz(store CreateBookingDetailStorage, bookingStore FindBookingStore, roomStore FindRoomStore, ps pubsub.Pubsub) *createBookingDetailBiz {
	return &createBookingDetailBiz{store: store, bookingStore: bookingStore, roomStore: roomStore, ps: ps}
}

func (biz *createBookingDetailBiz) CreateBookingDetail(ctx context.Context, data *bookingdetailmodel.BookingDetailRequest) error {
	booking, err := biz.bookingStore.FindDTOWithCondition(ctx, map[string]interface{}{"id": data.BookingId})
	if err != nil {
		return common.ErrEntityNotFound("Booking", err)
	}
	// Validate booking
	if err := data.CheckInvalidBooking(booking); err != nil {
		fmt.Println("err1: ", err)
		fmt.Println("err1: ", len(data.RoomIds))
		return common.ErrCannotCreateEntity(bookingdetailmodel.EntityName, err)
	}

	// Validate room
	rooms, err := biz.roomStore.FindRoomsDTOByIds(ctx, data.RoomIds)
	if err != nil {
		fmt.Println("err2:", err)

		return common.ErrCannotCreateEntity(bookingdetailmodel.EntityName, err)
	}
	if data.CheckInvalidRoom(rooms, booking.RoomTypeId) != nil {
		fmt.Println("err3: ", err)

		return common.ErrCannotCreateEntity(bookingdetailmodel.EntityName, err)
	}

	// Get old data
	oldIds, err := biz.store.ListRoomOfBooking(ctx, data.BookingId)
	if err != nil {
		fmt.Println("err4: ", err)

		return common.ErrEntityNotFound("BookingDetail", err)
	}

	// Create booking detail
	if err := biz.store.Create(ctx, data.ToCreate(), oldIds); err != nil {
		fmt.Println("err5: ", err)

		return common.ErrCannotCreateEntity(bookingdetailmodel.EntityName, err)
	}

	return nil
}
