package reviewmodel

import (
	"h5travelotobackend/common"
)

const EntityName = "Review"

type Review struct {
	common.MongoModel `json:",inline" bson:",inline"`
	UserId            int                `json:"-" bson:"user_id"`
	UserFakeId        *common.UID        `json:"user_id" bson:"-"`
	User              *common.SimpleUser `json:"user" gorm:"-" bson:"-"`
	HotelId           int                `json:"-" bson:"hotel_id"`
	HotelFakeId       *common.UID        `json:"hotel_id" bson:"-"`
	BookingId         int                `json:"-" bson:"booking_id,omitempty"`
	BookingFakeId     *common.UID        `json:"booking_id,omitempty" bson:"-"`
	RoomTypeId        int                `json:"-" bson:"room_type_id,omitempty"`
	RoomTypeFakeId    *common.UID        `json:"room_type_id,omitempty" bson:"-"`
	Rating            int                `json:"rating" bson:"rating"`
	Comment           string             `json:"comment" bson:"comment"`
	Image             common.Images      `json:"images,omitempty" bson:"images,omitempty"`
}

func (Review) CollectionName() string {
	return "reviews"
}

func (r *Review) Mask(isAdmin bool) {
	r.UserFakeId = common.NewUIDP(uint32(r.UserId), common.DbTypeUser, 1)
	r.HotelFakeId = common.NewUIDP(uint32(r.HotelId), common.DbTypeHotel, 1)
	r.BookingFakeId = common.NewUIDP(uint32(r.BookingId), common.DbTypeBooking, 1)
	r.RoomTypeFakeId = common.NewUIDP(uint32(r.RoomTypeId), common.DbTypeRoomType, 1)
	if r.User != nil {
		r.User.Mask(isAdmin)
	}
}

func (f *Review) UnMask() {
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
