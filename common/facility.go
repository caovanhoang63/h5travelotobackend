package common

type Facility struct {
	SqlModelNoMask `json:",inline"`
	Name           string `json:"name" gorm:"column:name"`
	NameEn         string `json:"name_en" gorm:"column:name_en"`
	NameVn         string `json:"name_vn" gorm:"column:name_vn" `
}

type FacilityDetail struct {
	SqlModelNoMask `json:",inline"`
	Name           string `json:"name" gorm:"column:name"`
	NameEn         string `json:"name_en" bson:"name_en" gorm:"column:name_en" `
	NameVn         string `json:"name_vn" bson:"name_vn" gorm:"column:name_vn" `
	DescriptionEn  string `json:"description_en" bson:"description_en" gorm:"column:description_en" `
	DescriptionVn  string `json:"description_vn" bson:"description_vn" gorm:"column:description_vn" `
	Type           int    `json:"-" gorm:"column:type"`
}

type Categories []FacilityDetail
