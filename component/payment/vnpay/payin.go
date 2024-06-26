package vnpay

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strconv"
	"time"
)

type payInParams struct {
	Version       string `json:"vnp_Version"`
	VnpAmount     int    `json:"vnp_Amount"`
	VnpCommand    string `json:"vnp_Command"`
	VnpCreateDate string `json:"vnp_CreateDate"`
	VnpCurrCode   string `json:"vnp_CurrCode"`
	VnpIpAddr     string `json:"vnp_IpAddr"`
	VnpLocale     string `json:"vnp_Locale"`
	VnpOrderInfo  string `json:"vnp_OrderInfo"`
	VnpOrderType  string `json:"vnp_OrderType"`
	VnpReturnUrl  string `json:"vnp_ReturnUrl"`
	VnpTmnCode    string `json:"vnp_TmnCode"`
	VnpTxnRef     string `json:"vnp_TxnRef"`
	VnpSecureHash string `json:"vnp_SecureHash"`
	VnpBankCode   string `json:"vnp_BankCode"`
}

func newPayInParams(Amount int, ip, orderInfo, txnRef, currency string, createdDate string) *payInParams {
	return &payInParams{
		VnpAmount:     Amount,
		VnpIpAddr:     ip,
		VnpOrderInfo:  orderInfo,
		VnpTxnRef:     txnRef,
		VnpCurrCode:   currency,
		VnpCreateDate: createdDate,
	}
}

func (p *payInParams) BuildUrl(pay *VnPay) string {
	baseUrl := "https://sandbox.vnpayment.vn/paymentv2/vpcpay.html?"
	param := "vnp_Amount=" + strconv.Itoa(p.VnpAmount) +
		"&vnp_Command=" + commandPay +
		"&vnp_CreateDate=" + time.Now().Format(vnPayTimeLayout) +
		"&vnp_CurrCode=" + p.VnpCurrCode +
		"&vnp_IpAddr=" + "127.0.0.1" +
		"&vnp_Locale=" + localeVn +
		"&vnp_OrderInfo=" + p.VnpOrderInfo +
		"&vnp_OrderType=" + orderType +
		"&vnp_ReturnUrl=" + "https%3A%2F%2Fapi.h5traveloto.site%2Fv1%2Fpayment%2Fvnpay%2Fipn" +
		"&vnp_TmnCode=" + pay.tmnCode +
		"&vnp_TxnRef=" + p.VnpTxnRef +
		"&vnp_Version=" + version
	hasher := hmac.New(sha512.New, []byte(pay.hashSecret))
	hasher.Write([]byte(param))
	return baseUrl + param + "&vnp_SecureHash=" + hex.EncodeToString(hasher.Sum(nil))
}
