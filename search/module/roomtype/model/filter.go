package rtsearchmodel

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"h5travelotobackend/common"
	"strconv"
)

type Filter struct {
	CacheKey     string            `json:"-"`
	HotelId      int               `json:"hotel_id"`
	Customer     float32           `json:"customer"`
	MinPrice     *float64          `json:"min_price"`
	MaxPrice     *float64          `json:"max_price"`
	RoomQuantity int               `json:"room_quantity" form:"room_quantity"`
	StartDate    *common.CivilDate `json:"start_date"`
	EndDate      *common.CivilDate `json:"end_date"`
}

func (f *Filter) SetDefault() {
	if f.MinPrice == nil {
		f.MinPrice = new(float64)
		*f.MinPrice = 0
	}

	if f.MaxPrice == nil {
		f.MaxPrice = new(float64)
		*f.MaxPrice = 100000000
	}
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
