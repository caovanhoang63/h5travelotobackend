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

type GetReviewInfoStore interface {
	GetTotalAndAvg(ctx context.Context, hotelId int) (int, float32, error)
}

type listHotelBiz struct {
	store   ListHotelStore
	rvStore GetReviewInfoStore
}

func NewListHotelBiz(store ListHotelStore, rvStore GetReviewInfoStore) *listHotelBiz {
	return &listHotelBiz{store: store, rvStore: rvStore}
}

func (biz *listHotelBiz) ListHotel(ctx context.Context, filter *hotelmodel.Filter, paging *common.Paging) ([]hotelmodel.Hotel, error) {
	data, err := biz.store.ListHotelWithCondition(ctx, filter, paging, "Province", "District", "Ward")
	if err != nil {
		return nil, common.ErrCannotListEntity(hotelmodel.EntityName, err)
	}

	for i := range data {
		tt, av, err := biz.rvStore.GetTotalAndAvg(ctx, data[i].Id)
		if err != nil {
			data[i].TotalRating = tt
			data[i].Rating = av
		}
	}

	return data, nil
}
