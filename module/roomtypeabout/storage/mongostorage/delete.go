package roomtypeaboutmongostorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
	roomtypeaboutmodel "h5travelotobackend/module/roomtypeabout/model"
)

func (s *mongoStore) Delete(ctx context.Context, roomTypeId int) error {
	filter := bson.D{{"room_type_id", roomTypeId}}

	coll := s.db.Collection(roomtypeaboutmodel.RoomTypeAbout{}.CollectionName())

	result, err := coll.DeleteOne(ctx, filter)

	if err != nil {
		return common.ErrDb(err)
	}

	if result.DeletedCount == 0 {
		return common.ErrEntityDeleted(roomtypeaboutmodel.EntityName, nil)
	}

	return nil
}
