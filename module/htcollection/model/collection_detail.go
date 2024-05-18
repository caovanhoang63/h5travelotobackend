package htcollection

import (
	"h5travelotobackend/common"
	"time"
)

const DetailEntityName = "HotelCollectionDetail"

type HotelCollectionDetail struct {
	CollectionFakeId *common.UID   `json:"collection_id,omitempty" gorm:"-"`
	CollectionId     int           `json:"-" form:"collection_id" gorm:"column:collection_id"`
	HotelId          int           `json:"-" gorm:"column:hotel_id"`
	Hotel            *common.Hotel `json:"hotel,omitempty" gorm:"foreignKey:HotelId;preload:false"`
	CreatedAt        *time.Time    `json:"created_at" gorm:"column:created_at"`
}

func (HotelCollectionDetail) TableName() string {
	return "hotels_collection_details"
}

func (h *HotelCollectionDetail) Mask() {

}

type HotelCollectionDetailCreate struct {
	CollectionId int `json:"-"  gorm:"column:collection_id"`
	HotelId      int `json:"-" gorm:"column:hotel_id"`
}

func (h *HotelCollectionDetailCreate) TableName() string {
	return HotelCollectionDetail{}.TableName()
}
