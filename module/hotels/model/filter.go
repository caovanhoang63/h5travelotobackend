package hotelmodel

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"h5travelotobackend/common"
)

type Filter struct {
	OwnerId         int         `json:"owner_id" form:"owner"`
	ProvinceCode    int         `json:"province_code" gorm:"column:province_code"`
	DistrictCode    int         `json:"district_code" gorm:"column:district_code"`
	WardCode        int         `json:"ward_Code" gorm:"column:ward_code"`
	HotelType       int         `json:"-" gorm:"column:hotel_type"`
	HotelTypeFakeId *common.UID `json:"hotel_type" gorm:"-"`
	Distance        float64     `json:"distance" form:"distance"`
	Lat             float64     `json:"lat" form:"lat"`
	Lng             float64     `json:"lng" form:"lng"`
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
