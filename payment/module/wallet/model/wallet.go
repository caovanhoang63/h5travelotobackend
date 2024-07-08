package walletmodel

import "h5travelotobackend/common"

const EntityName = "HotelWallet"

type HotelWallet struct {
	common.SqlModel `json:",inline"`
	HotelId         int     `json:"-" gorm:"column:hotel_id"`
	Balance         float64 `json:"balance" gorm:"column:balance"`
	Currency        string  `json:"currency" gorm:"column:currency"`
}

func (HotelWallet) TableName() string {
	return "hotel_wallets"
}

func (w *HotelWallet) Mask(isAdmin bool) {
	w.GenUID(common.DbTypeWallet)
}

type HotelWalletWithdrawal struct {
	common.SqlModel `json:",inline"`
	Amount          float64 `json:"amount" gorm:"column:balance"`
}

func (w *HotelWalletWithdrawal) UnMask() {
}

type HotelWalletCreate struct {
	HotelFakeId *common.UID `json:"hotel_id" gorm:"-"`
	HotelId     int         `json:"-" gorm:"column:hotel_id"`
	Balance     float64     `json:"balance" gorm:"column:balance"`
	Currency    string      `json:"currency" gorm:"column:currency"`
}

func (w *HotelWalletCreate) UnMask() {
	w.HotelId = int(w.HotelFakeId.GetLocalID())
}

func (HotelWalletCreate) TableName() string {
	return "hotel_wallets"
}

type HotelWalletUpdate struct {
	Amount float64 `json:"amount" gorm:"-"`
}
