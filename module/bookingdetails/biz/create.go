package bookingdetailbiz

import (
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
		condition map[string]interface{},
		ids []int,
	) ([]common.DTORoom, error)
}

type ListRoomBookedRepo interface {
	GetRoomIdsBooked(ctx context.Context, booking *common.DTOBooking) ([]int, error)
}

type createBookingDetailBiz struct {
	store        CreateBookingDetailStorage
	bookingStore FindBookingStore
	roomStore    FindRoomStore
	repo         ListRoomBookedRepo
	ps           pubsub.Pubsub
}

func NewCreateBookingDetailBiz(store CreateBookingDetailStorage, bookingStore FindBookingStore, roomStore FindRoomStore, repo ListRoomBookedRepo, ps pubsub.Pubsub) *createBookingDetailBiz {
	return &createBookingDetailBiz{store: store, bookingStore: bookingStore, roomStore: roomStore, repo: repo, ps: ps}
}

func (biz *createBookingDetailBiz) CreateBookingDetail(ctx context.Context, data *bookingdetailmodel.BookingDetailRequest) error {
	booking, err := biz.bookingStore.FindDTOWithCondition(ctx, map[string]interface{}{"id": data.BookingId})
	if err != nil {
		return common.ErrEntityNotFound("Booking", err)
	}
	// Validate booking
	if err := data.CheckInvalidBooking(booking); err != nil {

		return common.ErrCannotCreateEntity(bookingdetailmodel.EntityName, err)
	}

	// Validate room
	rooms, err := biz.roomStore.FindRoomsDTOByIds(ctx, nil, data.RoomIds)
	bookedRooms, err := biz.repo.GetRoomIdsBooked(ctx, booking)
	if err != nil {

		return common.ErrCannotCreateEntity(bookingdetailmodel.EntityName, err)
	}
	if err := data.CheckInvalidRoom(rooms, bookedRooms, booking.RoomTypeId); err != nil {
		return common.ErrCannotCreateEntity(bookingdetailmodel.EntityName, err)
	}

	// Get old data
	oldIds, err := biz.store.ListRoomOfBooking(ctx, data.BookingId)
	if err != nil {

		return common.ErrEntityNotFound("BookingDetail", err)
	}

	// Create booking detail
	if err := biz.store.Create(ctx, data.ToCreate(), oldIds); err != nil {

		return common.ErrCannotCreateEntity(bookingdetailmodel.EntityName, err)
	}

	return nil
}
