package roomfacilitiesbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	roomfacilitymodel "h5travelotobackend/module/roomfacilities/model"
)

type ListFaciOfRoomStore interface {
	ListFacilitiesOfRoomType(ctx context.Context, hotelId int) ([]roomfacilitymodel.RoomFacility, error)
}

type listFaciOfRoomBiz struct {
	store ListFaciOfRoomStore
}

func NewListFaciOfRoomBiz(store ListFaciOfRoomStore) *listFaciOfRoomBiz {
	return &listFaciOfRoomBiz{store: store}
}

func (biz *listFaciOfRoomBiz) ListFacilitiesOfRoom(ctx context.Context, hotelId int) ([]roomfacilitymodel.RoomFacility, error) {
	data, err := biz.store.ListFacilitiesOfRoomType(ctx, hotelId)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return data, nil
}
