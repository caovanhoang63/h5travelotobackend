package reviewmodel

import (
	"go.mongodb.org/mongo-driver/bson"
	"h5travelotobackend/common"
)

type Filter struct {
	UserId         int         `json:"-" bson:"user_id" form:"-" `
	UserFakeId     *common.UID `json:"user_id" bson:"-" form:"user_id"`
	HotelId        int         `json:"-" bson:"hotel_id" form:"-"`
	HotelFakeId    *common.UID `json:"hotel_id" bson:"-" form:"hotel_id"`
	BookingId      int         `json:"-" bson:"booking_id" form:"-"`
	BookingFakeId  *common.UID `json:"booking_id" bson:"-" form:"booking_id"`
	RoomTypeId     int         `json:"-" bson:"room_type_id" form:"-"`
	RoomTypeFakeId *common.UID `json:"room_type_id" bson:"-" form:"room_type_id"`
	Rating         int         `json:"rating" bson:"rating" form:"rating"`
}

func (f *Filter) UnMask() error {
	if f.UserFakeId != nil {
		f.UserId = int(f.UserFakeId.GetLocalID())
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
	return nil
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
