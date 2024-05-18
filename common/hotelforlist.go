package common

type Hotel struct {
	SqlModel     `json:",inline"`
	Name         string    `json:"name" gorm:"column:name"`
	Address      string    `json:"address" gorm:"column:address"`
	HotelType    int       `json:"-" gorm:"column:hotel_type"`
	Logo         *Image    `json:"logo" gorm:"column:logo"`
	Images       *Images   `json:"images" gorm:"column:images"`
	ProvinceCode int       `json:"-" gorm:"column:province_code"`
	Province     *Province `json:"province,inline" gorm:"foreignKey:ProvinceCode;references:Code"`
	DistrictCode int       `json:"-" gorm:"column:district_code"`
	District     *District `json:"district,inline" gorm:"foreignKey:DistrictCode;references:Code"`
	WardCode     int       `json:"-" gorm:"column:ward_code"`
	Ward         *Ward     `json:"ward,inline" gorm:"foreignKey:WardCode;references:Code"`
	Star         int       `json:"star" gorm:"star"`
	TotalRating  int       `json:"total_rating" gorm:"total_rating"`
	AvgRating    float64   `json:"avg_rating" gorm:"avg_rating"`
}

func (Hotel) TableName() string {
	return "hotels"
}

func (data *Hotel) Mask(isAdmin bool) {
	data.GenUID(DbTypeHotel)
}

type Province struct {
	Code     int    `json:"code" gorm:"column:code"`
	Name     string `json:"name" gorm:"column:name"`
	FullName string `json:"full_name" gorm:"column:full_name"`
}

func (Province) TableName() string {
	return "provinces"
}

type District struct {
	Code     int    `json:"code" gorm:"column:code"`
	Name     string `json:"name" gorm:"column:name"`
	FullName string `json:"full_name" gorm:"column:full_name"`
}

func (District) TableName() string {
	return "districts"
}

type Ward struct {
	Code     int    `json:"code" gorm:"column:code"`
	Name     string `json:"name" gorm:"column:name"`
	FullName string `json:"full_name" gorm:"column:full_name"`
}

func (Ward) TableName() string {
	return "wards"
}
