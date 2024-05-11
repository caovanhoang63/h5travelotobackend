package bookingbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
)

type CountBookedRoomStore interface {
	CountBookedRoom(ctx context.Context, rtId int,
		startDate *common.CivilDate, endDate *common.CivilDate) (*int, error)
}

type countBookedRoomBiz struct {
	store CountBookedRoomStore
}

func NewCountBookedRoomBiz(store CountBookedRoomStore) *countBookedRoomBiz {
	return &countBookedRoomBiz{store: store}
}

func (biz *countBookedRoomBiz) CountBookedRoom(ctx context.Context, rtId int,
	startDate *common.CivilDate, endDate *common.CivilDate) (*int, error) {
	result, err := biz.store.CountBookedRoom(ctx, rtId, startDate, endDate)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return result, nil
}
