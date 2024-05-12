package hotelsearchrepo

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/asyncjob"
	hotelmodel "h5travelotobackend/search/module/hotel/model"
	rtsearchmodel "h5travelotobackend/search/module/roomtype/model"
	"log"
)

type ListHotelStore interface {
	ListHotel(ctx context.Context, filter *hotelmodel.Filter, paging *common.Paging) ([]hotelmodel.Hotel, error)
}

type ListRoomTypeHandler interface {
	ListAvailableRt(ctx context.Context,
		filter *rtsearchmodel.Filter,
	) ([]rtsearchmodel.RoomType, error)
}

type listHotelRepo struct {
	store     ListHotelStore
	rTHandler ListRoomTypeHandler
}

func NewListHotelRepo(store ListHotelStore, rtHandler ListRoomTypeHandler) *listHotelRepo {
	return &listHotelRepo{store: store, rTHandler: rtHandler}
}

func (repo *listHotelRepo) ListHotelWithFilter(ctx context.Context,
	filter *hotelmodel.Filter,
	paging *common.Paging) ([]hotelmodel.Hotel, error) {

	result, err := repo.store.ListHotel(ctx, filter, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	var jobs []asyncjob.Job
	for i := range result {
		job := asyncjob.NewQueryJob(func(ctx context.Context) error {
			result[i].ListAvailableRoomType, err = repo.rTHandler.ListAvailableRt(ctx, &rtsearchmodel.Filter{
				QueryTime:    filter.QueryTime,
				HotelId:      result[i].Id,
				Customer:     filter.Customer,
				MinPrice:     filter.MinPrice,
				MaxPrice:     filter.MaxPrice,
				RoomQuantity: filter.RoomQuantity,
				StartDate:    filter.StartDate,
				EndDate:      filter.EndDate,
				Adult:        filter.Adults,
				Child:        filter.Children,
			})
			if err != nil {
				return common.ErrInternal(err)
			}
			return nil
		})

		jobs = append(jobs, job)
	}

	if err = asyncjob.NewGroup(true, jobs...).Run(ctx); err != nil {
		log.Println(err)
	}

	return result, nil
}
