package common

type successRes struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func SimpleSuccessResponse(data interface{}) *successRes {
	return &successRes{data, nil, nil}
}

func NewSuccessResponse(data, paging, filter interface{}) *successRes {
	return &successRes{data, paging, filter}
}
