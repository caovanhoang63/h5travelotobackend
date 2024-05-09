package hotelmodel

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"h5travelotobackend/common"
	"log"
	"time"
)

type SearchTerm string

const (
	SearchTermName     SearchTerm = "name"
	SearchTermProvince SearchTerm = "provinces"
	SearchTermDistrict SearchTerm = "districts"
	SearchTermWard     SearchTerm = "wards"
	SearchTermLocation SearchTerm = "location"
)

type Filter struct {
	SearchText   string            `json:"search_text" form:"search_text"`
	SearchTerm   *SearchTerm       `json:"search_term" form:"search_term"`
	Adults       int               `json:"adults" form:"adults"`
	Children     int               `json:"children" form:"children"`
	Star         int               `json:"star" form:"star"`
	ListFacility []string          `json:"list_facility" form:"list_facility"`
	Lat          *types.Float64    `json:"lat" form:"lat" `
	Lng          *types.Float64    `json:"lng" form:"lng"`
	StartDate    *common.CivilDate `json:"start_date" form:"start_date"`
	EndDate      *common.CivilDate `json:"end_date" form:"end_date"`
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

	log.Println("start date: ", f.StartDate.ToString())
	log.Println("end date: ", f.EndDate.ToString())

	if f.StartDate.After(*f.EndDate) {
		return ErrStartDateAfterEndDate
	}

	now := time.Now()
	if !f.StartDate.After(common.CivilDate(now)) {
		return ErrStartInPass
	}

	if *f.SearchTerm == SearchTermLocation {
		if f.Lat == nil || f.Lng == nil {
			return ErrLatLonEmpty
		}
		latBytes, err := f.Lat.MarshalJSON()
		if err != nil {
			return common.ErrInvalidRequest(err)
		}
		lngBytes, err := f.Lng.MarshalJSON()
		if err != nil {
			return common.ErrInvalidRequest(err)
		}
		if govalidator.IsLatitude(string(latBytes)) && govalidator.IsLongitude((string(lngBytes))) {
			return ErrInvalidLatLon
		}
	} else {
		if f.SearchText == "" {
			return ErrSearchTextIsEmpty
		}
	}

	return nil
}

func (f *Filter) ToSearchRequest() (*search.Request, error) {
	//TODO implement this
	return nil, nil
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
)
