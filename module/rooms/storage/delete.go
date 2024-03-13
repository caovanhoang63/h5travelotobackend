package roomstorage

import (
	"context"
	"h5travelotobackend/common"
	roommodel "h5travelotobackend/module/rooms/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.WithContext(ctx).Where("id = ?", id).
		Table(roommodel.Room{}.TableName()).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
