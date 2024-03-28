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

func (s *sqlStore) FindRoomDTOById(
	ctx context.Context,
	condition map[string]interface{},
) (*common.DTORoom, error) {
	var data common.DTORoom

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}

	return &data, nil
}

func (s *sqlStore) FindRoomsDTOByIds(
	ctx context.Context,
	condition map[string]interface{},
	ids []int,
) ([]common.DTORoom, error) {
	var data []common.DTORoom

	if err := s.db.Where(condition).Where("id IN ?", ids).Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return data, nil
}
