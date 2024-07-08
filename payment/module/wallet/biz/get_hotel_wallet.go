package walletbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	walletmodel "h5travelotobackend/payment/module/wallet/model"
)

type GetHotelWalletStore interface {
	FindHotelWallet(ctx context.Context, hotelId int) (*walletmodel.HotelWallet, error)
}

type getHotelWalletBiz struct {
	store GetHotelWalletStore
}

func NewGetHotelWalletBiz(store GetHotelWalletStore) *getHotelWalletBiz {
	return &getHotelWalletBiz{store: store}
}

func (biz *getHotelWalletBiz) GetHotelWalletById(ctx context.Context, hotelId int,
) (*walletmodel.HotelWallet, error) {
	wallet, err := biz.store.FindHotelWallet(ctx, hotelId)
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return nil, common.ErrInternal(err)
		}
		return nil, common.ErrEntityNotFound(walletmodel.EntityName, nil)
	}
	return wallet, nil
}
