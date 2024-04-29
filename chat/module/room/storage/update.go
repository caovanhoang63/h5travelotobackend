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
	coll.FindOneAndUpdate(ctx, filter, bson.D{{"$set", bson.D{{"last_message", message.ID}}}})
	update := bson.D{
		{"$set", bson.D{{"last_message", message.ID}}},
		{"$inc", bson.D{{"total_message", 1}}},
	}
	log.Println("From customer: ", message.IsFromCustomer)
	if message.IsFromCustomer {
		update = append(update, bson.E{Key: "$inc", Value: bson.D{{"hotel_unread", 1}}})
	} else {
		update = append(update, bson.E{Key: "$inc", Value: bson.D{{"user_unread", 1}}})
	}

	_, err := coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return common.ErrDb(err)
	}

	return nil
}
