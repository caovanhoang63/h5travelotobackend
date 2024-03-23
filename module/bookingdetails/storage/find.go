package bookingtrackingstorage

import (
	"context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
)

func (s *sqlStore) FindWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*bookingtrackingmodel.BookingTracking, error) {
	db := s.db.Table(bookingtrackingmodel.BookingTracking{}.TableName())
	db = db.Where(condition)
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}
	data := bookingtrackingmodel.BookingTracking{}
	if err := db.First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}

	return &data, nil
}
