package bookingsqlstorage

import (
	"context"
	"fmt"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/model"
)

func (s *sqlStore) Create(ctx context.Context, data *bookingmodel.BookingCreate) error {

	db := s.db
	err := db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *sqlStore) CreateFrontDeskBooking(ctx context.Context, data *bookingmodel.FrontDeskBookingCreate) error {
	db := s.db
	fmt.Println("OK")
	tx := db.Begin()
	if err := tx.WithContext(ctx).Create(&data.Booking).Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}
	data.Customer.BookingId = data.Booking.Id
	if err := tx.WithContext(ctx).Create(&data.Customer).Error; err != nil {
		tx.Rollback()
		return common.ErrDb(err)
	}

	tx.Commit()
	return nil
}
