package roomstorage

import (
	"context"
	"h5travelotobackend/common"
	roommodel "h5travelotobackend/module/rooms/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *roommodel.RoomUpdate) error {
	if err := s.db.Table(data.TableName()).Where("id = ?", id).Updates(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
