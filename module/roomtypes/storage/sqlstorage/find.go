package roomtypesqlstorage

import (
	"context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*roomtypemodel.RoomType, error) {
	var data roomtypemodel.RoomType

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}

	return &data, nil
}

func (s *sqlStore) FindDTODataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*common.RoomTypeDTO, error) {
	var data common.RoomTypeDTO

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
