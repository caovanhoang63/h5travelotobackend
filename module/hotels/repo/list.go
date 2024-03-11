package hotelrepo

import (
	"context"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type ListRestaurantStore interface {
	ListHotelWithCondition(
		ctx context.Context,
		filter *hotelmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]hotelmodel.Hotel, error)
}

type listRestaurantRepo struct {
	store ListRestaurantStore
	//likeStore LikeRestaurantStore
}

func NewListRestaurantRepo(store ListRestaurantStore) *listRestaurantRepo {
	return &listRestaurantRepo{store}
}

func (repo *listRestaurantRepo) List(
	ctx context.Context,
	filter *hotelmodel.Filter,
	paging *common.Paging,
) ([]hotelmodel.Hotel, error) {
	restaurants, err := repo.store.ListHotelWithCondition(ctx, filter, paging, "User")
	if err != nil {
		return nil, common.ErrCannotListEntity(hotelmodel.EntityName, err)
	}

	return restaurants, nil
}
