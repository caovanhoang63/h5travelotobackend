package hotelsearchbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
)

type ListRecentlyViewedStore interface {
	GetHotelCurrentlyViewed(ctx context.Context, userId int) ([]int, error)
}

type ListHotelByIdsRepo interface {
	ListHotelByIds(ctx context.Context, ids []int) ([]hotelmodel.Hotel, error)
}

func NewListRecentlyViewedBiz(store ListRecentlyViewedStore, hotelStore ListHotelByIdsRepo) *listRecentlyViewedBiz {
	return &listRecentlyViewedBiz{
		rStore: store,
		hRepo:  hotelStore,
	}
}

type listRecentlyViewedBiz struct {
	rStore ListRecentlyViewedStore
	hRepo  ListHotelByIdsRepo
}

func (biz *listRecentlyViewedBiz) ListRecentlyViewed(ctx context.Context, userId int) ([]hotelmodel.Hotel, error) {
	ids, err := biz.rStore.GetHotelCurrentlyViewed(ctx, userId)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	hotels, err := biz.hRepo.ListHotelByIds(ctx, ids)

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return hotels, err
}
