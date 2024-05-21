package walletmodel

import "h5travelotobackend/common"

const EntityName = "HotelWallet"

type HotelWallet struct {
	common.SqlModel `json:",inline"`
	HotelID         int     `json:"hotel_id" gorm:"column:hotel_id;index"`
	Balance         float64 `json:"balance" gorm:"column:balance"`
	Currency        string  `json:"currency" gorm:"column:currency"`
}

func (HotelWallet) TableName() string {
	return "hotel_wallets"
}
