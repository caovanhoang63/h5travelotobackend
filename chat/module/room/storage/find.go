package chatstorage

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	chatroom "h5travelotobackend/chat/module/room/model/room"
	"h5travelotobackend/common"
)

func (s *mongoStore) FindChatRoom(ctx context.Context, userId int, hotelId int) (*chatroom.Room, error) {
	var result chatroom.Room
	filter := bson.D{{Key: "user_id", Value: userId}, {Key: "hotel_id", Value: hotelId}}

	coll := s.db.Collection(chatroom.Room{}.CollectionName())

	one := coll.FindOne(ctx, filter)
	err := one.Err()
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, common.DocumentNotFound
		}
		return nil, common.ErrDb(err)
	}

	if one != nil {
		if err := one.Decode(&result); err != nil {
			return nil, common.ErrDb(err)
		}
	}

	return &result, nil
}
