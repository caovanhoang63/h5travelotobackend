package vnpay

type refundParams struct {
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

//func newRefundParams(requestId, txnRef, bookingId, transDate, transActionDate, ip, secureHash string, amount int) *refundParams {
//
//}

//func (r *refundParams) BuildUrl(v *VnPay) string {
//	baseUrl := "https://sandbox.vnpayment.vn/merchant_webapi/api/transaction"
//
//}
