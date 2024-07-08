package walletstorage

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	walletmodel "h5travelotobackend/payment/module/wallet/model"
)

func (s *store) CreateWallet(ctx context.Context, create *walletmodel.HotelWalletCreate) error {
	if err := s.db.Create(&create).Error; err != nil {
		return common.ErrDb(err)
	}
	return nil
}
