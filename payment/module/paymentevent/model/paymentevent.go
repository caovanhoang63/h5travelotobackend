package pemodel

import (
	"h5travelotobackend/common"
	"time"
)

const EntityName = "PaymentEvent"

type PaymentEvent struct {
	TxnId         string     `json:"txn_id" gorm:"column:txn_id"`
	HotelId       int        `json:"hotel_id" gorm:"column:hotel_id"`
	CustomerId    int        `json:"customer_id" gorm:"column:customer_id"`
	IsPaymentDone bool       `json:"is_payment_done" gorm:"column:is_payment_done"`
	PaymentType   string     `json:"payment_type" gorm:"column:payment_type"`
	Status        int        `json:"status" gorm:"column:status"`
	CreatedAt     *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     *time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (p PaymentEvent) TableName() string {
	return "payment_events"
}

type PaymentEventCreate struct {
	TxnId       string `json:"txn_id" gorm:"column:txn_id"`
	HotelId     int    `json:"hotel_id" gorm:"column:hotel_id"`
	CustomerId  int    `json:"customer_id" gorm:"column:customer_id"`
	PaymentType string `json:"payment_type" gorm:"column:payment_type"`
}

func (p PaymentEventCreate) TableName() string {
	return PaymentEvent{}.TableName()
}

func NewErrCannotCreatePaymentEvent(err error) error {
	return common.NewCustomError(
		err,
		" payment event cannot be created",
		"ErrCannotCreatePaymentEvent")
}
