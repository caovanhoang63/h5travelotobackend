package hotelmodel

import (
	"errors"
	"h5travelotobackend/common"
)

const (
	EntityName = "hotel"
)

type Hotel struct {
	common.SqlModel      `json:",inline"`
	Name                 string             `json:"name" gorm:"column:name"`
	Address              string             `json:"address" gorm:"column:address"`
	HotelType            int                `json:"-" gorm:"column:hotel_type"`
	HotelTypeFakeId      common.UID         `json:"hotel_type" gorm:"-"`
	OwnerID              int                `json:"-"`
	User                 *common.SimpleUser `json:"user" gorm:"foreignKey:OwnerID;preload:false"`
	NumberOfFloors       int                `json:"number_of_floors" gorm:"column:number_of_floors"`
	DistanceToCenterCity int                `json:"distance_to_center_city" gorm:"column:distance_to_center_city"`
	Hotline              string             `json:"hotline" gorm:"column:hotline"`
	ProvinceCode         int                `json:"province_code" gorm:"column:province_code"`
	DistrictCode         int                `json:"district_code" gorm:"column:district_code"`
	WardCode             int                `json:"ward_Code" gorm:"column:ward_code"`
	Lat                  float64            `json:"lat" gorm:"column:lat"`
	Lng                  float64            `json:"lng" gorm:"column:lng"`
	Star                 int                `json:"star" gorm:"star"`
	Rating               float32            `json:"rating" gorm:"rating"`
}

func (Hotel) TableName() string {
	return "hotels"
}

func (data *Hotel) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeHotel)
}

func (data *Hotel) UnMask() {
	data.HotelType = int(data.HotelTypeFakeId.GetLocalID())
}

type HotelCreate struct {
	common.SqlModel      `json:",inline"`
	Name                 string             `json:"name,omitempty" gorm:"column:name"`
	Address              string             `json:"address,omitempty" gorm:"column:address"`
	HotelType            int                `json:"-" gorm:"column:hotel_type"`
	HotelTypeFakeId      *common.UID        `json:"hotel_type,omitempty" gorm:"-"`
	OwnerID              int                `json:"-"`
	User                 *common.SimpleUser `json:"user" gorm:"foreignKey:OwnerID;preload:false"`
	NumberOfFloors       int                `json:"number_of_floors,omitempty" gorm:"column:number_of_floors"`
	DistanceToCenterCity int                `json:"distance_to_center_city,omitempty" gorm:"column:distance_to_center_city"`
	Hotline              string             `json:"hotline,omitempty" gorm:"column:hotline"`
	ProvinceCode         int                `json:"province_code,omitempty" gorm:"column:province_code"`
	DistrictCode         int                `json:"district_code,omitempty" gorm:"column:district_code"`
	WardCode             int                `json:"ward_Code,omitempty" gorm:"column:ward_code"`
	Lat                  float64            `json:"lat,omitempty" gorm:"column:lat"`
	Lng                  float64            `json:"lng,omitempty" gorm:"column:lng"`
}

func (HotelCreate) TableName() string {
	return "hotels"
}

func (data *HotelCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeHotel)
}

func (data *HotelCreate) UnMask() {
	data.HotelType = int(data.HotelTypeFakeId.GetLocalID())
}

func (data *HotelCreate) Validate() error {
	// TODO Validate data here
	if common.IsEmpty(data.Name) {
		return ErrNameIsEmpty
	}
	return nil
}

type HotelUpdate struct {
	Name                 string      `json:"name" gorm:"column:name"`
	HotelType            int         `json:"-" gorm:"column:hotel_type"`
	HotelTypeFakeId      *common.UID `json:"hotel_type" gorm:"-"`
	Address              string      `json:"address" gorm:"column:address"`
	NumberOfFloors       int         `json:"number_of_floors" gorm:"column:number_of_floors"`
	DistanceToCenterCity int         `json:"distance_to_center_city" gorm:"column:distance_to_center_city"`
	Hotline              string      `json:"hotline" gorm:"column:hotline"`
	ProvinceCode         int         `json:"province_code" gorm:"column:province_code"`
	DistrictCode         int         `json:"district_code" gorm:"column:district_code"`
	WardCode             int         `json:"ward_id" gorm:"column:ward_code"`
}

func (data HotelUpdate) UnMask() {
	data.HotelType = int(data.HotelTypeFakeId.GetLocalID())
}

func (HotelUpdate) TableName() string {
	return "hotels"
}

func (data *HotelUpdate) Validate() error {
	if common.IsEmpty(data.Name) {
		return ErrNameIsEmpty
	}
	return nil
}

var (
	ErrNameIsEmpty                = errors.New("name can not be empty")
	ErrCannotUpdateAdditionalData = common.NewErrorResponse(
		errors.New("cannot update additional data"),
		"cannot update additional data",
		"cannot update additional data",
		"CANNOT_UPDATE_HOTEL_ADDITIONAL_DATA",
	)
)
