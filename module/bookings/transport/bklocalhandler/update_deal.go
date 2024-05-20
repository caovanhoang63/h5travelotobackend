package bklocalhandler

import (
	"golang.org/x/net/context"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
	dealsqlstorage "h5travelotobackend/module/deals/storage"
)

func (b *CountBookedRoomHandler) UpdateDeal(ctx context.Context, id, dealId int) error {
	store := bookingsqlstorage.NewSqlStore(b.appCtx.GetGormDbConnection())
	dealStore := dealsqlstorage.NewSqlStore(b.appCtx.GetGormDbConnection())
	biz := bookingbiz.NewAddDealBiz(store, dealStore)
	if err := biz.AddDeal(ctx, id, dealId); err != nil {
		return err
	}
	return nil
}
