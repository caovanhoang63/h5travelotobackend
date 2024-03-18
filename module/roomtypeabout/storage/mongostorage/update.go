package roomtypeaboutmongostorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
	roomtypeaboutmodel "h5travelotobackend/module/roomtypeabout/model"
)

func (m *mongoStore) Update(ctx context.Context, roomTypeId int, updateData *roomtypeaboutmodel.RoomTypeAboutUpdate) error {
	filter := bson.D{{"room_type_id", roomTypeId}}

	coll := m.db.Collection(updateData.CollectionName())

	result, err := coll.UpdateOne(ctx, filter, bson.D{{"$set", updateData}})
	if err != nil {
		return common.ErrDb(err)
	}

	if result.ModifiedCount == 0 {
		return common.ErrEntityNotFound(roomtypeaboutmodel.EntityName, nil)
	}

	return nil
}
