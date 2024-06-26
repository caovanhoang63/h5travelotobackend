package chatroommodel

import (
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
)

type Filter struct {
	HotelId     int         `json:"-" bson:"hotel_id"`
	HotelFakeId *common.UID `json:"hotel_id" bson:"hotel_id"`
	UserId      int         `json:"-" bson:"user_id"`
	UserFakeId  *common.UID `json:"user_id" bson:"user_id"`
}

func (f *Filter) ToBson() (bson.D, error) {
	filter := bson.D{}
	if f.UserId != 0 {
		filter = append(filter, bson.E{Key: "user_id", Value: f.UserId})
	}
	if f.HotelId != 0 {
		filter = append(filter, bson.E{Key: "hotel_id", Value: f.HotelId})
	}
	return filter, nil
}

func (f *Filter) UnMask() {
	if f.HotelFakeId != nil {
		f.HotelId = int(f.HotelFakeId.GetLocalID())
	}
	if f.UserFakeId != nil {
		f.UserId = int(f.UserFakeId.GetLocalID())
	}
}

func (f *Filter) Mask(isAdmin bool) {
	f.HotelFakeId = common.NewUIDP(uint32(f.HotelId), common.DbTypeHotel, 0)
	f.UserFakeId = common.NewUIDP(uint32(f.UserId), common.DbTypeUser, 0)
}
