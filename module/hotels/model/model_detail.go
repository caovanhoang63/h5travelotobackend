package hotelmodel

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"h5travelotobackend/common"
)

type HotelPolicy struct {
	CheckInTime              string  `json:"check_in_time" bson:"check_in_time,omitempty"`
	CheckOutTime             string  `json:"check_out_time" bson:"check_out_time,omitempty"`
	RequiredDocuments        string  `json:"required_documents" bson:"required_documents,omitempty"`
	CheckInInstructions      string  `json:"check_in_instructions" bson:"check_in_instructions,omitempty"`
	Deposit                  float64 `json:"deposit" bson:"deposit,omitempty"`
	MinAge                   int     `json:"min_age" bson:"min_age,omitempty"`
	EarlyCheckInAllowed      bool    `json:"early_check_in_allowed" bson:"early_check_in_allowed,omitempty"`
	EarlyCheckInFee          float64 `json:"early_check_in_fee" bson:"early_check_in_fee,omitempty"`
	LateCheckOutAllowed      bool    `json:"late_check_out_allowed" bson:"late_check_out_allowed,omitempty"`
	LateCheckOutFee          float64 `json:"late_check_out_fee" bson:"late_check_out_fee,omitempty"`
	BreakfastIncluded        bool    `json:"breakfast_included" bson:"breakfast_included,omitempty"`
	BreakfastPrice           float64 `json:"breakfast_price" bson:"breakfast_price,omitempty"`
	SmokingPolicy            string  `json:"smoking_policy" bson:"smoking_policy,omitempty"`
	PetPolicy                string  `json:"pet_policy" bson:"pet_policy,omitempty"`
	AdditionalPolicies       string  `json:"additional_policies" bson:"additional_policies,omitempty"`
	AirportTransferAvailable bool    `json:"airport_transfer_available" bson:"airport_transfer_available,omitempty"`
	AirportTransferFee       float64 `json:"airport_transfer_fee" bson:"airport_transfer_fee,omitempty"`
}

type HotelAdditionalInfo struct {
	ID             primitive.ObjectID  `json:"_" bson:"_id,omitempty"`
	HotelID        int                 `json:"_" bson:"hotel_id"`
	Amenities      map[string][]string `json:"amenities" bson:"amenities,omitempty"`
	StayPolicies   HotelPolicy         `json:"stay_policies" bson:"stay_policies,omitempty"`
	AdditionalInfo string              `json:"additional_info" bson:"additional_info,omitempty"`
	Logo           *common.Image       `json:"logo" bson:"logo,omitempty"`
	Cover          *common.Images      `json:"cover" bson:"cover,omitempty"`
}

func (h *HotelAdditionalInfo) SetHotelId(id int) {
	h.HotelID = id
}

func (HotelAdditionalInfo) CollectionName() string {
	return "hotel_additional_info"
}
