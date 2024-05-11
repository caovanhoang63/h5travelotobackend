package bklocalhandler

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/appContext"
	bookingbiz "h5travelotobackend/module/bookings/biz"
	bookingsqlstorage "h5travelotobackend/module/bookings/storage"
)

type CountBookedRoomHandler struct {
	appCtx appContext.AppContext
}

func NewCountBookedRoomLocalHandler(appCtx appContext.AppContext) *CountBookedRoomHandler {
	return &CountBookedRoomHandler{
		appCtx: appCtx,
	}
}

func (h *CountBookedRoomHandler) CountBookedRoom(ctx context.Context,
	rtId int,
	startDate, endDate *common.CivilDate,
) (*int, error) {

	store := bookingsqlstorage.NewSqlStore(h.appCtx.GetGormDbConnection())
	biz := bookingbiz.NewCountBookedRoomBiz(store)

	return biz.CountBookedRoom(ctx, rtId, startDate, endDate)
}
