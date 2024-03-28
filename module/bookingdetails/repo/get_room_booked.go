package bookingdetailrepo

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
)

type ListBookingDetailStore interface {
	ListRoomIdsBooked(ctx context.Context, booking *common.DTOBooking) ([]int, error)
}

type getRoomBookedRepo struct {
	store ListBookingDetailStore
}

func NewGetRoomBookedRepo(store ListBookingDetailStore) *getRoomBookedRepo {
	return &getRoomBookedRepo{store: store}
}

func (repo *getRoomBookedRepo) GetRoomIdsBooked(ctx context.Context, booking *common.DTOBooking) ([]int, error) {
	ids, err := repo.store.ListRoomIdsBooked(ctx, booking)
	if err != nil {
		return nil, common.ErrCannotListEntity(bookingdetailmodel.EntityName, err)
	}
	return ids, nil

}
