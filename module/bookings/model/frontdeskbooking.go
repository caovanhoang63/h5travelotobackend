package bookingmodel

type FrontDeskBookingCreate struct {
	Booking  *BookingCreate     `json:"booking"`
	Customer *FrontDeskCustomer `json:"customer"`
}
