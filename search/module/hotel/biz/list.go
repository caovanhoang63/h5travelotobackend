package hotelsearchbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
	"log"
)

type ListHotelRepo interface {
	ListHotelWithFilter(ctx context.Context,
		filter *hotelmodel.Filter,
		paging *common.Paging) ([]hotelmodel.Hotel, error)
}

type listHotelBiz struct {
	repo ListHotelRepo
}

func NewListHotelBiz(repo ListHotelRepo) *listHotelBiz {
	return &listHotelBiz{repo: repo}
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

	log.Println("result: ", len(result))

	for i := len(result) - 1; i >= 0; i-- {
		log.Println("result[i]: ", result[i].Name)
		if result[i].ListAvailableRoomType == nil {
			result = append(result[:i], result[i+1:]...)
		}
	}

	if len(result) == 0 {
		if paging.Total > int64(paging.Limit*paging.Page) {
			log.Println("paging:", paging.Page)
			log.Println("paging TOTAL:", paging.Total)
			paging.Page++
			return biz.ListHotelWithFilter(ctx, filter, paging)
		} else {
			return nil, nil
		}
	}

	return result, nil
}
