package hotelstorage

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.
		Table(hotelmodel.Hotel{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": common.StatusDeleted}).
		Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
