package hotelfacilitiesbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelfacilitymodel "h5travelotobackend/module/hotelfacilities/model"
)

type ListHotelFacilitiesRepo interface {
	ListAllHotelFacilityType(ctx context.Context) ([]hotelfacilitymodel.HotelFacilityType, error)
}

type listHotelFacilitiesBiz struct {
	repo ListHotelFacilitiesRepo
}

func NewListHotelFacilities(repo ListHotelFacilitiesRepo) *listHotelFacilitiesBiz {
	return &listHotelFacilitiesBiz{repo: repo}
}

func (biz *listHotelFacilitiesBiz) ListAllHotelFacilities(ctx context.Context) ([]hotelfacilitymodel.HotelFacilityType, error) {
	data, err := biz.repo.ListAllHotelFacilityType(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(hotelfacilitymodel.EntityName, err)
	}

	return data, nil
}
