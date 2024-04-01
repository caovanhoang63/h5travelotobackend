package hotelfacilitysqlstore

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelfacilitymodel "h5travelotobackend/module/hotelfacilities/model"
)

func (s *sqlStore) ListAllHotelFacilityType(ctx context.Context) ([]hotelfacilitymodel.HotelFacilityType, error) {
	var data []hotelfacilitymodel.HotelFacilityType
	db := s.db.Table(hotelfacilitymodel.HotelFacilityType{}.TableName())
	if err := db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (s *sqlStore) ListHotelFacilityByType(ctx context.Context, typeId int) ([]hotelfacilitymodel.HotelFacility, error) {
	var data []hotelfacilitymodel.HotelFacility
	db := s.db.Table(hotelfacilitymodel.HotelFacility{}.TableName())
	if err := db.Where("type = ?", typeId).Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return data, nil
}
