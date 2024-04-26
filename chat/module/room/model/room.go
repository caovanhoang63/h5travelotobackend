package chatmodel

import "h5travelotobackend/common"

const EntityName = "Room"

type Room struct {
	common.MongoModel `bson:",inline"`
	HotelId           int         `json:"-" bson:"hotel_id"`
	HotelFakeId       *common.UID `json:"hotel_id" bson:"hotel_id"`
	UserId            int         `json:"-" bson:"user_id"`
	UserFakeId        *common.UID `json:"user_id" bson:"user_id"`
	Messages          *Messages   `json:"messages,omitempty" bson:"messages,omitempty"`
}

func (r Room) CollectionName() string {
	return "rooms"
}

func (r *Room) Mask(isAdmin bool) {
	r.HotelFakeId = common.NewUIDP(uint32(r.HotelId), common.DbTypeHotel, 0)
	r.UserFakeId = common.NewUIDP(uint32(r.UserId), common.DbTypeUser, 0)
	if r.Messages != nil {
		r.Messages.Mask()
	}
}

type RoomCreate struct {
	HotelId     int         `json:"-" bson:"hotel_id"`
	HotelFakeId *common.UID `json:"hotel_id" bson:"hotel_id"`
	UserId      int         `json:"-" bson:"user_id"`
	UserFakeId  *common.UID `json:"user_id" bson:"user_id"`
}

func (r *RoomCreate) UnMask() {
	r.HotelId = int(r.HotelFakeId.GetLocalID())
	r.UserId = int(r.UserFakeId.GetLocalID())
}
