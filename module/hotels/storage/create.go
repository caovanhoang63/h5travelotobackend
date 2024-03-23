package hotelstorage

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

func (s *sqlStore) CreateCompleted(ctx context.Context, data *hotelmodel.HotelCreate) error {
	data.Status = common.StatusActive
	if err := s.db.Where("id = ? ", data.Id).Updates(&data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *sqlStore) CreateUncompleted(ctx context.Context, data *hotelmodel.HotelCreate) error {
	data.Status = common.StatusUncompleted
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
