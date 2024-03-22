package hoteltypestorage

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
)

func (s *sqlStore) FindById(ctx context.Context, id int) (*hoteltypemodel.HotelType, error) {
	var data hoteltypemodel.HotelType
	fmt.Println("id", id)
	if err := s.db.Where("id = ?", id).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		} else {
			return nil, common.ErrDb(err)
		}
	}

	return &data, nil

}
