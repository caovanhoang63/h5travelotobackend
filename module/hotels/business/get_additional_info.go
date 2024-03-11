package hotelbiz

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type FindHotelAddInfoRepo interface {
	FindAdditionalDataById(ctx context.Context, id int) (*hotelmodel.HotelAdditionalInfo, error)
}

type findHotelAddInfoBiz struct {
	repo FindHotelAddInfoRepo
}

func NewFindAddInfoHotelBiz(repo FindHotelAddInfoRepo) *findHotelAddInfoBiz {
	return &findHotelAddInfoBiz{repo: repo}
}

func (b *findHotelAddInfoBiz) GetHotelAdditionalInfoById(ctx context.Context,
	id int) (*hotelmodel.HotelAdditionalInfo, error) {

	data, err := b.repo.FindAdditionalDataById(ctx, id)
	if err != nil {
		return nil, err
	}

	if data.Status == 0 {
		return nil, common.ErrEntityNotFound(hotelmodel.EntityName, err)
	}

	return data, nil

}
