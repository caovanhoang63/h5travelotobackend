package districtmodel

const EntityName = "District"

type District struct {
	Code         string `json:"code" gorm:"column:code"`
	Name         string `json:"name" gorm:"column:name"`
	NameEn       string `json:"name_en" gorm:"column:name_en"`
	FullName     string `json:"full_name" gorm:"column:full_name"`
	FullNameEn   string `json:"full_name_en" gorm:"column:full_name_en"`
	CodeName     string `json:"code_name" gorm:"column:code_name"`
	ProvinceCode string `json:"province_code" gorm:"column:province_code"`
}

func (District) TableName() string {
	return "districts"
}
