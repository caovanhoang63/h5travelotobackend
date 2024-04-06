package hoteldetailmodel

import (
	"h5travelotobackend/common"
)

const EntityName = "HotelDetail"

// HotelDetail struct
type HotelDetail struct {
	common.SqlModel      `json:",inline"`
	HotelId              int     `json:"hotel_id" gorm:"column:hotel_id"`
	NumberOfFloor        int     `json:"number_of_floor" gorm:"column:number_of_floor"`
	DistanceToCenterCity float64 `json:"distance_to_center_city" gorm:"column:distance_to_center_city"`
	Description          string  `json:"description" gorm:"column:description"`
	LocationDetail       string  `json:"location_detail" gorm:"column:location_detail"`
	CheckInTime          string  `json:"check_in_time" gorm:"column:check_in_time"`
	CheckOutTime         string  `json:"check_out_time" gorm:"column:check_out_time"`
	RequireDocument      string  `json:"require_document" gorm:"column:require_document"`
	MinimumAge           int     `json:"minimum_age" gorm:"column:minimum_age"`
	CancellationPolicy   float32 `json:"cancellation_policy" gorm:"column:cancellation_policy"`
	SmokingPolicy        string  `json:"smoking_policy" gorm:"column:smoking_policy"`
	PetPolicy            string  `json:"pet_policy" gorm:"column:pet_policy"`
	AdditionalPolicy     string  `json:"additional_policy" gorm:"column:additional_policy"`
}

// TableName func
func (HotelDetail) TableName() string {
	return "hotel_details"
}

type HotelDetailCreate HotelDetail

func (HotelDetailCreate) TableName() string {
	return HotelDetail{}.TableName()
}

type HotelDetailUpdate struct {
	HotelId              int     `json:"hotel_id" gorm:"column:hotel_id"`
	NumberOfFloor        int     `json:"number_of_floor" gorm:"column:number_of_floor"`
	DistanceToCenterCity float64 `json:"distance_to_center_city" gorm:"column:distance_to_center_city"`
	Description          string  `json:"description" gorm:"column:description"`
	LocationDetail       string  `json:"location_detail" gorm:"column:location_detail"`
	CheckInTime          string  `json:"check_in_time" gorm:"column:check_in_time"`
	CheckOutTime         string  `json:"check_out_time" gorm:"column:check_out_time"`
	RequireDocument      string  `json:"require_document" gorm:"column:require_document"`
	MinimumAge           int     `json:"minimum_age" gorm:"column:minimum_age"`
	CancellationPolicy   float32 `json:"cancellation_policy" gorm:"column:cancellation_policy"`
	SmokingPolicy        string  `json:"smoking_policy" gorm:"column:smoking_policy"`
	PetPolicy            string  `json:"pet_policy" gorm:"column:pet_policy"`
	AdditionalPolicy     string  `json:"additional_policy" gorm:"column:additional_policy"`
}

func (HotelDetailUpdate) TableName() string {
	return HotelDetail{}.TableName()
}
