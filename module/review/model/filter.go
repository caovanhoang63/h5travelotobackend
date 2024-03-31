package reviewmodel

import (
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
)

type Filter struct {
	UserId         int         `json:"-" bson:"user_id"`
	UserFakeId     *common.UID `json:"user_id" bson:"-"`
	HotelId        int         `json:"-" bson:"hotel_id"`
	HotelFakeId    *common.UID `json:"hotel_id" bson:"-"`
	BookingId      int         `json:"-" bson:"booking_id"`
	BookingFakeId  *common.UID `json:"booking_id" bson:"-"`
	RoomTypeId     int         `json:"-" bson:"room_type_id"`
	RoomTypeFakeId *common.UID `json:"room_type_id" bson:"-"`
	Rating         int         `json:"rating" bson:"rating"`
}

func (f *Filter) UnMask() {
	if f.UserFakeId != nil {
		f.UserId = int(f.HotelFakeId.GetLocalID())
	}
	if f.HotelFakeId != nil {
		f.HotelId = int(f.HotelFakeId.GetLocalID())
	}
	if f.BookingFakeId != nil {
		f.BookingId = int(f.BookingFakeId.GetLocalID())
	}
	if f.RoomTypeFakeId != nil {
		f.RoomTypeId = int(f.RoomTypeFakeId.GetLocalID())
	}
}

func (f *Filter) ToBsonD() (bson.D, error) {
	filter := bson.D{}
	if f.UserId != 0 {
		filter = append(filter, bson.E{Key: "user_id", Value: f.UserId})
	}
	if f.HotelId != 0 {
		filter = append(filter, bson.E{Key: "hotel_id", Value: f.HotelId})
	}
	if f.BookingId != 0 {
		filter = append(filter, bson.E{Key: "booking_id", Value: f.BookingId})
	}
	if f.RoomTypeId != 0 {
		filter = append(filter, bson.E{Key: "room_type_id", Value: f.RoomTypeId})
	}
	if f.Rating != 0 {
		filter = append(filter, bson.E{Key: "rating", Value: f.Rating})
	}
	return filter, nil
}
