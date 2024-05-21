package common

type PaymentBooking struct {
	TxnId     string  `json:"txn_id" `
	BookingId int     `json:"booking_id"`
	Amount    float64 `json:"amount"`
}
