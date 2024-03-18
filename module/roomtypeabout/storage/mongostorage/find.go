package roomtypeaboutmongostorage

import (
	"context"
	"h5travelotobackend/common"
	roomtypeaboutmodel "h5travelotobackend/module/roomtypeabout/model"
)

func (m *mongoStore) FindWithCondition(ctx context.Context, condition map[string]interface{}) (*roomtypeaboutmodel.RoomTypeAbout, error) {
	var data roomtypeaboutmodel.RoomTypeAbout
	result := m.db.Collection(roomtypeaboutmodel.RoomTypeAbout{}.CollectionName()).
		FindOne(ctx, condition)
	if result.Err() != nil {
		return nil, common.ErrDb(result.Err())
	}
	err := result.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
