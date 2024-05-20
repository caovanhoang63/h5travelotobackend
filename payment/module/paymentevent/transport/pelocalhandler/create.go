package pelocalhandler

import (
	"golang.org/x/net/context"
	pebiz "h5travelotobackend/payment/module/paymentevent/biz"
	paymenteventmodel "h5travelotobackend/payment/module/paymentevent/model"
	paymenteventstore "h5travelotobackend/payment/module/paymentevent/store"
)

func (h *peLocalHandler) Create(ctx context.Context, create *paymenteventmodel.PaymentEventCreate) error {
	store := paymenteventstore.NewStore(h.appCtx.GetGormDbConnection())
	biz := pebiz.NewPEBiz(store, h.appCtx.GetUUID())

	err := biz.Create(ctx, create)
	if err != nil {
		return err
	}
	return nil
}
