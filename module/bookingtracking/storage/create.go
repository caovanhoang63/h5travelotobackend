package bookingtrackingstorage

import (
	"context"
	"h5travelotobackend/common"
	bookingtrackingmodel "h5travelotobackend/module/bookingtracking/model"
)

func (s *sqlStore) Create(ctx context.Context, data *bookingtrackingmodel.BookingTrackingCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
