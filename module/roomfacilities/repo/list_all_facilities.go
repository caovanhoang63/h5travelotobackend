package roomfacilityrepo

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	roomfacilitymodel "h5travelotobackend/module/roomfacilities/model"
)

type ListFacilityTypesStore interface {
	ListAllRoomFacilityType(ctx context.Context) ([]roomfacilitymodel.RoomFacilityType, error)
}

type ListFacilitiesStore interface {
	ListRoomFacilityByType(ctx context.Context, typeId int) ([]roomfacilitymodel.RoomFacility, error)
}

type listFacilityTypesRepo struct {
	facilityStore     ListFacilitiesStore
	facilityTypeStore ListFacilityTypesStore
}

func NewListFacilityTypesRepo(facilityStore ListFacilitiesStore, facilityTypeStore ListFacilityTypesStore) *listFacilityTypesRepo {
	return &listFacilityTypesRepo{facilityStore: facilityStore, facilityTypeStore: facilityTypeStore}
}

func (repo *listFacilityTypesRepo) ListAllRoomFacilityType(ctx context.Context) ([]roomfacilitymodel.RoomFacilityType, error) {
	types, err := repo.facilityTypeStore.ListAllRoomFacilityType(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(roomfacilitymodel.EntityName, err)
	}

	for i := range types {
		types[i].Facilities, err = repo.facilityStore.ListRoomFacilityByType(ctx, types[i].Id)
		if err != nil {
			return nil, common.ErrCannotListEntity(roomfacilitymodel.EntityName, err)
		}
	}
	return types, nil
}
