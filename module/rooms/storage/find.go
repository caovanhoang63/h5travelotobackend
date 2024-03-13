package roomstorage

import (
	"context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	roommodel "h5travelotobackend/module/rooms/model"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*roommodel.Room, error) {
	var data roommodel.Room

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
