package vnpay

import (
	"h5travelotobackend/common"
	"log"
)

type IPNRequest struct {
	VnpAmount            int    `json:"vnp_Amount" form:"vnp_Amount"`
	VnpBankCode          string `json:"vnp_BankCode" form:"vnp_BankCode"`
	VnpBankTranNo        string `json:"vnp_BankTranNo" form:"vnp_BankTranNo"`
	VnpCardType          string `json:"vnp_CardType" form:"vnp_CardType"`
	VnpOrderInfo         string `json:"vnp_OrderInfo" form:"vnp_OrderInfo"`
	VnpPayDate           string `json:"vnp_PayDate" form:"vnp_PayDate"`
	VnpResponseCode      string `json:"vnp_ResponseCode" form:"vnp_ResponseCode"`
	VnpTmnCode           string `json:"vnp_TmnCode" form:"vnp_TmnCode"`
	VnpTransactionNo     string `json:"vnp_TransactionNo" form:"vnp_TransactionNo"`
	VnpTransactionStatus string `json:"vnp_TransactionStatus" form:"vnp_TransactionStatus"`
	VnpTxnRef            string `json:"vnp_TxnRef" form:"vnp_TxnRef"`
	VnpSecureHash        string `json:"vnp_SecureHash" form:"vnp_SecureHash"`
	VnpSecureHashType    string `json:"vnp_SecureHashType" form:"vnp_SecureHashType"`
}

func (i *IPNRequest) GetBookingId() (int, error) {
	log.Println("VnpOrderInfo: ", i.VnpOrderInfo)
	uid, err := common.FromBase58(i.VnpOrderInfo)
	if err != nil {
		return 0, common.ErrInvalidRequest(err)
	}
	return int(uid.GetLocalID()), nil
}

func (i *IPNRequest) GetAmount() float64 {
	return float64(i.VnpAmount) / 100
}

func (i *IPNRequest) getParamString() string {
	return ""
}
