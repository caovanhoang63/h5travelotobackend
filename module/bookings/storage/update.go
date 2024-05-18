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
