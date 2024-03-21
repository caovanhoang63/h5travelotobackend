package bookingtrackingstorage

import (
	"context"
	"h5travelotobackend/common"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
)

func (s *sqlStore) Delete(ctx context.Context, bookingId int) error {
	db := s.db.Where("booking_id = ?", bookingId).Table(bookingtrackingmodel.BookingTracking{}.TableName())
	if err := db.Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
