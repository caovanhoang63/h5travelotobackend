package hotelrepo

import (
	"context"
	"gorm.io/gorm"
	"h5travelotobackend/common"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type FindHotelStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*hotelmodel.Hotel, error)
}

type findHotelRepo struct {
	store FindHotelStore
}

func NewFindHotelRepo(store FindHotelStore) *findHotelRepo {
	return &findHotelRepo{store: store}
}

func (r *findHotelRepo) FindHotelBaseInfoWithCondition(ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*hotelmodel.Hotel, error) {
	result, err := r.store.FindDataWithCondition(ctx, condition, moreKeys...)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrEntityNotFound(hotelmodel.EntityName, err)
	}

	return result, nil
}
