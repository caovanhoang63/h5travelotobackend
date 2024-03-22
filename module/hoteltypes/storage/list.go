package hoteltypestorage

import (
	"context"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
)

func (s *sqlStore) ListAllHotelTypes(ctx context.Context) ([]hoteltypemodel.HotelType, error) {
	var data []hoteltypemodel.HotelType

	if err := s.db.Find(&data, map[string]interface{}{"status": 1}).Error; err != nil {
		return nil, err
	}

	return data, nil
}
