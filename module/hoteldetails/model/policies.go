package hoteldetailmodel

import "time"

type Policies struct {
	CheckInTime         *time.Time `json:"check_in_time,omitempty" bson:"check_in_time,omitempty"`
	CheckOutTime        *time.Time `json:"check_out_time,omitempty" bson:"check_out_time,omitempty"`
	RequiredDocuments   string     `json:"required_documents,omitempty" bson:"required_documents,omitempty"`
	Deposit             float64    `json:"deposit,omitempty" bson:"deposit,omitempty"`
	MinAge              int        `json:"min_age,omitempty" bson:"min_age,omitempty"`
	CancellationPolicy  string     `json:"cancellation_policy,omitempty" bson:"cancellation_policy,omitempty"`
	EarlyCheckInAllowed bool       `json:"early_check_in_allowed,omitempty" bson:"early_check_in_allowed,omitempty"`
	EarlyCheckInFee     float64    `json:"early_check_in_fee,omitempty" bson:"early_check_in_fee,omitempty"`
	LateCheckOutAllowed bool       `json:"late_check_out_allowed,omitempty" bson:"late_check_out_allowed,omitempty"`
	LateCheckOutFee     float64    `json:"late_check_out_fee,omitempty" bson:"late_check_out_fee,omitempty"`
	SmokingPolicy       string     `json:"smoking_policy,omitempty" bson:"smoking_policy,omitempty"`
	PetPolicy           string     `json:"pet_policy,omitempty" bson:"pet_policy,omitempty"`
	AdditionalPolicies  string     `json:"additional_policies,omitempty" bson:"additional_policies,omitempty"`
}
