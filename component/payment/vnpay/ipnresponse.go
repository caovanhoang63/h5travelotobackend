package vnpay

type IPNResponse struct {
	Message string `json:"Message"`
	RspCode string `json:"RspCode"`
}

func NewSuccessResponse() *IPNResponse {
	return &IPNResponse{
		Message: "Confirm Success",
		RspCode: "00",
	}

}

func NewOrderNotFound() *IPNResponse {
	return &IPNResponse{
		Message: "Order not found",
		RspCode: "01",
	}
}

func NewOrderAlreadyConfirmed() *IPNResponse {
	return &IPNResponse{
		Message: "Order already confirmed",
		RspCode: "02",
	}
}

func NewInvalidCheckSum() *IPNResponse {
	return &IPNResponse{
		Message: "Invalid checksum",
		RspCode: "97",
	}
}

func NewInvalidAmount() *IPNResponse {
	return &IPNResponse{
		Message: "Invalid amount",
		RspCode: "04",
	}
}

func NewOtherError() *IPNResponse {
	return &IPNResponse{
		Message: "Other error",
		RspCode: "99",
	}
}

func NewFailedTransaction() *IPNResponse {
	return &IPNResponse{
		Message: "Failed transaction",
		RspCode: "02",
	}
}
