package hotelbiz

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type ListRestaurantRepo interface {
	List(
		ctx context.Context,
		filter *hotelmodel.Filter,
		paging *common.Paging,
	) ([]hotelmodel.Hotel, error)
}

type listRestaurantBiz struct {
	repo ListRestaurantRepo
}

func NewListRestaurantBiz(repo ListRestaurantRepo) *listRestaurantBiz {
	return &listRestaurantBiz{repo: repo}
}

func (biz *listRestaurantBiz) List(
	ctx context.Context,
	filter *hotelmodel.Filter,
	paging *common.Paging,
) ([]hotelmodel.Hotel, error) {
	result, err := biz.repo.List(ctx, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(hotelmodel.EntityName, err)
	}

	return result, nil
}
