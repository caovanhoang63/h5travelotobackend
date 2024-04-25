package hotelstorage

import (
	"context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*hotelmodel.Hotel, error) {
	var data hotelmodel.Hotel
	db := s.db.Table(hotelmodel.Hotel{}.TableName())

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	if err := db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
