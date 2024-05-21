package bookingsqlstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	bookingmodel "h5travelotobackend/module/bookings/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *bookingmodel.BookingUpdate) error {
	if err := s.db.
		Table(bookingmodel.Booking{}.TableName()).
		Where("id = ?", id).
		Updates(data).
		Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *sqlStore) UpdateStateToPaid(ctx context.Context, id int) error {
	if err := s.db.WithContext(ctx).
		Table(bookingmodel.Booking{}.TableName()).
		Where("id = ?", id).
		Update("state", common.BookingStatePaid).
		Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
