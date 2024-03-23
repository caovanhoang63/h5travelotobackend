package hoteldetailmodel

type Facility struct {
	Name       string          `json:"name" bson:"name"`
	Categories map[string]bool `json:"categories,omitempty" bson:"categories,omitempty"`
}

type FacilityList []Facility
