package hotelmodel

import (
	"errors"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"strings"
	"time"
)

type Filter struct {
	SearchText string     `json:"search_text"`
	Adults     int        `json:"adults"`
	Children   int        `json:"children"`
	Star       int        `json:"star"`
	Lat        float64    `json:"lat"`
	Lng        float64    `json:"lng"`
	StartDate  *time.Time `json:"start_date"`
	EndDate    *time.Time `json:"date_date"`
}

func (f *Filter) Validate() error {
	f.SearchText = strings.TrimSpace(f.SearchText)
	if f.SearchText == "" && f.Lat == 0 && f.Lng == 0 {
		return ErrSearchTextIsEmpty
	}
	return nil
}

func (f *Filter) ToString() *search.Request {

	req := search.Request{

		Query: &types.Query{
			Match: map[string]types.MatchQuery{
				"name":     {Query: f.SearchText},
				"province": {Query: f.SearchText},
				"district": {Query: f.SearchText},
				"ward":     {Query: f.SearchText},
			},
		},
		Aggregations: map[string]types.Aggregations{},
	}
	if f.Lng != 0 && f.Lat != 0 {
		req.Query.GeoDistance = &types.GeoDistanceQuery{
			Boost:            nil,
			Distance:         "30km",
			DistanceType:     nil,
			GeoDistanceQuery: map[string]types.GeoLocation{},
		}

	}

	return &req
}

var (
	ErrSearchTextIsEmpty = errors.New("search text can not be empty")
)
