package chatroomstorage

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	chatroommodel "h5travelotobackend/chat/module/room/model"
	"h5travelotobackend/common"
	"log"
)

func (s *mongoStore) HandleNewMessage(ctx context.Context, message *chatroommodel.ChatMessage) error {
	log.Println("Update chat room")
	coll := s.db.Collection(chatroommodel.Room{}.CollectionName())

	filter := bson.D{
		{"_id", message.RoomId},
	}

	update := bson.D{
		{"$set", bson.D{{"last_message", message.ID}}},
		{"$inc", bson.D{{"total_message", 1}}},
		{"$inc", bson.D{{"user_unread", bson.M{"$cond": bson.A{bson.M{"$eq": bson.A{"$user_id", message.UserId}}, 1, 0}}}}},
		{"$inc", bson.D{{"hotel_unread", bson.M{"$cond": bson.A{bson.M{"$ne": bson.A{"$user_id", message.UserId}}, 1, 0}}}}},
	}

	one, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return common.ErrDb(err)
	}

	log.Println("Modified count", one.ModifiedCount)

	return nil
}
