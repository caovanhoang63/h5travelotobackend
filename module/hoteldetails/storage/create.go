package hoteldetailsqlstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
)

func (s *sqlStore) Create(ctx context.Context, data *hoteldetailmodel.HotelDetailCreate) error {
	db := s.db

	if err := db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}

	return nil
}
