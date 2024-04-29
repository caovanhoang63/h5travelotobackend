package chatmessagemodel

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"h5travelotobackend/common"
)

type Filter struct {
	RoomId     *primitive.ObjectID `json:"room_id" bson:"room_id"`
	UserId     int                 `json:"-" bson:"from"`
	UserFakeId *common.UID         `json:"from" bson:"-"`
	Status     int                 `json:"status" bson:"status"`
}

func (f *Filter) ToBsonD() (bson.D, error) {
	filter := bson.D{}
	if f.UserId != 0 {
		filter = append(filter, bson.E{Key: "user_id", Value: f.UserId})
	}
	if f.RoomId != nil {
		filter = append(filter, bson.E{Key: "room_id", Value: f.RoomId})
	}
	return filter, nil
}
