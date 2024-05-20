package payinmodel

import (
	"h5travelotobackend/common"
	"time"
)

const EntityName = "PaymentBooking"

type PaymentBooking struct {
	BookingId      int         `json:"-" gorm:"column:booking_id"`
	BookingFakeID  *common.UID `json:"booking_id" gorm:"-"`
	TxnId          string      `json:"txn_id" gorm:"column:txn_id"`
	CustomerId     int         `json:"-" gorm:"column:customer_id"`
	CustomerFakeId *common.UID `json:"customer_id" gorm:"-"`
	HotelId        int         `json:"-" gorm:"column:hotel_id"`
	HotelFakeId    *common.UID `json:"hotel_id" gorm:"-"`
	Amount         float64     `json:"amount" gorm:"column:amount"`
	Currency       string      `json:"currency" gorm:"column:currency"`
	Method         string      `json:"method" gorm:"column:method"`
	PaymentStatus  string      `json:"payment_status" gorm:"column:payment_status"`
	LedgerUpdated  bool        `json:"ledger_updated" gorm:"column:ledger_updated"`
	WalletUpdated  bool        `json:"wallet_updated" gorm:"column:wallet_updated"`
	Status         int         `json:"status" gorm:"column:status"`
	CreatedAt      *time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt      *time.Time  `json:"updated_at" gorm:"column:updated_at"`
}

func (p PaymentBooking) TableName() string {
	return "payment_bookings"
}

func (p *PaymentBooking) Mask(isAdmin bool) error {
	if p.BookingFakeID != nil {
		p.BookingId = int(p.BookingFakeID.GetLocalID())
	}
	if p.CustomerFakeId != nil {
		p.CustomerId = int(p.CustomerFakeId.GetLocalID())
	}
	if p.HotelFakeId != nil {
		p.HotelId = int(p.HotelFakeId.GetLocalID())
	}
	return nil
}

type PaymentBookingCreate struct {
	BookingId     int         `json:"-" gorm:"column:booking_id" `
	BookingFakeId *common.UID `json:"booking_id" form:"booking_id" binding:"required" gorm:"-"`
	DealId        int         `json:"-" gorm:"-"`
	DealFakeId    *common.UID `json:"deal_id" form:"deal_id" gorm:"-"`
	CustomerId    int         `json:"-" gorm:"customer_id"`
	HotelId       int         `json:"-" gorm:"column:hotel_id"`
	TxnId         string      `json:"txn_id" gorm:"column:txn_id"`
	Method        string      `json:"method" gorm:"column:method"`
	Amount        float64     `json:"amount" gorm:"column:amount"`
	Currency      string      `json:"currency" form:"currency" binding:"required" gorm:"column:currency"`
}

func (p PaymentBookingCreate) TableName() string {
	return PaymentBooking{}.TableName()
}

type PaymentInfoResponse struct {
	PaymentUrl string      `json:"payment_url"`
	Currency   string      `json:"currency"`
	TxnId      string      `json:"txn_id"`
	BookingId  *common.UID `json:"booking_id"`
	Method     string      `json:"method"`
	DealId     *common.UID `json:"deal_id"`
	Amount     float64     `json:"amount"`
}

func (p *PaymentBookingCreate) UnMask() error {
	if p.DealFakeId != nil {

		p.DealId = int(p.DealFakeId.GetLocalID())
	}
	if p.BookingFakeId != nil {
		p.BookingId = int(p.BookingFakeId.GetLocalID())
	}
	return nil
}

type PaymentBookingUpdateStatus struct {
	PaymentStatus string `json:"payment_status" form:"payment_status" binding:"required"`
}

func (p PaymentBookingUpdateStatus) TableName() string {
	return PaymentBooking{}.TableName()
}

var (
	ErrPaymentSuccessfullOrExecuting = common.NewCustomError(
		nil,
		"Payment is already successfull or executing",
		"ErrPaymentSuccessfullOrExecuting")
)
