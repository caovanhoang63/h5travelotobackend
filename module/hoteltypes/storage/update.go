package hoteltypestorage

import (
	"context"
	"h5travelotobackend/common"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, update *hoteltypemodel.HotelTypeUpdate) error {
	if err := s.db.Where("id = ?", id).Updates(update).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
