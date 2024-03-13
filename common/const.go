package common

import "log"

// DbType is a type to represent the type of database
const (
	DbTypeHotel    = 1
	DbTypeUser     = 2
	DbTypeRoomType = 3
	DbTypeRoom     = 4
	DbTypeWorker   = 5
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
