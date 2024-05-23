package hotelsearchbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
)

type ListHotelRepo interface {
	ListHotelWithFilter(ctx context.Context,
		filter *hotelmodel.Filter,
		paging *common.Paging) ([]hotelmodel.Hotel, error)
}

type listHotelBiz struct {
	repo    ListHotelRepo
	rtStore GetRatingStore
}

type GetRatingStore interface {
	GetHotelRating(ctx context.Context, id int) (float64, int, error)
}

func NewListHotelBiz(repo ListHotelRepo, rtStore GetRatingStore) *listHotelBiz {
	return &listHotelBiz{repo: repo, rtStore: rtStore}
}

func (biz *listHotelBiz) ListHotelWithFilter(ctx context.Context,
	filter *hotelmodel.Filter,
	paging *common.Paging) ([]hotelmodel.Hotel, error) {
	err := filter.Validate()
	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	result, err := biz.repo.ListHotelWithFilter(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	for i := len(result) - 1; i >= 0; i-- {
		if result[i].ListAvailableRoomType == nil {
			result = append(result[:i], result[i+1:]...)
		} else {
			result[i].DisplayPrice = result[i].ListAvailableRoomType[0].Price
			result[i].Rating, result[i].TotalRating, _ = biz.rtStore.GetHotelRating(ctx, result[i].Id)
		}
	}

	if len(result) == 0 {
		if paging.Total > int64(paging.Limit*paging.Page) {
			paging.Page++
			return biz.ListHotelWithFilter(ctx, filter, paging)
		} else {
			return nil, nil
		}
	}

	return result, nil
}
