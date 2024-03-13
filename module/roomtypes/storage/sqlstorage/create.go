package roomtypesqlstorage

import (
	"context"
	"h5travelotobackend/common"
	roomtypemodel "h5travelotobackend/module/roomtypes/model"
)

func (s *sqlStore) Create(ctx context.Context, data *roomtypemodel.RoomTypeCreate) error {
	if err := s.db.Create(data).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
