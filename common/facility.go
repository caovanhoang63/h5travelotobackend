package common

type Facility struct {
	Name       string           `json:"name" bson:"name"`
	NameEn     string           `json:"name_en" bson:"name_en"`
	NameVn     string           `json:"name_vn" bson:"name_vn"`
	Categories []FacilityDetail `json:"categories" bson:"categories"`
}

type FacilityDetail struct {
	Name          string `json:"name" bson:"name"`
	NameEn        string `json:"name_en" bson:"name_en"`
	NameVn        string `json:"name_vn" bson:"name_vn"`
	DescriptionEn string `json:"description_en" bson:"description_en"`
	DescriptionVn string `json:"description_vn" bson:"description_vn"`
}
