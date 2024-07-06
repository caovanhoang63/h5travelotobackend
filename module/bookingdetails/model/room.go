package bookingdetailmodel

import "h5travelotobackend/common"

type Room struct {
	common.SqlModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name;"`
	Floor           int    `json:"floor" gorm:"column:floor;"`
}

func (Room) TableName() string {
	return "rooms"
}

func (r *Room) Mask(isAdmin bool) {
	r.GenUID(common.DbTypeRoom)
}
