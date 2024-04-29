package chatmessagemodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"h5travelotobackend/common"
)

type Filter struct {
	RoomId     *primitive.ObjectID `json:"room_id" bson:"room_id"`
	UserId     int                 `json:"-" bson:"from"`
	UserFakeId *common.UID         `json:"from" bson:"-"`
	Status     int                 `json:"status" bson:"status"`
	IsRead     bool                `json:"is_read" bson:"is_read"`
}
