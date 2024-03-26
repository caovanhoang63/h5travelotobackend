package bookingdetailstorage

import (
	"context"
	"h5travelotobackend/common"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
)

func (s *sqlStore) DeleteOne(ctx context.Context, bookingId int, roomId int) error {
	db := s.db.Table(bookingdetailmodel.BookingDetail{}.TableName())
	db = db.Where("booking_id = ? and room_id = ? ", bookingId, roomId)
	if err := db.Delete(nil).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *sqlStore) DeleteAll(ctx context.Context, bookingId int) error {
	db := s.db.Table(bookingdetailmodel.BookingDetail{}.TableName())
	db = db.Where("booking_id = ?", bookingId)
	if err := db.Delete(nil).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
