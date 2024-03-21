package bookingtrackingstorage

import (
	"context"
	"h5travelotobackend/common"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
)

func (s *sqlStore) Update(ctx context.Context, bookingId int, data *bookingtrackingmodel.BookingTrackingUpdate) error {
	if err := s.db.Table(bookingtrackingmodel.BookingTracking{}.TableName()).
		Where("booking_id = ?", bookingId).Updates(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
