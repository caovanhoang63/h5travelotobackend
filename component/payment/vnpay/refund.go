package vnpay

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strconv"
)

type RefundParams struct {
	VnpRequestId       string `json:"vnp_RequestId" form:"vnp_RequestId"`
	VnpVersion         string `json:"vnp_Version" form:"vnp_Version"`
	VnpCommand         string `json:"vnp_Command" form:"vnp_Command"`
	VnpTmnCode         string `json:"vnp_TmnCode" form:"vnp_TmnCode"`
	VnpTransactionType string `json:"vnp_TransactionType" form:"vnp_TransactionType"`
	VnpTxnRef          string `json:"vnp_TxnRef" form:"vnp_TxnRef"`
	VnpAmount          int    `json:"vnp_Amount" form:"vnp_Amount"`
	VnpOrderInfo       string `json:"vnp_OrderInfo" form:"vnp_OrderInfo"`
	VnpTransDate       string `json:"vnp_TransDate" form:"vnp_TransDate"`
	VnpTransactionDate string `json:"vnp_TransactionDate" form:"vnp_TransactionDate"`
	VnpCreateBy        string `json:"vnp_CreateBy" form:"vnp_CreateBy"`
	VnpCreateDate      string `json:"vnp_CreateDate" form:"vnp_CreateDate"`
	VnpIdAddr          string `json:"vnp_IpAddr" form:"vnp_IpAddr"`
	VnpSecureHash      string `json:"vnp_SecureHash" form:"vnp_SecureHash"`
}

func newRefundParams(requestId, txnRef, orderInfo, createdBy, transDate, createdDate, transType, ip string, amount int) *RefundParams {
	return &RefundParams{
		VnpRequestId:       requestId,
		VnpTxnRef:          txnRef,
		VnpOrderInfo:       orderInfo,
		VnpTransactionDate: transDate,
		VnpTransactionType: transType,
		VnpCreateBy:        createdBy,
		VnpIdAddr:          ip,
		VnpAmount:          amount,
		VnpCreateDate:      createdDate,
	}
}

func (r *RefundParams) BuildUrl(v *VnPay) string {
	baseUrl := "https://sandbox.vnpayment.vn/merchant_webapi/api/transaction"
	data := r.VnpRequestId + "|" +
		version + "|" +
		commandRefund + "|" +
		v.tmnCode + "|" +
		r.VnpTransactionType + "|" +
		r.VnpTxnRef + "|" +
		strconv.Itoa(r.VnpAmount) + "|" +
		r.VnpTransDate + "|" +
		r.VnpCreateBy + "|" +
		r.VnpCreateDate + "|" +
		"127.0.0.1" + "|" +
		r.VnpOrderInfo

	hasher := hmac.New(sha512.New, []byte(v.hashSecret))
	hasher.Write([]byte(data))

	param := "vnp_RequestId=" + r.VnpRequestId +
		"&vnp_Version=" + version +
		"&vnp_Command=" + commandRefund +
		"&vnp_TmnCode=" + v.tmnCode +
		"&vnp_TransactionType=" + r.VnpTransactionType +
		"&vnp_TxnRef=" + r.VnpTxnRef +
		"&vnp_Amount=" + strconv.Itoa(r.VnpAmount) +
		"&vnp_OrderInfo=" + r.VnpOrderInfo +
		"&vnp_TransactionDate=" + r.VnpTransactionDate +
		"&vnp_CreateBy=" + r.VnpCreateBy +
		"&vnp_IpAddr=" + "127.0.0.1" +
		"&vnp_SecureHash" + hex.EncodeToString(hasher.Sum(nil))
	return baseUrl + param
}
