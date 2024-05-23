package localhoteldetail

import (
	"golang.org/x/net/context"
	"h5travelotobackend/component/appContext"
	hoteldetailbiz "h5travelotobackend/module/hoteldetails/biz"
	hoteldetailmodel "h5travelotobackend/module/hoteldetails/model"
	hoteldetailsqlstorage "h5travelotobackend/module/hoteldetails/storage"
)

type HotelDetailLocalHandler struct {
	appCtx appContext.AppContext
}

func NewHotelDetailLocalHandler(appCtx appContext.AppContext) *HotelDetailLocalHandler {
	return &HotelDetailLocalHandler{appCtx: appCtx}
}

func (h *HotelDetailLocalHandler) GetHotelDetailById(ctx context.Context, id int) (*hoteldetailmodel.HotelDetail, error) {
	store := hoteldetailsqlstorage.NewSqlStore(h.appCtx.GetGormDbConnection())
	biz := hoteldetailbiz.NewGetHotelDetailByIdBiz(store)
	data, err := biz.GetHotelDetailById(ctx, id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
