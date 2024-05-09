package hotelmodel

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"log"
	"strings"
	"time"
)

type Filter struct {
	SearchText   string         `json:"search_text" form:"search_text"`
	Adults       int            `json:"adults" form:"adults"`
	Children     int            `json:"children" form:"children"`
	Star         int            `json:"star" form:"star"`
	ByLocation   bool           `json:"by_location" form:"by_location"`
	ByProvince   bool           `json:"by_province" form:"by_province"`
	ByDistrict   bool           `json:"by_district" form:"by_district"`
	ByWard       bool           `json:"by_ward" form:"by_ward"`
	ByHotelName  bool           `json:"by_hotel_name" form:"by_hotel_name"`
	ProvinceCode string         `json:"province_code" form:"province_code" `
	DistrictCode string         `json:"district_code" form:"district_code" `
	WardCode     string         `json:"ward_code" form:"ward_code"`
	ListFacility []string       `json:"list_facility" form:"list_facility"`
	Lat          *types.Float64 `json:"lat" form:"lat" `
	Lng          *types.Float64 `json:"lng" form:"lng"`
	StartDate    *time.Time     `json:"start_date" form:"start_date"`
	EndDate      *time.Time     `json:"end_date" form:"end_date"`
}

func (f *Filter) Validate() error {
	if f.Adults+f.Children == 0 {
		return ErrOccupancyEmpty
	}
	if f.StartDate == nil {
		return ErrStartIsEmpty
	}
	if f.EndDate == nil {
		return ErrEndIsEmpty
	}

	start := *f.StartDate
	end := *f.EndDate

	if start.After(end) {
		return ErrStartDateAfterEndDate
	}

	if start.After(time.Now()) {
		return ErrStartDateIsAfterNow
	}

	trueCount := 0
	if f.ByLocation {
		trueCount++
		if f.Lat == nil || f.Lng == nil {
			return ErrLatLonEmpty
		}
		lat, err := f.Lat.MarshalJSON()
		if err != nil {
			return ErrInvalidLatLon
		}
		lng, err := f.Lat.MarshalJSON()
		if err != nil {
			return ErrInvalidLatLon
		}
		if !govalidator.IsLatitude(string(lat)) || !govalidator.IsLongitude(string(lng)) {
			return ErrInvalidLatLon
		}
	}
	if f.ByProvince {
		trueCount++
		if f.ProvinceCode == "" {
			return ErrProvinceCodeEmpty
		}
	}
	if f.ByDistrict {
		trueCount++
		if f.DistrictCode == "" {
			return ErrDistrictCodeEmpty
		}
	}
	if f.ByWard {
		trueCount++
		if f.WardCode == "" {
			return ErrWardCodeEmpty
		}
	}
	if f.ByHotelName {
		trueCount++
		if f.SearchText == "" {
			return ErrHotelNameIsEmpty
		}
	}

	if trueCount > 1 {
		return ErrTooManyFilters
	}

	f.SearchText = strings.TrimSpace(f.SearchText)
	if trueCount == 0 && f.SearchText == "" {
		return ErrSearchTextIsEmpty
	}
	return nil
}

func (f *Filter) ToSearchRequest() (*search.Request, error) {
	if !(f.ByWard && f.ByLocation && f.ByProvince && f.ByDistrict && f.ByHotelName) {
		log.Println("hotel", f.SearchText)
		return &search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Should: []types.Query{
						{
							Match: map[string]types.MatchQuery{
								"name": {Query: f.SearchText},
							},
						},
					},
				},
			},
		}, nil
	}

	if f.ByLocation {
		if f.Lng != nil && f.Lat != nil {
			return &search.Request{
				Query: &types.Query{
					Bool: &types.BoolQuery{
						Should: []types.Query{
							{
								Match: map[string]types.MatchQuery{
									"name": {Query: f.SearchText},
								},
							},
						},
					},
					GeoDistance: &types.GeoDistanceQuery{
						Distance: "20km",
						GeoDistanceQuery: map[string]types.GeoLocation{
							"location": types.LatLonGeoLocation{Lat: *f.Lat, Lon: *f.Lng},
						},
					},
				},
			}, nil

		}

	}

	if f.ByProvince {
		return &search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"province.province_code": {Query: f.ProvinceCode},
				},
			},
		}, nil
	}
	if f.ByDistrict {
		return &search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"district.district_code": {Query: f.DistrictCode},
				},
			},
		}, nil
	}

	if f.ByWard {
		return &search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"ward.ward_code": {Query: f.WardCode},
				},
			},
		}, nil
	}
	if f.ByHotelName {
		return &search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"name": {Query: f.SearchText},
				},
			},
		}, nil
	}
	return nil, ErrNoFilterIsSet
}

var (
	ErrSearchTextIsEmpty     = errors.New("search text can not be empty")
	ErrNoFilterIsSet         = errors.New("no filter is set")
	ErrTooManyFilters        = errors.New("too many filters are set")
	ErrHotelNameIsEmpty      = errors.New("hotel name can not be empty")
	ErrProvinceCodeEmpty     = errors.New("province code can not be empty")
	ErrDistrictCodeEmpty     = errors.New("district code can not be empty")
	ErrWardCodeEmpty         = errors.New("ward code can not be empty")
	ErrLatLonEmpty           = errors.New("lat and lon can not be empty")
	ErrInvalidLatLon         = errors.New("lat and lon are invalid")
	ErrOccupancyEmpty        = errors.New("occupancy can not be empty")
	ErrStartDateAfterEndDate = errors.New("start date can not be after end date")
	ErrStartDateIsAfterNow   = errors.New("start date can not be after now")
	ErrStartIsEmpty          = errors.New("start date can not be empty")
	ErrEndIsEmpty            = errors.New("end date can not be empty")
)
