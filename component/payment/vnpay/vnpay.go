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

func (v *VnPay) CheckSum(url string) bool {

	return true
}

func (v *VnPay) NewPayInUrl(amount float64, currency, bookingId, ip, txnRef string) string {
	orderInfo := fmt.Sprintf("%s", bookingId)
	params := newPayInParams(int(amount)*100, ip, orderInfo, txnRef, currency)
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
