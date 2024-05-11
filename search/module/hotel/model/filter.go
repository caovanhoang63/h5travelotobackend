package hotelmodel

import (
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"h5travelotobackend/common"
	"time"
)

type SearchTerm string

const (
	SearchTermName     SearchTerm = "name"
	SearchTermProvince SearchTerm = "province"
	SearchTermDistrict SearchTerm = "district"
	SearchTermWard     SearchTerm = "ward"
	SearchTermLocation SearchTerm = "location"
)

type Filter struct {
	SearchText   *string           `json:"search_text" form:"search_text"`
	Id           *string           `json:"id" form:"id"`
	SearchTerm   *SearchTerm       `json:"search_term" form:"search_term"`
	Adults       int               `json:"adults" form:"adults"`
	Children     int               `json:"children" form:"children"`
	RoomQuantity int               `json:"room_quantity" form:"room_quantity"`
	Star         int               `json:"star" form:"star"`
	ListFacility []string          `json:"list_facility" form:"list_facility"`
	Lat          *string           `json:"lat" form:"lat" `
	Lng          *string           `json:"lng" form:"lng"`
	MaxPrice     *float64          `json:"max_price" form:"max_price"`
	MinPrice     *float64          `json:"min_price" form:"min_price"`
	StartDate    *common.CivilDate `json:"start_date" form:"start_date"`
	EndDate      *common.CivilDate `json:"end_date" form:"end_date"`
	Customer     float32           `json:"customer"`
}

func (f *Filter) SetDefault() {
	f.Customer = float32(f.Adults) + float32(f.Children)/2.0
	if f.MinPrice == nil {
		f.MinPrice = new(float64)
		*f.MinPrice = 0
	}

	if f.MaxPrice == nil {
		f.MaxPrice = new(float64)
		*f.MaxPrice = 100000000
	}
}

func (f *Filter) Validate() error {
	if f.RoomQuantity == 0 {
		return ErrRoomQuantityIsZero
	}
	if f.Adults+f.Children == 0 {
		return ErrOccupancyEmpty
	}
	if f.StartDate == nil {
		return ErrStartIsEmpty
	}
	if f.EndDate == nil {
		return ErrEndIsEmpty
	}

	if f.StartDate.After(*f.EndDate) {
		return ErrStartDateAfterEndDate
	}

	now := time.Now()
	if !f.StartDate.After(common.CivilDate(now)) {
		return ErrStartInPass
	}

	searchTerm := *f.SearchTerm

	if *f.SearchTerm == SearchTermLocation {
		if f.Lat == nil || f.Lng == nil {
			return ErrLatLonEmpty
		}
		if !govalidator.IsLatitude(*f.Lat) && !govalidator.IsLongitude(*f.Lng) {
			return ErrInvalidLatLon
		}
	}

	if searchTerm == SearchTermProvince || searchTerm == SearchTermDistrict || searchTerm == SearchTermWard {
		if f.Id == nil {
			return ErrIdIsEmpty
		}
	}

	if searchTerm == SearchTermName {
		if f.SearchText == nil {
			return ErrSearchTextIsEmpty
		}
	}

	return nil
}

func (f *Filter) ToSearchRequest() (*search.Request, error) {
	if f.SearchTerm == nil {
		return nil, ErrNoFilter
	}
	searchTerm := string(*f.SearchTerm)
	var field string
	if *f.SearchTerm == SearchTermLocation {
		LatLonGeo := fmt.Sprintf("%s,%s", *f.Lat, *f.Lng)
		return &search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Filter: []types.Query{
						{
							GeoDistance: &types.GeoDistanceQuery{
								Distance: "30km",
								GeoDistanceQuery: map[string]types.GeoLocation{
									"location_geo_point": LatLonGeo,
								},
							},
						},
					},
				},
			},
		}, nil
	} else if *f.SearchTerm == SearchTermName {
		field = string(SearchTermName)
		return &search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Should: []types.Query{
						{
							Match: map[string]types.MatchQuery{
								field: {Query: *f.SearchText},
							},
						},
					},
				},
			},
		}, nil
	} else {
		field = fmt.Sprintf("%s.%s_code", searchTerm, searchTerm)
		return &search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Must: []types.Query{
						{
							Term: map[string]types.TermQuery{
								field: {Value: *f.Id},
							},
						},
					},
				},
			},
		}, nil
	}
}

var (
	ErrSearchTextIsEmpty     = errors.New("search text can not be empty")
	ErrLatLonEmpty           = errors.New("lat and lon can not be empty")
	ErrInvalidLatLon         = errors.New("lat and lon are invalid")
	ErrOccupancyEmpty        = errors.New("occupancy can not be empty")
	ErrStartDateAfterEndDate = errors.New("start date can not be after end date")
	ErrStartDateIsAfterNow   = errors.New("start date can not be after now")
	ErrStartIsEmpty          = errors.New("start date can not be empty")
	ErrEndIsEmpty            = errors.New("end date can not be empty")
	ErrStartInPass           = errors.New("start date can not be in the past")
	ErrIdIsEmpty             = errors.New("id can not be empty")
	ErrNoFilter              = errors.New("no filter")
	ErrRoomQuantityIsZero    = errors.New("room quantity can not be zero")
)
