package reviewmodel

import (
	"h5travelotobackend/common"
)

const EntityName = "Review"

type Review struct {
	common.MongoModel `json:",inline" bson:",inline"`
	UserId            int         `json:"-" bson:"user_id"`
	UserFakeId        *common.UID `json:"user_id" bson:"-"`
	HotelId           int         `json:"-" bson:"hotel_id"`
	HotelFakeId       *common.UID `json:"hotel_id" bson:"-"`
	BookingId         int         `json:"-" bson:"booking_id"`
	BookingFakeId     *common.UID `json:"booking_id" bson:"-"`
	RoomTypeId        int         `json:"-" bson:"room_type_id"`
	RoomTypeFakeId    *common.UID `json:"room_type_id" bson:"-"`
	Rating            int         `json:"rating" bson:"rating"`
	Comment           string      `json:"comment" bson:"comment"`
}

func (Review) CollectionName() string {
	return "reviews"
}

func (r *Review) Mask(isAdmin bool) {
	*r.HotelFakeId = common.NewUID(uint32(r.HotelId), common.DbTypeHotel, 1)
	*r.BookingFakeId = common.NewUID(uint32(r.BookingId), common.DbTypeBooking, 1)
	*r.RoomTypeFakeId = common.NewUID(uint32(r.RoomTypeId), common.DbTypeRoomType, 1)
}

func (r *Review) UnMask() {
	r.HotelId = int(r.HotelFakeId.GetLocalID())
	r.BookingId = int(r.BookingFakeId.GetLocalID())
	r.RoomTypeId = int(r.RoomTypeFakeId.GetLocalID())
}
