package hotelmodel

type Filter struct {
	OwnerId    int     `json:"owner_id" form:"owner_id"`
	ProvinceID int     `json:"province_id" gorm:"column:province_id"`
	DistrictID int     `json:"district_id" gorm:"column:district_id"`
	WardID     int     `json:"ward_id" gorm:"column:ward_id"`
	Distance   float64 `json:"distance"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
}
