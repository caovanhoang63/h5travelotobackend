package rtsearchmodel

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"h5travelotobackend/common"
	"strconv"
	"time"
)

type Filter struct {
	QueryTime    int64             `json:"query_time" form:"query_time"`
	HotelId      int               `json:"-" form:"-"`
	HotelFakeId  *common.UID       `json:"hotel_id" form:"hotel_id" binding:"required"`
	Customer     float32           `json:"-"`
	MinPrice     *float64          `json:"min_price" form:"min_price"`
	MaxPrice     *float64          `json:"max_price" form:"max_price"`
	RoomQuantity int               `json:"room_quantity" form:"room_quantity"`
	StartDate    *common.CivilDate `json:"start_date" form:"start_date" binding:"required"`
	EndDate      *common.CivilDate `json:"end_date" form:"end_date" binding:"required"`
	Adults       int               `json:"adults" form:"adults"`
	Children     int               `json:"children" form:"children"`
}

func (f *Filter) Mask(isAdmin bool) {
	f.HotelFakeId = common.NewUIDP(uint32(f.HotelId), common.DbTypeHotel, 0)
}

func (f *Filter) UnMask(isAdmin bool) {
	f.HotelId = int(f.HotelFakeId.GetLocalID())
}

func (f *Filter) SetDefault() {
	if f.QueryTime == 0 {
		f.QueryTime = time.Now().Unix()
	}

	if f.MinPrice == nil {
		f.MinPrice = new(float64)
		*f.MinPrice = 0
	}
	if f.MaxPrice == nil {
		f.MaxPrice = new(float64)
		*f.MaxPrice = 100000000
	}
}

func (f *Filter) SetHotelId(hotelId int) {
	f.HotelId = hotelId
}

func (f *Filter) ToSearchRequest() (*search.Request, error) {
	minPriceStr := strconv.FormatFloat(*f.MinPrice, 'f', -1, 64)
	maxPriceStr := strconv.FormatFloat(*f.MaxPrice, 'f', -1, 64)
	var minPrice, maxPrice types.Float64
	if err := minPrice.UnmarshalJSON([]byte(minPriceStr)); err != nil {
		return nil, err
	}
	if err := maxPrice.UnmarshalJSON([]byte(maxPriceStr)); err != nil {
		return nil, err
	}
	return &search.Request{
		Query: &types.Query{
			Bool: &types.BoolQuery{
				Must: []types.Query{
					{
						Term: map[string]types.TermQuery{
							"hotel_id": {Value: strconv.Itoa(f.HotelId)},
						},
					},
					{
						Term: map[string]types.TermQuery{
							"status": {Value: 1},
						},
					},
				},
				Filter: []types.Query{
					{
						Range: map[string]types.RangeQuery{
							"price": types.NumberRangeQuery{
								Gte: &minPrice,
								Lte: &maxPrice,
							},
						},
					},
				},
			},
		},
	}, nil
}
