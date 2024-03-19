package common

import "time"

type DTOBooking struct {
	SqlModel  `json:",inline"`
	StartDate *time.Time `json:"start_date" gorm:"column:start_date"`
	EndDate   *time.Time `json:"end_date" gorm:"column:end_date"`
}

func (DTOBooking) TableName() string {
	return "bookings"
}

func (u *DTOBooking) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}
