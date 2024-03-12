package hotelmodel

import (
	"h5travelotobackend/common"
)

type HotelPolicy struct {
	CheckInTime              string  `json:"check_in_time,omitempty" bson:"check_in_time,omitempty"`
	CheckOutTime             string  `json:"check_out_time,omitempty" bson:"check_out_time,omitempty"`
	RequiredDocuments        string  `json:"required_documents,omitempty" bson:"required_documents,omitempty"`
	CheckInInstructions      string  `json:"check_in_instructions,omitempty" bson:"check_in_instructions,omitempty"`
	Deposit                  float64 `json:"deposit,omitempty" bson:"deposit,omitempty"`
	MinAge                   int     `json:"min_age,omitempty" bson:"min_age,omitempty"`
	EarlyCheckInAllowed      bool    `json:"early_check_in_allowed,omitempty" bson:"early_check_in_allowed,omitempty"`
	EarlyCheckInFee          float64 `json:"early_check_in_fee,omitempty" bson:"early_check_in_fee,omitempty"`
	LateCheckOutAllowed      bool    `json:"late_check_out_allowed,omitempty" bson:"late_check_out_allowed,omitempty"`
	LateCheckOutFee          float64 `json:"late_check_out_fee,omitempty" bson:"late_check_out_fee,omitempty"`
	BreakfastIncluded        bool    `json:"breakfast_included,omitempty" bson:"breakfast_included,omitempty"`
	BreakfastPrice           float64 `json:"breakfast_price,omitempty" bson:"breakfast_price,omitempty"`
	SmokingPolicy            string  `json:"smoking_policy,omitempty" bson:"smoking_policy,omitempty"`
	PetPolicy                string  `json:"pet_policy,omitempty" bson:"pet_policy,omitempty"`
	AdditionalPolicies       string  `json:"additional_policies,omitempty" bson:"additional_policies,omitempty"`
	AirportTransferAvailable bool    `json:"airport_transfer_available,omitempty" bson:"airport_transfer_available,omitempty"`
	AirportTransferFee       float64 `json:"airport_transfer_fee,omitempty" bson:"airport_transfer_fee,omitempty"`
}

type HotelAdditionalInfo struct {
	common.MongoModel `bson:",inline" json:",inline"`
	HotelID           int                 `json:"-" bson:"hotel_id"`
	Amenities         map[string][]string `json:"amenities,omitempty" bson:"amenities,omitempty"`
	StayPolicies      *HotelPolicy        `json:"stay_policies,omitempty" bson:"stay_policies,omitempty"`
	AdditionalInfo    string              `json:"additional_info,omitempty" bson:"additional_info,omitempty"`
	Logo              *common.Image       `json:"logo,omitempty" bson:"logo,omitempty"`
	Cover             *common.Images      `json:"cover,omitempty" bson:"cover,omitempty"`
}

func (h *HotelAdditionalInfo) SetHotelId(id int) {
	h.HotelID = id
}

func (HotelAdditionalInfo) CollectionName() string {
	return "hotel_additional_info"
}
