package subcriber

import (
	"encoding/json"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	"h5travelotobackend/component/pubsub"
	walletmodel "h5travelotobackend/payment/module/wallet/model"
	walletstorage "h5travelotobackend/payment/module/wallet/storage"
	"log"
)

func CreateHotelWallet(appCtx appContext.AppContext, ctx context.Context) consumerJob {
	return consumerJob{
		Title: "Create wallet when hotel created",
		Handler: func(ctx context.Context, message *pubsub.Message) error {
			var hotel common.DTOHotel
			err := json.Unmarshal(message.Data, &hotel)
			if err != nil {
				log.Println(err)
			}
			create := walletmodel.HotelWalletCreate{
				HotelId:  hotel.Id,
				Currency: common.VND,
				Balance:  0,
			}
			store := walletstorage.NewStore(appCtx.GetGormDbConnection())
			return store.CreateWallet(ctx, &create)
		},
	}
}
