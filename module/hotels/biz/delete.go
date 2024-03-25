package hotelbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type DeleteHotelStore interface {
	Delete(ctx context.Context, id int) error
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*hotelmodel.Hotel, error)
}

type deleteHotelBiz struct {
	store DeleteHotelStore
	ps    pubsub.Pubsub
}

func NewDeleteHotelBiz(store DeleteHotelStore, ps pubsub.Pubsub) *deleteHotelBiz {
	return &deleteHotelBiz{
		store: store,
		ps:    ps,
	}
}

func (biz *deleteHotelBiz) DeleteHotel(ctx context.Context, id int) error {
	data, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(hotelmodel.EntityName, err)
	}
	if data.Status == common.StatusDeleted {
		return common.ErrEntityDeleted(hotelmodel.EntityName, nil)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(hotelmodel.EntityName, err)
	}

	return nil
}
