package chatroom

import (
	chatmessage "h5travelotobackend/chat/module/room/model/message"
	"h5travelotobackend/common"
)

const EntityName = "Room"

type Room struct {
	common.MongoModel `bson:",inline"`
	HotelId           int                   `json:"-" bson:"hotel_id"`
	HotelFakeId       *common.UID           `json:"hotel_id" bson:"-"`
	UserId            int                   `json:"-" bson:"user_id"`
	UserFakeId        *common.UID           `json:"user_id" bson:"-"`
	Messages          *chatmessage.Messages `json:"messages,omitempty" bson:"messages,omitempty"`
	LastMessage       *chatmessage.Message  `json:"last_message,omitempty" bson:"-"`
}

type RoomTiny struct {
	common.MongoModel `bson:",inline"`
	HotelId           int                  `json:"-" bson:"hotel_id"`
	HotelFakeId       *common.UID          `json:"hotel_id" bson:"-"`
	UserId            int                  `json:"-" bson:"user_id"`
	UserFakeId        *common.UID          `json:"user_id" bson:"-"`
	LastMessage       *chatmessage.Message `json:"last_message,omitempty" bson:"-"`
}

func (r Room) CollectionName() string {
	return "chat_rooms"
}

func (r *Room) Mask(isAdmin bool) {
	r.HotelFakeId = common.NewUIDP(uint32(r.HotelId), common.DbTypeHotel, 0)
	r.UserFakeId = common.NewUIDP(uint32(r.UserId), common.DbTypeUser, 0)
	if r.Messages != nil {
		r.Messages.Mask()
	}
}

type RoomCreate struct {
	common.MongoModel `bson:",inline"`
	HotelId           int `json:"-" bson:"hotel_id"`
	UserId            int `json:"-" bson:"user_id"`
}
