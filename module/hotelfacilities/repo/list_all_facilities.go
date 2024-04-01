package hotelfacilityrepo

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelfacilitymodel "h5travelotobackend/module/hotelfacilities/model"
)

type ListFacilityTypesStore interface {
	ListAllHotelFacilityType(ctx context.Context) ([]hotelfacilitymodel.HotelFacilityType, error)
}

type ListFacilitiesStore interface {
	ListHotelFacilityByType(ctx context.Context, typeId int) ([]hotelfacilitymodel.HotelFacility, error)
}

type listFacilityTypesRepo struct {
	facilityStore     ListFacilitiesStore
	facilityTypeStore ListFacilityTypesStore
}

func NewListFacilityTypesRepo(facilityStore ListFacilitiesStore, facilityTypeStore ListFacilityTypesStore) *listFacilityTypesRepo {
	return &listFacilityTypesRepo{facilityStore: facilityStore, facilityTypeStore: facilityTypeStore}
}

func (repo *listFacilityTypesRepo) ListAllHotelFacilityType(ctx context.Context) ([]hotelfacilitymodel.HotelFacilityType, error) {
	types, err := repo.facilityTypeStore.ListAllHotelFacilityType(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(hotelfacilitymodel.EntityName, err)
	}

	for i := range types {
		types[i].Facilities, err = repo.facilityStore.ListHotelFacilityByType(ctx, types[i].Id)
		if err != nil {
			return nil, common.ErrCannotListEntity(hotelfacilitymodel.EntityName, err)
		}
	}

	return types, nil
}
