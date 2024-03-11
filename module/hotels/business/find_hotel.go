package hotelbiz

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type FindHotelRepo interface {
	FindBaseDataWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*hotelmodel.Hotel, error)
	FindAllDataWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*hotelmodel.Hotel, error)
}

type findHotelBiz struct {
	repo FindHotelRepo
}

func NewFindHotelBiz(repo FindHotelRepo) *findHotelBiz {
	return &findHotelBiz{repo: repo}
}

func (b *findHotelBiz) GetHotelById(ctx context.Context,
	id int, incAdd bool) (*hotelmodel.Hotel, error) {
	var data *hotelmodel.Hotel
	var err error

	if incAdd {
		data, err = b.repo.FindAllDataWithCondition(ctx, map[string]interface{}{"id": id})

	} else {
		data, err = b.repo.FindBaseDataWithCondition(ctx, map[string]interface{}{"id": id})
	}
	if err != nil {
		return nil, err
	}

	if data.Status == 0 {
		return nil, common.ErrEntityNotFound(hotelmodel.EntityName, err)
	}

	return data, nil

}
