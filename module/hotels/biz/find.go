package hotelbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type FindHotelStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*hotelmodel.Hotel, error)
}

type findHotelBiz struct {
	store FindHotelStore
	ps    pubsub.Pubsub
}

func NewFindHotelBiz(store FindHotelStore, ps pubsub.Pubsub) *findHotelBiz {
	return &findHotelBiz{store: store, ps: ps}
}

func (biz *findHotelBiz) FindWithConditionHotel(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*hotelmodel.Hotel, error) {
	data, err := biz.store.FindDataWithCondition(ctx, condition, moreKeys...)
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.RecordNotFound
		} else {
			return nil, common.ErrEntityNotFound(hotelmodel.EntityName, err)
		}
	}

	if data.Status == common.StatusDeleted {
		return nil, common.ErrEntityDeleted(hotelmodel.EntityName, nil)
	}

	return data, nil
}
