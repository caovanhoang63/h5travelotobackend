package hotelbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type ListHotelStore interface {
	ListHotelWithCondition(
		ctx context.Context,
		filter *hotelmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]hotelmodel.Hotel, error)
}

type listHotelBiz struct {
	store ListHotelStore
}

func NewListHotelBiz(store ListHotelStore) *listHotelBiz {
	return &listHotelBiz{store: store}
}

func (biz *listHotelBiz) ListHotel(ctx context.Context, filter *hotelmodel.Filter, paging *common.Paging) ([]hotelmodel.Hotel, error) {
	data, err := biz.store.ListHotelWithCondition(ctx, filter, paging, "Province", "District", "Ward")
	if err != nil {
		return nil, common.ErrCannotListEntity(hotelmodel.EntityName, err)
	}

	return data, nil
}
