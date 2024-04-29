package chatroommodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"h5travelotobackend/common"
)

const EntityName = "Room"

type Room struct {
	common.MongoModel `bson:",inline"`
	HotelId           int                 `json:"-" bson:"hotel_id"`
	HotelFakeId       *common.UID         `json:"hotel_id" bson:"-"`
	UserId            int                 `json:"-" bson:"user_id"`
	UserFakeId        *common.UID         `json:"user_id" bson:"-"`
	TotalMessage      int                 `json:"total_message" bson:"total_message"`
	LastMessageId     *primitive.ObjectID `json:"last_message" bson:"last_message"`
	UserUnRead        int                 `json:"user_unread" bson:"user_unread"`
	HotelUnRead       int                 `json:"hotel_unread" bson:"hotel_unread"`
}

func (r Room) CollectionName() string {
	return "chat_rooms"
}

func (r *Room) Mask(isAdmin bool) {
	r.HotelFakeId = common.NewUIDP(uint32(r.HotelId), common.DbTypeHotel, 0)
	r.UserFakeId = common.NewUIDP(uint32(r.UserId), common.DbTypeUser, 0)
}

type RoomCreate struct {
	common.MongoModel `bson:",inline"`
	HotelId           int                 `json:"-" bson:"hotel_id"`
	UserId            int                 `json:"-" bson:"user_id"`
	TotalMessage      int                 `json:"total_message" bson:"total_message"`
	LastMessageId     *primitive.ObjectID `json:"last_message" bson:"last_message"`
	UserUnRead        int                 `json:"user_unread" bson:"user_unread"`
	HotelUnRead       int                 `json:"hotel_unread" bson:"hotel_unread"`
}
