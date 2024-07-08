package walletstorage

import (
	"errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	walletmodel "h5travelotobackend/payment/module/wallet/model"
)

func (s *store) FindHotelWallet(ctx context.Context, hotelId int) (*walletmodel.HotelWallet, error) {
	var result walletmodel.HotelWallet
	if err := s.db.Table(walletmodel.HotelWallet{}.TableName()).Where("hotel_id = ?", hotelId).First(&result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDb(err)
	}
	return &result, nil
}
