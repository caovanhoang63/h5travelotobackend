package htcollection

import (
	"h5travelotobackend/common"
)

const EntityName = "HotelCollection"

type HotelCollection struct {
	common.SqlModel `json:",inline"`
	Name            string        `json:"name" gorm:"column:name;"`
	UserId          int           `json:"user_id" gorm:"column:user_id"`
	Cover           *common.Image `json:"cover" gorm:"column:cover"`
	UserFakeId      *common.UID   `json:"-" form:"-" gorm:"column:user_id;"`
	IsPrivate       bool          `json:"is_private" gorm:"column:is_private"`
}

func (HotelCollection) TableName() string {
	return "hotels_collections"
}

func (h *HotelCollection) Mask(isAdmin bool) {
	h.GenUID(common.DbTypeUser)
	h.UserFakeId = common.NewUIDP(uint32(h.UserId), common.DbTypeUser, 0)
}

type HotelCollectionCreate struct {
	common.SqlModel `json:",inline"`
	UserId          int           `json:"-" form:"-" gorm:"column:user_id"`
	Name            string        `json:"name" form:"name" gorm:"column:name"`
	Cover           *common.Image `json:"cover" gorm:"column:cover"`
	IsPrivate       *bool         `json:"is_private" gorm:"column:is_private"`
}

func (hc *HotelCollectionCreate) Mask(isAdmin bool) {
	hc.GenUID(common.DbTypeUser)
}

func (h *HotelCollectionCreate) TableName() string {
	return HotelCollection{}.TableName()
}

type HotelCollectionUpdate struct {
	Name      string        `json:"name"  gorm:"column:name"`
	Cover     *common.Image `json:"cover" gorm:"column:cover"`
	IsPrivate *bool         `json:"is_private" gorm:"column:is_private"`
}

func (h *HotelCollectionUpdate) TableName() string {
	return HotelCollection{}.TableName()
}
