package invoicemodel

import (
	"github.com/shopspring/decimal"
	"h5travelotobackend/common"
)

const EntityName = "Invoice"

type Invoice struct {
	common.SqlModel `json:",inline"`
	BookingId       int             `json:"-" gorm:"column:booking_id;"`
	BookingFakeId   *common.UID     `json:"booking_id" gorm:"column:booking_id;"`
	DealId          int             `json:"-" gorm:"column:deal_id;"`
	PayInHotel      bool            `json:"pay_in_hotel" gorm:"column:pay_in_hotel;"`
	DealFakeId      *common.UID     `json:"deal_id" gorm:"column:deal_id;"`
	TotalAmount     decimal.Decimal `json:"total_amount" gorm:"column:total_amount;"`
	DiscountAmount  decimal.Decimal `json:"discount_amount" gorm:"column:discount_amount;"`
	FinalAmount     decimal.Decimal `json:"final_amount" gorm:"column:final_amount;"`
	Currency        string          `json:"currency" gorm:"column:currency;"`
	PaymentStatus   string          `json:"payment_status" gorm:"column:payment_status;"`
	LedgerUpdated   bool            `json:"ledger_updated" gorm:"column:ledger_updated;"`
	WalletUpdated   bool            `json:"wallet_updated" gorm:"column:wallet_updated;"`
}

func (Invoice) TableName() string {
	return "invoices"
}

func (i *Invoice) Mask(isAdmin bool) {
	i.GenUID(common.DbTypeInvoice)
	i.BookingFakeId = common.NewUIDP(uint32(i.BookingId), common.DbTypeBooking, 0)
	if i.DealId != 0 {
		i.DealFakeId = common.NewUIDP(uint32(i.DealId), common.DbTypeDeal, 0)
	}
}

func (i *Invoice) UnMask() {
	i.BookingId = int(i.BookingFakeId.GetLocalID())
	if i.DealFakeId != nil {
		i.DealId = int(i.DealFakeId.GetLocalID())
	}
}

type InvoiceCreate struct {
	common.SqlModel `json:",inline"`
	BookingId       int         `json:"-" gorm:"column:booking_id;"`
	BookingFakeId   *common.UID `json:"booking_id" gorm:"column:booking_id;"`
	DealId          int         `json:"-" gorm:"column:deal_id;"`
	DealFakeId      *common.UID `json:"deal_id" gorm:"column:deal_id;"`
	Currency        string      `json:"currency" gorm:"column:currency;"`
	PayInHotel      bool        `json:"pay_in_hotel" gorm:"column:pay_in_hotel;"`
}

func (InvoiceCreate) TableName() string {
	return Invoice{}.TableName()
}

func (i *InvoiceCreate) Mask(isAdmin bool) {
	i.GenUID(common.DbTypeInvoice)
	i.BookingFakeId = common.NewUIDP(uint32(i.BookingId), common.DbTypeBooking, 0)
	if i.DealId != 0 {
		i.DealFakeId = common.NewUIDP(uint32(i.DealId), common.DbTypeDeal, 0)
	}
}

func (i *InvoiceCreate) UnMask() {
	i.BookingId = int(i.BookingFakeId.GetLocalID())
	if i.DealFakeId != nil {
		i.DealId = int(i.DealFakeId.GetLocalID())
	}
}

type InvoiceUpdate struct {
	BookingId      int             `json:"-" gorm:"column:booking_id;"`
	BookingFakeId  *common.UID     `json:"booking_id" gorm:"column:booking_id;"`
	DealId         int             `json:"-" gorm:"column:deal_id;"`
	DealFakeId     *common.UID     `json:"deal_id" gorm:"column:deal_id;"`
	TotalAmount    decimal.Decimal `json:"total_amount" gorm:"column:total_amount;"`
	DiscountAmount decimal.Decimal `json:"discount_amount" gorm:"column:discount_amount;"`
	FinalAmount    decimal.Decimal `json:"final_amount" gorm:"column:final_amount;"`
	Currency       string          `json:"currency" gorm:"column:currency;"`
	PaymentStatus  string          `json:"payment_status" gorm:"column:payment_status;"`
	LedgerUpdated  bool            `json:"ledger_updated" gorm:"column:ledger_updated;"`
	WalletUpdated  bool            `json:"wallet_updated" gorm:"column:wallet_updated;"`
}

func (InvoiceUpdate) TableName() string {
	return Invoice{}.TableName()
}

func (i *InvoiceUpdate) Mask(isAdmin bool) {
	i.BookingFakeId = common.NewUIDP(uint32(i.BookingId), common.DbTypeBooking, 0)
	if i.DealId != 0 {
		i.DealFakeId = common.NewUIDP(uint32(i.DealId), common.DbTypeDeal, 0)
	}
}

func (i *InvoiceUpdate) UnMask() {
	i.BookingId = int(i.BookingFakeId.GetLocalID())
	if i.DealFakeId != nil {
		i.DealId = int(i.DealFakeId.GetLocalID())
	}
}
