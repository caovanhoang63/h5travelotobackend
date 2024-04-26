package chatstorage

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	chatmodel "h5travelotobackend/chat/module/room/model"
	"h5travelotobackend/common"
)

func (s *mongoStore) CreateRoom(ctx context.Context, create *chatmodel.RoomCreate) error {
	coll := s.db.Collection(chatmodel.Room{}.CollectionName())
	_, err := coll.InsertOne(ctx, create)
	if err != nil {
		return common.ErrDb(err)
	}
	return nil
}

func (s *mongoStore) CreateMessage(ctx context.Context, roomId string, create *chatmodel.Message) error {
	coll := s.db.Collection(chatmodel.Room{}.CollectionName())
	update := bson.M{"$push": bson.M{"messages": create}}
	filter := bson.M{"_id": roomId}
	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return common.ErrDb(err)
	}
	return nil
}
