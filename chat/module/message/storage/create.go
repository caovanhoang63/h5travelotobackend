package chatmessagestorage

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/net/context"
	chatmodel "h5travelotobackend/chat/module/message/model"
	chatroommodel "h5travelotobackend/chat/module/room/model"
	"h5travelotobackend/common"
	"log"
)

func (s *mongoStore) CreateMessage(ctx context.Context, roomId string, create *chatmodel.Message) error {
	coll := s.db.Collection(chatroommodel.Room{}.CollectionName())

	update := bson.M{"$push": bson.M{"messages": &create}}
	id, _ := primitive.ObjectIDFromHex(roomId)
	filter := bson.M{"_id": id}
	one, err := coll.UpdateOne(ctx, filter, update)
	log.Println("UpdateOne : ", one.ModifiedCount)
	if err != nil {
		return common.ErrDb(err)
	}
	return nil
}
