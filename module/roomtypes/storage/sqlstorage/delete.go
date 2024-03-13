package roomtypesqlstorage

import (
	"context"
	"h5travelotobackend/common"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.WithContext(ctx).Where("id = ?", id).
		Table(roomtypemodel.RoomType{}.TableName()).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
