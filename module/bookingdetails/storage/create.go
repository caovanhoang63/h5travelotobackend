package bookingtrackingstorage

import (
	"context"
	"h5travelotobackend/common"
	bookingdetailmodel "h5travelotobackend/module/bookingdetails/model"
)

func (s *sqlStore) Create(ctx context.Context, data *bookingdetailmodel.BookingDetailCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
