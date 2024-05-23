package refundmodel

import (
	"h5travelotobackend/common"
	"time"
)

type RefundBooking struct {
	BookingId      int         `json:"-" gorm:"column:booking_id"`
	BookingFakeID  *common.UID `json:"booking_id" gorm:"-"`
	PayInTxnId     string      `json:"pay_in_txn_id" gorm:"column:pay_in_txn_id"`
	TxnId          string      `json:"txn_id" gorm:"column:txn_id"`
	CustomerId     int         `json:"-" gorm:"column:customer_id"`
	CustomerFakeId *common.UID `json:"customer_id" gorm:"-"`
	HotelId        int         `json:"-" gorm:"column:hotel_id"`
	HotelFakeId    *common.UID `json:"hotel_id" gorm:"-"`
	Reason         string      `json:"reason" form:"reason" binding:"required" gorm:"column:reason"`
	Amount         float64     `json:"amount" gorm:"column:amount"`
	Currency       string      `json:"currency" gorm:"column:currency"`
	Method         string      `json:"method" gorm:"column:method"`
	PaymentStatus  string      `json:"payment_status" gorm:"column:payment_status"`
	LedgerUpdated  bool        `json:"-" gorm:"column:ledger_updated"`
	WalletUpdated  bool        `json:"-" gorm:"column:wallet_updated"`
	Status         int         `json:"status" gorm:"column:status"`
	CreatedAt      *time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      *time.Time  `json:"updated_at" gorm:"column:updated_at"`
}

func (p RefundBooking) TableName() string {
	return "refund_bookings"
}

type RefundBookingCreate struct {
	BookingId     int         `json:"-" gorm:"column:booking_id" `
	BookingFakeId *common.UID `json:"booking_id" form:"booking_id" binding:"required" gorm:"-"`
	PayInTxnId    string      `json:"pay_in_txn_id" gorm:"column:pay_in_txn_id"`
	CustomerId    int         `json:"-" gorm:"customer_id"`
	UserFakeId    *common.UID `json:"customer_id" gorm:"-"`
	HotelId       int         `json:"-" gorm:"column:hotel_id"`
	HotelFakeId   *common.UID `json:"hotel_id" gorm:"-"`
	TxnId         string      `json:"txn_id" gorm:"column:txn_id"`
	Method        string      `json:"method" gorm:"column:method"`
	Amount        float64     `json:"amount" gorm:"column:amount"`
	Reason        string      `json:"reason" form:"reason" binding:"required" gorm:"column:reason"`
	Currency      string      `json:"currency" form:"currency" binding:"required" gorm:"column:currency"`
	CreatedAt     *time.Time  `json:"created_at" gorm:"column:created_at"`
	PayInDate     *time.Time  `json:"pay_in_dat" gorm:"-"`
	Type          string      `json:"-" gorm:"-"`
}

func (r *RefundBookingCreate) Mask(isAdmin bool) {
	r.BookingFakeId = common.NewUIDP(uint32(r.BookingId), common.DbTypeBooking, 0)
	r.UserFakeId = common.NewUIDP(uint32(r.CustomerId), common.DbTypeUser, 0)
	r.HotelFakeId = common.NewUIDP(uint32(r.HotelId), common.DbTypeHotel, 0)

}

func (RefundBookingCreate) TableName() string {
	return RefundBooking{}.TableName()
}

type RefundBookingUpdateStatus struct {
	PaymentStatus string `json:"payment_status" form:"payment_status" binding:"required"`
}

func (RefundBookingUpdateStatus) TableName() string {
	return RefundBooking{}.TableName()
}
