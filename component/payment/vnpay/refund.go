package vnpay

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"strconv"
)

const (
	RefundTypeFull = "02"
	RefundTypePart = "03"
)

type RefundParams struct {
	VnpRequestId       string `json:"vnp_RequestId" form:"vnp_RequestId"`
	VnpVersion         string `json:"vnp_Version" form:"vnp_Version"`
	VnpCommand         string `json:"vnp_Command" form:"vnp_Command"`
	VnpTmnCode         string `json:"vnp_TmnCode" form:"vnp_TmnCode"`
	VnpTransactionType string `json:"vnp_TransactionType" form:"vnp_TransactionType"`
	VnpTxnRef          string `json:"vnp_TxnRef" form:"vnp_TxnRef"`
	VnpAmount          string `json:"vnp_Amount" form:"vnp_Amount"`
	VnpOrderInfo       string `json:"vnp_OrderInfo" form:"vnp_OrderInfo"`
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
		VnpAmount:          strconv.Itoa(amount),
		VnpCreateDate:      createdDate,
	}
}

func (r *RefundParams) BuildUrl(v *VnPay) *RefundParams {
	data := r.VnpRequestId + "|" +
		version + "|" +
		commandRefund + "|" +
		v.tmnCode + "|" +
		r.VnpTransactionType + "|" +
		r.VnpTxnRef + "|" +
		r.VnpAmount + "|" +
		r.VnpTransactionDate + "|" +
		r.VnpCreateBy + "|" +
		r.VnpCreateDate + "|" +
		"127.0.0.1" + "|" +
		r.VnpOrderInfo

	hasher := hmac.New(sha512.New, []byte(v.hashSecret))
	hasher.Write([]byte(data))

	r.VnpVersion = version
	r.VnpCommand = commandRefund
	r.VnpTmnCode = v.tmnCode
	r.VnpSecureHash = hex.EncodeToString(hasher.Sum(nil))
	r.VnpIdAddr = "127.0.0.1"

	return r
}
