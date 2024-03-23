package hotelstorage

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, update *hotelmodel.HotelUpdate) error {
	if err := s.db.Table(update.TableName()).Where("id = ?", id).Updates(update).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
