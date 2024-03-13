package roomtypesqlstorage

import (
	"context"
	"h5travelotobackend/common"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, update *roomtypemodel.RoomTypeUpdate) error {

	//TODO: Sua loi khi update bed
	if err := s.db.Table(update.TableName()).Where("id = ?", id).Updates(update).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
