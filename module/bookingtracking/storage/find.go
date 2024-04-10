package bookingtrackingstorage

import (
	"context"
	"h5travelotobackend/common"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
)

func (s *sqlStore) GetBookingTracking(ctx context.Context, bookingId int) (*bookingtrackingmodel.BookingTracking, error) {

	db := s.db
	var data bookingtrackingmodel.BookingTracking
	if err := db.Where("booking_id = ?", bookingId).First(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return &data, nil
}

func (s *sqlStore) GetBookingTrackings(ctx context.Context, bookingId int) ([]bookingtrackingmodel.BookingTracking, error) {
	db := s.db
	var data []bookingtrackingmodel.BookingTracking
	if err := db.Where("booking_id = ?", bookingId).Find(&data).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return data, nil
}
