package vnpay

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
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
	hash       hash.Hash
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

func (v *VnPay) hashString(data string) string {
	v.hash.Write([]byte(data))
	return hex.EncodeToString(v.hash.Sum(nil))
}

func NewVnPay(hashSecret, tmnCode, localIp string) *VnPay {
	hash := hmac.New(sha512.New, []byte(hashSecret))
	return &VnPay{
		hashSecret: hashSecret,
		tmnCode:    tmnCode,
		localIp:    localIp,
		hash:       hash,
	}
}
