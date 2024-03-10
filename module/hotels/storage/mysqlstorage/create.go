package hotelmysqlstorage

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *sqlStore) Create(ctx context.Context, data *hotelmodel.HotelCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
