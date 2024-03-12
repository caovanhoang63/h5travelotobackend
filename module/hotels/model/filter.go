package hotelmodel

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

type Filter struct {
	OwnerId    int     `json:"owner_id" form:"owner"`
	ProvinceID int     `json:"province_id" gorm:"column:province_id" form:"province"`
	DistrictID int     `json:"district_id" gorm:"column:district_id" form:"district"`
	WardID     int     `json:"ward_id" gorm:"column:ward_id" form:"ward"`
	Distance   float64 `json:"distance" form:"distance"`
	Lat        float64 `json:"lat" form:"lat"`
	Lng        float64 `json:"lng" form:"lng"`
}

func (f *Filter) Validate() error {
	if (f.Distance+f.Lng+f.Lat) != 0 && (f.Distance == 0 || f.Lat == 0 || f.Lng == 0) {
		return ErrInvalidLocation
	}
	if f.Lat != 0 && f.Lng != 0 {
		if !govalidator.IsLatitude(govalidator.ToString(f.Lat)) {
			return ErrInvalidLocation
		}
		if !govalidator.IsLongitude(govalidator.ToString(f.Lng)) {
			return ErrInvalidLocation
		}
	}
	return nil
}

var ErrInvalidLocation = errors.New("invalid location")
