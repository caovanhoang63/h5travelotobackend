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
	DbTypeDeal              = 16
	DbTypeInvoice           = 17
	DbTypeHotelCollection   = 18
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

const RefreshTokenAliveTime = 24 * 30 * 60 * 60
const AccessTokenAliveTime = 60 * 60

// AppRecover is an intelligent function to recover from panic
func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovery error:", err)
	}
}

const IsDebug = false

// Currency
const (
	VND = "VND"
)

const DbTimeStampLayout = "2006-01-02T15:04:05.999999"

const (
	PaymentMethodVnPay = "vnpay"
	PaymentMethodMomo  = "momo"
)

const (
	PaymentStatusNotStarted = "not_started"
	PaymentStatusExecuting  = "executing"
	PaymentStatusSuccess    = "success"
	PaymentStatusFailed     = "failed"
	PaymentStatusExpired    = "expired"
)

const (
	PaymentTypePayIn  = "pay_in"
	PaymentTypePayOut = "pay_out"
)
