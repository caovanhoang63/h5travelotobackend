package bookingdetailstorage

import (
	"context"
	"h5travelotobackend/common"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
)

func (s *sqlStore) Create(ctx context.Context, data []bookingdetailmodel.BookingDetailCreate, oldIds []int) error {
	db := s.db.Table(bookingdetailmodel.BookingDetail{}.TableName())

	db = db.Begin()

	// Delete old data
	if len(oldIds) > 0 {
		if err := db.Where("booking_id = ? and room_id in ?", data[0].BookingId, oldIds).Delete(nil).Error; err != nil {
			db.Rollback()
			return common.ErrDb(err)
		}
	}

	// Create new data
	if err := db.Create(&data).Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}

	db.Commit()
	return nil
}
