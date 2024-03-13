package roomstorage

import (
	"context"
	roommodel "h5travelotobackend/module/rooms/model"
)

func (s sqlStore) Create(ctx context.Context, create *roommodel.RoomCreate) error {
	db := s.db

	if err := db.Create(create).Error; err != nil {
		return err
	}

	return nil
}
