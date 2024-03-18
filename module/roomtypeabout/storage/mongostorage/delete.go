package roomtypeaboutmongostorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
	roomtypeaboutmodel "h5travelotobackend/module/roomtypeabout/model"
)

func (s *mongoStore) Delete(ctx context.Context, room_type_id int) error {
	filter := bson.D{{"room_type_id", room_type_id}, {"status", 1}}

	if result, err := s.db.Collection(roomtypeaboutmodel.RoomTypeAbout{}.CollectionName()).
		UpdateMany(ctx, filter, bson.D{{"status", 0}}); err != nil {
		return common.ErrDb(err)
	} else if result.ModifiedCount == 0 {
		return common.ErrEntityDeleted(roomtypeaboutmodel.EntityName, nil)
	}

	return nil
}
