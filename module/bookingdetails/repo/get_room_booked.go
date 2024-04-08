package bookingdetailrepo

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
	"time"
)

type ListBookingDetailStore interface {
	ListRoomIdsBooked(ctx context.Context,
		startDate *time.Time,
		endDate *time.Time,
		condition map[string]interface{}) ([]int, error)
}

type getRoomBookedRepo struct {
	store ListBookingDetailStore
}

func NewGetRoomBookedRepo(store ListBookingDetailStore) *getRoomBookedRepo {
	return &getRoomBookedRepo{store: store}
}

func (repo *getRoomBookedRepo) GetRoomIdsBooked(
	ctx context.Context,
	startDate *time.Time,
	endDate *time.Time,
	condition map[string]interface{}) ([]int, error) {
	ids, err := repo.store.ListRoomIdsBooked(ctx, startDate, endDate, condition)
	if err != nil {
		return nil, common.ErrCannotListEntity(bookingdetailmodel.EntityName, err)
	}
	return ids, nil
}
