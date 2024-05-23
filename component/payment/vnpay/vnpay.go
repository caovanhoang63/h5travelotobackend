package vnpay

import (
	"fmt"
	"time"
)

const vnPayTimeLayout = "20060102150405"
const version = "2.1.0"
const commandPay = "pay"
const commandRefund = "refund"
const localeVn = "vn"
const orderType = "order"

type VnPay struct {
	hashSecret string
	tmnCode    string
	localIp    string
}

func (v *VnPay) CheckSum(url string) bool {

	return true
}

func (v *VnPay) NewPayInUrl(amount float64, currency, bookingId, ip, txnRef string, time *time.Time) string {
	orderInfo := fmt.Sprintf("%s", bookingId)
	createdDate := time.Format(vnPayTimeLayout)
	params := newPayInParams(int(amount*100), ip, orderInfo, txnRef, currency, createdDate)
	return params.BuildUrl(v)
}

func (v *VnPay) NewRefundUrl(requestId, txnRef, bookingId, createdBy, transType, ip string, amount float64, transDate, createdDate *time.Time) string {
	orderInfo := fmt.Sprintf("%s", bookingId)
	transDateStr := transDate.Format(vnPayTimeLayout)
	createdDateStr := createdDate.Format(vnPayTimeLayout)
	params := newRefundParams(requestId, txnRef, orderInfo, createdBy, transDateStr, createdDateStr, transType, ip, int(amount*100))
	return params.BuildUrl(v)
}

func NewVnPay(hashSecret, tmnCode, localIp string) *VnPay {
	return &VnPay{
		hashSecret: hashSecret,
		tmnCode:    tmnCode,
		localIp:    localIp,
	}
}
