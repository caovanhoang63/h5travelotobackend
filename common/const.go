package common

import (
	"log"
)

// DbType is a type to represent the type of database
const (
	DbTypeHotel             = 1
	DbTypeUser              = 2
	DbTypeRoomType          = 3
	DbTypeRoom              = 4
	DbTypeWorker            = 5
	DbTypeBooking           = 6
	DbTypeBookingRoomType   = 7
	DbTypeBookingTracking   = 8
	DbTypeBookingDetail     = 9
	DbTypeHotelType         = 10
	DbTypeHotelDetail       = 11
	DbTypeHotelFacilityType = 12
	DbTypeHotelFacility     = 13
	DbTypeRoomFacilityType  = 14
	DbTypeRoomFacility      = 15
)

// RecordStatus is a type to represent the status of a record
const (
	StatusDeleted     = 0
	StatusActive      = 1
	StatusInactive    = 2
	StatusUncompleted = 3
	StatusPending     = 4
	StatusRejected    = 5
	StatusProhibited  = 6
)

const CurrentUser = "user"
const CurrentWorker = "worker"

const (
	RoleAdmin    = "admin"
	RoleWorker   = "worker"
	RoleCustomer = "customer"
	RoleOwner    = "owner"
	RoleStaff    = "staff"
	RoleManager  = "manager"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

// AppRecover is an intelligent function to recover from panic
func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}
