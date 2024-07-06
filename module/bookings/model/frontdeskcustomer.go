package bookingmodel

import "h5travelotobackend/common"

type FrontDeskCustomer struct {
	common.SqlModel `json:",inline"`
	BookingId       int    `json:"booking_id" gorm:"booking_id"`
	Name            string `json:"name" gorm:"name"`
	Address         string `json:"address" gorm:"address"`
	Gender          string `json:"gender" gorm:"gender"`
}
