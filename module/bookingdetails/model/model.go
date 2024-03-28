package bookingdetailmodel

import (
	"errors"
	"fmt"
	"h5travelotobackend/common"
	"time"
)

const EntityName = "BookingDetail"

type BookingDetail struct {
	BookingId     int        `json:"-" gorm:"column:booking_id"`
	BookingFakeId common.UID `json:"booking_id" gorm:"-"`
	RoomId        int        `json:"-" gorm:"column:room_id"`
	RoomFakeId    common.UID `json:"room_id" gorm:"-"`
	Status        int        `json:"status" gorm:"column:status;default:1"`
	CreatedAt     *time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func (BookingDetail) TableName() string {
	return "booking_details"
}

func (b *BookingDetail) Mask(isAdmin bool) {
	b.BookingFakeId = common.NewUID(uint32(b.BookingId), common.DbTypeBookingDetail, 1)
	b.RoomFakeId = common.NewUID(uint32(b.RoomId), common.DbTypeBookingDetail, 1)
}

func (b *BookingDetail) UnMask() {
	b.BookingId = int(b.BookingFakeId.GetLocalID())
	b.RoomId = int(b.RoomFakeId.GetLocalID())
}

type BookingDetailCreate struct {
	BookingId int `json:"booking_id" gorm:"column:booking_id"`
	RoomId    int `json:"room_id" gorm:"column:room_id"`
}

func (BookingDetailCreate) TableName() string {
	return BookingDetail{}.TableName()
}

type BookingDetailRequest struct {
	BookingId     int         `json:"-"`
	BookingFakeId common.UID  `json:"booking_id" gorm:"-"`
	RoomIds       []int       `json:"-" `
	RoomFakeIds   common.UIDS `json:"room_ids" gorm:"-"`
}

func (b *BookingDetailRequest) UnMask() {
	b.RoomIds = make([]int, len(b.RoomFakeIds))
	for i := range b.RoomFakeIds {
		fmt.Println("RoomFakeIds: ", b.RoomFakeIds[i].GetLocalID())
		b.RoomIds[i] = int(b.RoomFakeIds[i].GetLocalID())
	}
}

func (b *BookingDetailRequest) ToCreate() []BookingDetailCreate {
	var result []BookingDetailCreate
	for i := range b.RoomIds {
		result = append(result, BookingDetailCreate{
			BookingId: b.BookingId,
			RoomId:    b.RoomIds[i],
		})
	}
	return result

}

type BookingDetailUpdate struct {
	RoomFakeId common.UID `json:"room_id" gorm:"-"`
	RoomId     int        `json:"-" gorm:"column:room_id"`
}

var (
	ErrBookingNotFound      = errors.New("Booking not found")
	ErrRoomNotFound         = errors.New("Room not found")
	ErrRoomQuantityExceeded = errors.New("Room quantity exceeded")
	ErrRoomBooked           = errors.New("Room has been booked")
)
