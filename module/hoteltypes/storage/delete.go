package hoteltypestorage

import (
	"context"
	"h5travelotobackend/common"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.Table(hoteltypemodel.HotelType{}.TableName()).
		Where("id = ?", id).Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
