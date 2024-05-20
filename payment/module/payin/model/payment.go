package payinmodel

import "h5travelotobackend/common"

type PaymentInfo struct {
	BookingId     int         `json:"-"`
	BookingFakeId *common.UID `json:"booking_id" form:"booking_id" binding:"required"`
	DealId        int         `json:"-"`
	DealFakeId    *common.UID `json:"deal_id" form:"deal_id"`
	Currency      string      `json:"currency" form:"currency" binding:"required"`
}

type PaymentInfoResponse struct {
	PaymentUrl string      `json:"payment_url"`
	Currency   string      `json:"currency"`
	BookingId  *common.UID `json:"booking_id"`
	Method     string      `json:"method"`
	DealId     *common.UID `json:"deal_id"`
	Amount     int         `json:"amount"`
}

func (p *PaymentInfo) UnMask() error {
	if p.DealFakeId != nil {
		p.DealId = int(p.DealFakeId.GetLocalID())
	}
	if p.BookingFakeId != nil {
		p.BookingId = int(p.BookingFakeId.GetLocalID())
	}
	return nil
}
