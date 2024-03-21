package bookingsqlstorage

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/bookingmodel"
)

func (s *sqlStore) Create(ctx context.Context, data *bookingmodel.BookingCreate) error {

	db := s.db
	err := db.WithContext(ctx).Create(&data).Error
	if err != nil {
		return common.ErrDb(err)
	}
	return nil
}
