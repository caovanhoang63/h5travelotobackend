package common

type DTOReview struct {
	HotelId int `json:"hotel_id" gorm:"column:hotel_id;"`
	UserId  int `json:"user_id" gorm:"column:user_id;"`
	Rating  int `json:"rating" gorm:"column:rating;"`
}
