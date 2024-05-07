package hotelmodel

import (
	"errors"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"h5travelotobackend/common"
	"strings"
	"time"
)

type Filter struct {
	SearchText   string     `json:"search_text"`
	Adults       int        `json:"adults"`
	Children     int        `json:"children"`
	Star         int        `json:"star"`
	ByLocation   bool       `json:"by_location"`
	ByProvince   bool       `json:"by_province"`
	ByDistrict   bool       `json:"by_district"`
	ByWard       bool       `json:"by_ward"`
	ByHotelName  bool       `json:"by_hotel_name"`
	ProvinceCode string     `json:"province_code"`
	DistrictCode string     `json:"district_code"`
	WardCode     string     `json:"ward_code"`
	Name         string     `json:"name"`
	Lat          float64    `json:"lat"`
	Lng          float64    `json:"lng"`
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"date_date"`
}

func (f *Filter) Validate() error {
	trueCount := 0

	if f.ByLocation {
		trueCount++
		if f.Lat == 0 || f.Lng == 0 {
			return ErrLatLonEmpty
		}
		if !common.IsLatitude(f.Lat) || !common.IsLongitude(f.Lng) {
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
		if f.Name == "" {
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
	if f.ByWard && f.ByLocation && f.ByProvince && f.ByDistrict && f.ByHotelName {
		return &search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"name":                   {Query: f.SearchText},
					"province.province_name": {Query: f.SearchText},
					"district.district_name": {Query: f.SearchText},
					"ward.name":              {Query: f.SearchText},
				},
			},
		}, nil
	}

	if f.ByLocation {
		if f.Lng != 0 && f.Lat != 0 {
			return &search.Request{
				Query: &types.Query{
					GeoDistance: &types.GeoDistanceQuery{
						Distance: "20km",
						GeoDistanceQuery: map[string]types.GeoLocation{
							"location_example": types.LatLonGeoLocation{Lat: types.Float64(f.Lat), Lon: types.Float64(f.Lng)},
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
					"name": {Query: f.Name},
				},
			},
		}, nil
	}
	return nil, ErrNoFilterIsSet
}

var (
	ErrSearchTextIsEmpty = errors.New("search text can not be empty")
	ErrNoFilterIsSet     = errors.New("no filter is set")
	ErrTooManyFilters    = errors.New("too many filters are set")
	ErrHotelNameIsEmpty  = errors.New("hotel name can not be empty")
	ErrProvinceCodeEmpty = errors.New("province code can not be empty")
	ErrDistrictCodeEmpty = errors.New("district code can not be empty")
	ErrWardCodeEmpty     = errors.New("ward code can not be empty")
	ErrLatLonEmpty       = errors.New("lat and lon can not be empty")
	ErrInvalidLatLon     = errors.New("lat and lon are invalid")
)
