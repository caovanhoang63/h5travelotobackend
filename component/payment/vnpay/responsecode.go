package vnpay

const (
	IPNSuccess            = "00"
	IPNSuspect            = "07"
	IPNNoInternetBanking  = "09"
	IPN3TimeWrongCardInfo = "10"
	IPNExpired            = "11"
	IPNCardBlocked        = "12"
	IPNWrongOTP           = "13"
	IPNTxnCancelled       = "24"
	IPNBalanceNotEnough   = "51"
	IPNExceedLimit        = "53"
	IPNBankMaintaining    = "75"
	INPWrongPassword      = "79"
	IPNOtherError         = "99"
)
