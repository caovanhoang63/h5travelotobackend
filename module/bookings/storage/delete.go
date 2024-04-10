package bookingsqlstorage

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/module/bookings/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.
		Table(bookingmodel.Booking{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
