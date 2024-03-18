package roomtypeaboutmongostorage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"h5travelotobackend/common"
	roomtypeaboutmodel "h5travelotobackend/module/roomtypeabout/model"
)

func (m *mongoStore) FindWithCondition(ctx context.Context, condition map[string]interface{}) (*roomtypeaboutmodel.RoomTypeAbout, error) {
	var data roomtypeaboutmodel.RoomTypeAbout
	result := m.db.Collection(roomtypeaboutmodel.RoomTypeAbout{}.CollectionName()).
		FindOne(ctx, condition)

	err := result.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, common.ErrEntityNotFound(roomtypeaboutmodel.EntityName, nil)
		}
		return nil, common.ErrDb(err)
	}

	err = result.Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
