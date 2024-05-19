package vnpay

type IPNResponse struct {
	VnpAmount            string `json:"vnp_Amount" form:"vnp_Amount"`
	VnpBankCode          string `json:"vnp_BankCode" form:"vnp_BankCode"`
	VnpBankTranNo        string `json:"vnp_BankTranNo" form:"vnp_BankTranNo"`
	VnpCardType          string `json:"vnp_CardType" form:"vnp_CardType"`
	VnpCommand           string `json:"vnp_Command" form:"vnp_Command"`
	VnpOrderInfo         string `json:"vnp_OrderInfo" form:"vnp_OrderInfo"`
	VnpPayDate           string `json:"vnp_PayDate" form:"vnp_PayDate"`
	VnpResponseCode      string `json:"vnp_ResponseCode" form:"vnp_ResponseCode"`
	VnpTmnCode           string `json:"vnp_TmnCode" form:"vnp_TmnCode"`
	VnpTransactionNo     string `json:"vnp_TransactionNo" form:"vnp_TransactionNo"`
	VnpTransactionStatus string `json:"vnp_TransactionStatus" form:"vnp_TransactionStatus"`
	VnpSecureHash        string `json:"vnp_SecureHash" form:"vnp_SecureHash"`
}
