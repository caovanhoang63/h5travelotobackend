package hotelfacilitiesbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelfacilitymodel "h5travelotobackend/module/hotelfacilities/model"
)

type ListHotelFacilitiesStore interface {
	ListAllRoomFacilities(ctx context.Context) ([]hotelfacilitymodel.HotelFacility, error)
}

type listHotelFacilitiesBiz struct {
	store ListHotelFacilitiesStore
}

func NewListHotelFacilities(store ListHotelFacilitiesStore) *listHotelFacilitiesBiz {
	return &listHotelFacilitiesBiz{store: store}
}

func (biz *listHotelFacilitiesBiz) ListAllHotelFacilities(ctx context.Context) ([]hotelfacilitymodel.HotelFacility, error) {
	data, err := biz.store.ListAllRoomFacilities(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(hotelfacilitymodel.EntityName, err)
	}

	return data, nil
}
