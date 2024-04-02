package roomfacilitiesbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	roomfacilitymodel "h5travelotobackend/module/roomfacilities/model"
)

type ListRoomFacilitiesRepo interface {
	ListAllRoomFacilityType(ctx context.Context) ([]roomfacilitymodel.RoomFacilityType, error)
}

type listRoomFacilitiesBiz struct {
	repo ListRoomFacilitiesRepo
}

func NewListRoomFacilities(repo ListRoomFacilitiesRepo) *listRoomFacilitiesBiz {
	return &listRoomFacilitiesBiz{repo: repo}
}

func (biz *listRoomFacilitiesBiz) ListAllRoomFacilities(ctx context.Context) ([]roomfacilitymodel.RoomFacilityType, error) {
	data, err := biz.repo.ListAllRoomFacilityType(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(roomfacilitymodel.EntityName, err)
	}

	return data, nil
}
