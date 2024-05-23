package chatroommodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"h5travelotobackend/common"
)

type ChatMessage struct {
	common.MongoModel `bson:",inline"`
	RoomId            *primitive.ObjectID `json:"room_id" bson:"room_id"`
	UserId            int                 `json:"-" bson:"from"`
	Message           string              `json:"message,omitempty" bson:"message,omitempty"`
	IsFromCustomer    bool                `json:"is_from_customer" bson:"-"`
}
