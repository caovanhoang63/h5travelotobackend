package roomtypeaboutmongostorage

import (
	"context"
	"h5travelotobackend/common"
	"h5travelotobackend/module/roomtypeabout/model"
)

func (m *mongoStore) Create(ctx context.Context, data *roomtypeaboutmodel.RoomTypeAbout) error {
	if _, err := m.db.Collection(data.CollectionName()).InsertOne(ctx, data); err != nil {
		return common.ErrDb(err)
	}
	return nil
}
