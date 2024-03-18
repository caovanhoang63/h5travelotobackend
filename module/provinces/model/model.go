package provincemodel

const EntityName = "Province"

type Province struct {
	Code       int    `json:"code" gorm:"column:code"`
	Name       string `json:"name" gorm:"column:name"`
	NameEn     string `json:"name_en" gorm:"column:name_en"`
	FullName   string `json:"full_name" gorm:"column:full_name"`
	FullNameEn string `json:"full_name_en" gorm:"column:full_name_en"`
	CodeName   string `json:"code_name" gorm:"column:code_name"`
}

func (Province) TableName() string {
	return "provinces"
}
