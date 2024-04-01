package common

type DTOHotel struct {
	Id            int      `json:"id" gorm:"column:id"`
	Name          string   `json:"name" gorm:"column:name"`
	OwnerId       int      `json:"owner_id" gorm:"column:owner_id"`
	Status        int      `json:"status" gorm:"column:status"`
	FacilitiesIds []string `json:"facilities_ids" gorm:"-"`
}

func (DTOHotel) TableName() string {
	return "hotels"
}
