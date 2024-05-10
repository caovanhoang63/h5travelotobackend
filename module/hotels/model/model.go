package hotelmodel

import (
	"errors"
	"h5travelotobackend/common"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
)

const (
	EntityName = "hotel"
)

type Hotel struct {
	common.SqlModel `json:",inline"`
	OwnerID         int                `json:"-" gorm:"column:owner_id"`
	User            *common.SimpleUser `json:"user" gorm:"foreignKey:OwnerID;preload:false"`
	Name            string             `json:"name" gorm:"column:name"`
	Address         string             `json:"address" gorm:"column:address"`
	HotelType       int                `json:"hotel_type" gorm:"column:hotel_type"`
	Hotline         string             `json:"hotline" gorm:"column:hotline"`
	Logo            *common.Image      `json:"logo" gorm:"column:logo"`
	Images          *common.Images     `json:"images" gorm:"column:images"`
	ProvinceCode    string             `json:"-" gorm:"column:province_code"`
	Province        *Province          `json:"province,inline" gorm:"foreignKey:ProvinceCode;references:Code"`
	DistrictCode    string             `json:"-" gorm:"column:district_code"`
	District        *District          `json:"district,inline" gorm:"foreignKey:DistrictCode;references:Code"`
	WardCode        string             `json:"-" gorm:"column:ward_code"`
	Ward            *Ward              `json:"ward,inline" gorm:"foreignKey:WardCode;references:Code"`
	Lat             float64            `json:"lat" gorm:"column:lat"`
	Lng             float64            `json:"lng" gorm:"column:lng"`
	Star            int                `json:"star" gorm:"star"`
	Rating          float32            `json:"rating" gorm:"rating"`
	TotalRating     int                `json:"total_rating" gorm:"total_rating"`
	TotalRoomType   int                `json:"total_room_type" gorm:"total_room_type"`
	AvgPrice        float64            `json:"avg_price" gorm:"avg_price"`
}

func (Hotel) TableName() string {
	return "hotels"
}

func (data *Hotel) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeHotel)
}

func (data *Hotel) UnMask() {
}

type HotelCreate struct {
	common.SqlModel `json:",inline"`
	OwnerID         int                                 `json:"-" gorm:"column:owner_id"`
	Name            string                              `json:"name" gorm:"column:name"`
	Address         string                              `json:"address" gorm:"column:address"`
	HotelType       int                                 `json:"hotel_type" gorm:"column:hotel_type"`
	Hotline         string                              `json:"hotline" gorm:"column:hotline"`
	ProvinceCode    string                              `json:"province_code" gorm:"column:province_code"`
	DistrictCode    string                              `json:"district_code" gorm:"column:district_code"`
	WardCode        string                              `json:"ward_Code" gorm:"column:ward_code"`
	Logo            *common.Image                       `json:"logo" gorm:"column:logo"`
	Images          *common.Images                      `json:"images" gorm:"column:images"`
	Lat             float64                             `json:"lat" gorm:"column:lat"`
	Lng             float64                             `json:"lng" gorm:"column:lng"`
	Star            int                                 `json:"star" gorm:"star"`
	FacilityIds     []int                               `json:"facility_ids" gorm:"-"`
	HotelDetail     *hoteldetailmodel.HotelDetailCreate `json:"hotel_detail" gorm:"foreignKey:HotelId;references:Id"`
}

func (HotelCreate) TableName() string {
	return "hotels"
}

func (data *HotelCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeHotel)
}

func (data *HotelCreate) UnMask() {
}

func (data *HotelCreate) Validate() error {
	// TODO Validate data here
	if data.HotelType == 0 {
		return ErrInvalidHotelType
	}

	if common.IsEmpty(data.Name) {
		return ErrNameIsEmpty
	}
	return nil
}

type HotelUpdate struct {
	Name            string         `json:"name,omitempty" gorm:"column:name"`
	Address         string         `json:"address" gorm:"column:address"`
	HotelType       int            `json:"-" gorm:"column:hotel_type"`
	HotelTypeFakeId *common.UID    `json:"hotel_type" gorm:"-"`
	Hotline         string         `json:"hotline" gorm:"column:hotline"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo"`
	Images          *common.Images `json:"images" gorm:"column:images"`
	ProvinceCode    string         `json:"province_code" gorm:"column:province_code"`
	DistrictCode    string         `json:"district_code" gorm:"column:district_code"`
	WardCode        string         `json:"ward_Code" gorm:"column:ward_code"`
	Lat             float64        `json:"lat" gorm:"column:lat"`
	Lng             float64        `json:"lng" gorm:"column:lng"`
	Star            int            `json:"star" gorm:"star"`
	Rating          float32        `json:"rating" gorm:"rating"`
	TotalRating     int            `json:"total_rating" gorm:"total_rating"`
	TotalRoomType   int            `json:"total_room_type" gorm:"total_room_type"`
	AvgPrice        float64        `json:"avg_price" gorm:"avg_price"`
}

func (data HotelUpdate) UnMask() {
	data.HotelType = int(data.HotelTypeFakeId.GetLocalID())
}

func (HotelUpdate) TableName() string {
	return "hotels"
}

func (data *HotelUpdate) Validate() error {
	return nil
}

var (
	ErrNameIsEmpty      = errors.New("name can not be empty")
	ErrInvalidHotelType = errors.New("invalid hotel type")
)

type Province struct {
	Code string `json:"code" gorm:"column:code"`
	Name string `json:"name" gorm:"column:name"`
}

func (Province) TableName() string {
	return "provinces"
}

type District struct {
	Code string `json:"code" gorm:"column:code"`
	Name string `json:"name" gorm:"column:name"`
}

func (District) TableName() string {
	return "districts"
}

type Ward struct {
	Code string `json:"code" gorm:"column:code"`
	Name string `json:"name" gorm:"column:name"`
}

func (Ward) TableName() string {
	return "wards"
}
