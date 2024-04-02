package roomfacilitysqlstore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	roomfacilitymodel "h5travelotobackend/module/roomfacilities/model"
)

func (s *sqlStore) ListAllRoomFacilityType(ctx context.Context) ([]roomfacilitymodel.RoomFacilityType, error) {
	var data []roomfacilitymodel.RoomFacilityType
	db := s.db.Table(roomfacilitymodel.RoomFacilityType{}.TableName())
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (s *sqlStore) ListRoomFacilityByType(ctx context.Context, typeId int) ([]roomfacilitymodel.RoomFacility, error) {
	var data []roomfacilitymodel.RoomFacility
	db := s.db.Table(roomfacilitymodel.RoomFacility{}.TableName())
	if err := db.Where("type = ?", typeId).Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return data, nil
}
