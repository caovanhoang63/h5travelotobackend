package hotelmodel

import (
	"errors"
	"h5travelotobackend/common"
)

const (
	EntityName = "hotel"
)

type Hotel struct {
	common.SqlModel     `json:",inline"`
	Name                string               `json:"name" gorm:"column:name"`
	Address             string               `json:"address" gorm:"column:address"`
	Logo                *common.Image        `json:"logo" gorm:"logo"`
	Cover               *common.Images       `json:"cover" gorm:"column:cover"`
	ProvinceID          int                  `json:"province_id" gorm:"column:province_id"`
	DistrictID          int                  `json:"district_id" gorm:"column:district_id"`
	WardID              int                  `json:"ward_id" gorm:"column:ward_id"`
	Lat                 float64              `json:"lat" gorm:"column:lat"`
	Lng                 float64              `json:"lng" gorm:"column:lat"`
	OwnerID             int                  `json:"-"`
	User                *common.SimpleUser   `json:"user" gorm:"foreignKey:OwnerID;preload:false"`
	Rating              float32              `json:"rating" gorm:"rating"`
	HotelAdditionalInfo *HotelAdditionalInfo `json:"hotel_additional_info,omitempty" gorm:"-"`
}

func (Hotel) TableName() string {
	return "hotels"
}

func (data *Hotel) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeHotel)
}

type HotelCreate struct {
	common.SqlModel     `json:",inline" `
	Name                string               `json:"name" gorm:"column:name"`
	Address             string               `json:"address" gorm:"column:address"`
	ProvinceID          int                  `json:"province_id" gorm:"column:province_id"`
	DistrictID          int                  `json:"district_id" gorm:"column:district_id"`
	WardID              int                  `json:"ward_id" gorm:"column:ward_id"`
	Lat                 float64              `json:"lat" gorm:"column:lat"`
	Lng                 float64              `json:"lng" gorm:"column:lat"`
	OwnerID             int                  `json:"-" gorm:"column:owner_id"`
	HotelAdditionalInfo *HotelAdditionalInfo `json:"hotel_additional_info" gorm:"-"`
	Logo                *common.Image        `json:"logo" gorm:"logo"`
	Cover               *common.Images       `json:"cover" gorm:"column:cover"`
}

func (HotelCreate) TableName() string {
	return "hotels"
}

func (data *HotelCreate) Mask(isAdmin bool) {
	data.GenUID(common.DbTypeHotel)
}

func (data *HotelCreate) Validate() error {
	if common.IsEmpty(data.Name) {
		return ErrNameIsEmpty
	}
	return nil
}

type HotelUpdate struct {
	Name       string         `json:"name" gorm:"column:name"`
	Address    string         `json:"address" gorm:"column:address"`
	ProvinceID int            `json:"province_id" gorm:"column:province_id"`
	DistrictID int            `json:"district_id" gorm:"column:district_id"`
	WardID     int            `json:"ward_id" gorm:"column:ward_id"`
	Lat        float64        `json:"lat" gorm:"column:lat"`
	Lng        float64        `json:"lng" gorm:"column:lat"`
	Logo       *common.Image  `json:"logo" gorm:"logo"`
	Cover      *common.Images `json:"cover" gorm:"column:cover"`
}

func (HotelUpdate) TableName() string {
	return "hotels"
}

var (
	ErrNameIsEmpty = errors.New("name can not be empty")
)
