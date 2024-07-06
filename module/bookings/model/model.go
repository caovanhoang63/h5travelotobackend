package bookingmodel

import (
	"errors"
	"h5travelotobackend/common"
	"time"
)

const EntityName = "Booking"

type RoomType struct {
	common.SqlModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
}

type Booking struct {
	common.SqlModel `json:",inline"`
	HotelId         int               `json:"-" gorm:"column:hotel_id"`
	Hotel           *common.Hotel     `json:"hotel,omitempty" gorm:"foreignKey:HotelId;preload:false"`
	UserFakeId      *common.UID       `json:"user_id"`
	UserId          int               `json:"-" gorm:"column:user_id"`
	RoomTypeId      int               `json:"-" gorm:"column:room_type_id"`
	RoomTypeFakeId  *common.UID       `json:"room_type_id" gorm:"-"`
	RoomType        *RoomType         `json:"room_type" gorm:"foreignKey:RoomTypeId;preload:false"`
	RoomQuantity    int               `json:"room_quantity" gorm:"column:room_quantity"`
	Rooms           []int             `json:"-" gorm:"-"`
	RoomFakeIds     []*common.UID     `json:"rooms" gorm:"-"`
	Adults          int               `json:"adults" gorm:"column:adults"`
	Children        int               `json:"children" gorm:"column:children"`
	DealId          *int              `json:"deal_id" gorm:"column:deal_id"`
	TotalAmount     float64           `json:"total_amount" gorm:"column:total_amount"`
	DiscountAmount  float64           `json:"discount_amount" gorm:"column:discount_amount"`
	FinalAmount     float64           `json:"final_amount" gorm:"column:final_amount"`
	Currency        string            `json:"currency" gorm:"column:currency"`
	State           string            `json:"state" gorm:"column:state"`
	PayInHotel      bool              `json:"pay_in_hotel" gorm:"column:pay_in_hotel"`
	StartDate       *common.CivilDate `json:"start_date" gorm:"column:start_date"`
	EndDate         *common.CivilDate `json:"end_date" gorm:"column:end_date"`
}

func (b *Booking) CheckExpired() {
	if b.State == common.BookingStatePending {
		if time.Now().Sub(*b.CreatedAt).Hours() > 1 {
			b.State = common.BookingStateExpired
		}
	}
}

func (Booking) TableName() string {
	return "bookings"
}

func (b *Booking) Mask(isAdmin bool) {
	b.GenUID(common.DbTypeBooking)
	b.UserFakeId = common.NewUIDP(uint32(b.UserId), common.DbTypeUser, 0)
	b.RoomTypeFakeId = common.NewUIDP(uint32(b.RoomTypeId), common.DbTypeRoomType, 0)
	if b.Hotel != nil {
		b.Hotel.Mask(isAdmin)
	}
	if b.RoomType != nil {
		b.RoomType.GenUID(common.DbTypeRoomType)
	}
	if len(b.Rooms) > 0 {
		for _, id := range b.Rooms {
			b.RoomFakeIds = append(b.RoomFakeIds, common.NewUIDP(uint32(id), common.DbTypeRoom, 0))
		}
	}
}

func (b *Booking) UnMask() {
	b.UserId = int(b.UserFakeId.GetLocalID())
	b.RoomTypeId = int(b.RoomTypeFakeId.GetLocalID())
}

type BookingCreate struct {
	common.SqlModel `json:",inline"`
	HotelId         int               `json:"-" gorm:"column:hotel_id"`
	HotelFakeId     *common.UID       `json:"hotel_id" gorm:"-"`
	UserId          int               `json:"-" gorm:"column:user_id"`
	RoomTypeId      int               `json:"-" gorm:"column:room_type_id"`
	RoomTypeFakeId  *common.UID       `json:"room_type_id" gorm:"-"`
	Adults          int               `json:"adults" gorm:"column:adults"`
	Children        int               `json:"children" gorm:"column:children"`
	TotalAmount     float64           `json:"total_amount" gorm:"column:total_amount"`
	DiscountAmount  float64           `json:"discount_amount" gorm:"column:discount_amount"`
	FinalAmount     float64           `json:"final_amount" gorm:"column:final_amount"`
	Currency        string            `json:"currency" gorm:"column:currency"`
	RoomQuantity    int               `json:"room_quantity" gorm:"column:room_quantity"`
	StartDate       *common.CivilDate `json:"start_date" gorm:"column:start_date"`
	EndDate         *common.CivilDate `json:"end_date" gorm:"column:end_date"`
}

type UpdateAmountResponse struct {
	TotalAmount    float64 `json:"total_amount"`
	DiscountAmount float64 `json:"discount_amount"`
	FinalAmount    float64 `json:"final_amount"`
}

func (b *BookingCreate) UnMask() {
	b.HotelId = int(b.HotelFakeId.GetLocalID())
	b.RoomTypeId = int(b.RoomTypeFakeId.GetLocalID())
}

func (b *BookingCreate) Mask(isAdmin bool) {
	b.GenUID(common.DbTypeBooking)
	b.HotelFakeId = common.NewUIDP(uint32(b.HotelId), common.DbTypeHotel, 0)
	b.RoomTypeId = int(b.RoomTypeFakeId.GetLocalID())
}

func (BookingCreate) TableName() string {
	return Booking{}.TableName()
}

type BookingUpdate struct {
	DealId         *int     `json:"deal_id"`
	DiscountAmount *float64 `json:"-" gorm:"column:discount_amount"`
	FinalAmount    *float64 `json:"-" gorm:"column:final_amount"`
}

type BookingUpdateDeal struct {
	DealId    int `json:"deal_id"`
	BookingId int `json:"booking_id"`
}

func (BookingUpdate) TableName() string {
	return Booking{}.TableName()
}

var (
	ErrInvalidRoomType = common.NewCustomError(
		nil,
		"room type is invalid",
		"ERR_INVALID_ROOM_TYPE")

	ErrRoomNotAvailable = common.NewCustomError(
		nil,
		"room is not available",
		"ERR_ROOM_NOT_AVAILABLE",
	)
	ErrInvalidDeal = common.NewCustomError(
		nil,
		"deal is invalid",
		"ERR_INVALID_DEAL",
	)
)

var (
	ErrOccupancyEmpty        = errors.New("occupancy can not be empty")
	ErrStartDateAfterEndDate = errors.New("start date can not be after end date")
	ErrStartIsEmpty          = errors.New("start date can not be empty")
	ErrEndIsEmpty            = errors.New("end date can not be empty")
	ErrStartInPass           = errors.New("start date can not be in the past")
	ErrRoomQuantityIsZero    = errors.New("room quantity can not be zero")
	ErrDealNotAvailable      = errors.New("deal is not available")
)
