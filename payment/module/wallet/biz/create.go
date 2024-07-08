package walletbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	walletmodel "h5travelotobackend/payment/module/wallet/model"
)

type CreateWalletStore interface {
	CreateWallet(ctx context.Context, create *walletmodel.HotelWalletCreate) error
	FindHotelWallet(ctx context.Context, hotelId int) (*walletmodel.HotelWallet, error)
}

type createWalletBiz struct {
	store CreateWalletStore
}

func NewCreateWalletBiz(store CreateWalletStore) *createWalletBiz {
	return &createWalletBiz{store: store}
}

func (biz *createWalletBiz) CreateWallet(ctx context.Context, create *walletmodel.HotelWalletCreate) error {
	if _, err := biz.store.FindHotelWallet(ctx, create.HotelId); err != nil {
		if !errors.Is(err, common.RecordNotFound) {
			return common.ErrInternal(err)
		}
	} else {
		return common.ErrInvalidRequest(errors.New("hotelId already exists"))
	}
	create.Balance = 0
	create.Currency = common.VND
	if err := biz.store.CreateWallet(ctx, create); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
