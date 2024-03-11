package hotelmodel

type Filter struct {
	OwnerId    int     `json:"owner_id" form:"owner"`
	ProvinceID int     `json:"province_id" gorm:"column:province_id" form:"province"`
	DistrictID int     `json:"district_id" gorm:"column:district_id" form:"district"`
	WardID     int     `json:"ward_id" gorm:"column:ward_id" form:"ward"`
	Distance   float64 `json:"distance" form:"distance"`
	Lat        float64 `json:"lat" form:"lat"`
	Lng        float64 `json:"lng" form:"lng"`
}
