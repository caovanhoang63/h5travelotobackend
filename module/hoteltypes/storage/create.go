package hoteltypestorage

import (
	"context"
	"h5travelotobackend/common"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
)

func (s *sqlStore) Create(ctx context.Context, data *hoteltypemodel.HotelTypeCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
