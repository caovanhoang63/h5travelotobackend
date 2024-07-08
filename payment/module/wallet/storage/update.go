package walletstorage

import (
	"errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"h5travelotobackend/common"
	walletmodel "h5travelotobackend/payment/module/wallet/model"
)

func (s *store) UpdateBalance(ctx context.Context, id int, amount float64) error {
	var wallet *walletmodel.HotelWallet
	db := s.db.Begin().Table(walletmodel.HotelWallet{}.TableName()).Clauses(clause.Locking{Strength: "UPDATE"})
	if err := db.Find(&wallet, map[string]interface{}{"hotel_id": id}).Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}

	if err := db.
		Where("hotel_id = ?", id).Update("balance", gorm.Expr(" balance + ?", amount)).Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}

	db.Commit()
	return nil
}

func (s *store) Withdrawal(ctx context.Context, id int, amount float64) error {
	var wallet *walletmodel.HotelWallet
	db := s.db.Begin().Table(walletmodel.HotelWallet{}.TableName()).Clauses(clause.Locking{Strength: "UPDATE"})
	if err := db.Find(&wallet, map[string]interface{}{"hotel_id": id}).Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}

	if wallet != nil {

		if wallet.Balance < amount {
			db.Rollback()
			return common.ErrInvalidRequest(errors.New("insufficient balance"))
		}
	}

	if err := db.
		Where("hotel_id = ?", id).Update("balance", gorm.Expr(" balance -  ?", amount)).Error; err != nil {
		db.Rollback()
		return common.ErrDb(err)
	}

	db.Commit()
	return nil
}
