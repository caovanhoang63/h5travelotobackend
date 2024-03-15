package roomtypesqlstorage

import (
	"context"
	"gorm.io/gorm"
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

func (s *sqlStore) IncreaseTotalRoom(ctx context.Context, id int) error {
	if err := s.db.Table(roomtypemodel.RoomType{}.TableName()).
		Where("id = ?", id).
		Update("total_room", gorm.Expr("total_room + ?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *sqlStore) DecreaseTotalRoom(ctx context.Context, id int) error {
	if err := s.db.Table(roomtypemodel.RoomType{}.TableName()).
		Where("id = ?", id).
		Update("total_room", gorm.Expr("total_room - ?", 1)).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
