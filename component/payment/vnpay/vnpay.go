package vnpay

import (
	"fmt"
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

func (v *VnPay) NewPayInUrl(amount int, bookingId, ip, txnRef string) string {
	orderInfo := fmt.Sprintf("%s%d", bookingId, amount)
	params := newPayInParams(amount*100, ip, orderInfo, txnRef)
	return params.BuildUrl(v)
}

//func (v *VnPay) NewRefundUrl(amount int, txnRef, orderInfo string) string {
//	params := newRefundParams(amount*100, txnRef, orderInfo)
//	return params.BuildUrl(v)
//}

func NewVnPay(hashSecret, tmnCode, localIp string) *VnPay {
	return &VnPay{
		hashSecret: hashSecret,
		tmnCode:    tmnCode,
		localIp:    localIp,
	}
}
