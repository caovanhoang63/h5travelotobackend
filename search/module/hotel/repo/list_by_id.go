package hotelsearchrepo

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
)

type GetRatingStore interface {
	GetHotelRating(ctx context.Context, id int) (float64, int, error)
}

type ListByIdsStore interface {
	ListByIds(ctx context.Context, ids []int) ([]hotelmodel.Hotel, error)
}
type GetMinPrice interface {
	GetMinPriceByHotelId(ctx context.Context, hotelId int) (float64, error)
}

type listHotelByIdsRepo struct {
	hStore  ListByIdsStore
	rStore  GetMinPrice
	rtStore GetRatingStore
}

func NewHotelsByIdsRepo(hStore ListByIdsStore, rStore GetMinPrice, rtStore GetRatingStore) *listHotelByIdsRepo {
	return &listHotelByIdsRepo{
		hStore:  hStore,
		rStore:  rStore,
		rtStore: rtStore,
	}
}

func (r *listHotelByIdsRepo) ListHotelByIds(ctx context.Context, ids []int) ([]hotelmodel.Hotel, error) {
	hotels, err := r.hStore.ListByIds(ctx, ids)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	for i := len(hotels) - 1; i >= 0; i-- {
		a, err := r.rStore.GetMinPriceByHotelId(ctx, hotels[i].Id)
		if err != nil {
			hotels = append(hotels[:i], hotels[i+1:]...)
			continue
		}
		hotels[i].DisplayPrice = &a
		hotels[i].Rating, hotels[i].TotalRating, _ = r.rtStore.GetHotelRating(ctx, hotels[i].Id)
	}

	return hotels, nil
}
