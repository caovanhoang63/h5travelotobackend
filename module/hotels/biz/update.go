package hotelbiz

import (
	"fmt"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	"h5travelotobackend/component/pubsub"
	hotelmodel "h5travelotobackend/module/hotels/model"
)

type UpdateHotelStore interface {
	Update(ctx context.Context, id int, update *hotelmodel.HotelUpdate) error
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*hotelmodel.Hotel, error)
}

type updateHotelBiz struct {
	store UpdateHotelStore
	ps    pubsub.Pubsub
}

func NewUpdateHotelBiz(store UpdateHotelStore, ps pubsub.Pubsub) *updateHotelBiz {
	return &updateHotelBiz{store: store, ps: ps}
}

func (biz *updateHotelBiz) UpdateHotel(ctx context.Context, id int, data *hotelmodel.HotelUpdate) error {
	if err := data.Validate(); err != nil {
		return err
	}

	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(hotelmodel.EntityName, err)
	}
	if oldData.Status == common.StatusDeleted {
		return common.ErrEntityDeleted(hotelmodel.EntityName, nil)
	}

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(hotelmodel.EntityName, err)
	}

	if err := biz.ps.Publish(ctx, common.TopicUpdateHotel, pubsub.NewMessage(data)); err != nil {
		fmt.Println(common.ErrCannotPublishMessage(common.TopicUpdateHotel, err))
	}

	return nil
}
