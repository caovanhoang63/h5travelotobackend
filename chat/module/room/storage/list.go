package chatroomstorage

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"h5travelotobackend/chat/module/room/model"
	"h5travelotobackend/common"
)

func (s *mongoStore) ListChatRoom(ctx context.Context, filter *chatroommodel.Filter, paging *common.Paging) ([]chatroommodel.Room, error) {
	var result []chatroommodel.Room
	filterB, err := filter.ToBson()
	filterB = append(filterB, bson.E{Key: "status", Value: 1})
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	option := &options.FindOptions{}

	if v := paging.FakeCursor; v == "" {
		option.SetSkip(int64(paging.GetOffSet()))
	} else {
		cursor := primitive.ObjectID{}
		err := cursor.UnmarshalText([]byte(v))
		if err != nil {
			filterB = append(filterB, bson.E{Key: "_id", Value: bson.M{"$lt": cursor}})
		} else {
			option.SetSkip(int64(paging.GetOffSet()))
		}
	}

	option.SetLimit(int64(paging.Limit)).SetSort(bson.D{{Key: "_id", Value: -1}})

	coll := s.db.Collection(chatroommodel.Room{}.CollectionName())
	if count, err := coll.CountDocuments(ctx, filterB); err != nil {
		return nil, common.ErrDb(err)
	} else {
		paging.Total = count
	}

	cur, err := coll.Find(ctx, filterB, option)
	if err != nil {
		return nil, common.ErrDb(err)
	}

	for cur.Next(ctx) {
		var chatRoom chatroommodel.Room
		if err := cur.Decode(&chatRoom); err != nil {
			return nil, common.ErrDb(err)
		}
		result = append(result, chatRoom)
	}

	if err := cur.Err(); err != nil {
		return nil, common.ErrDb(err)
	}

	if err := cur.Close(ctx); err != nil {
		return nil, common.ErrDb(err)
	}

	if len(result) > 0 {
		paging.NextCursor = result[len(result)-1].ID.Hex()
	}

	return result, nil
}
