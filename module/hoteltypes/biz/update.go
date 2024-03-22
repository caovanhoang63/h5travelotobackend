package hoteltypebiz

import (
	"context"
	"h5travelotobackend/common"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
)

type UpdateHotelTypeStore interface {
	Update(ctx context.Context, id int, update *hoteltypemodel.HotelTypeUpdate) error
}

type updateHotelTypeBiz struct {
	store UpdateHotelTypeStore
}

func NewUpdateHotelTypeBiz(store UpdateHotelTypeStore) *updateHotelTypeBiz {
	return &updateHotelTypeBiz{store: store}
}

func (biz *updateHotelTypeBiz) UpdateHotelType(ctx context.Context, id int, data *hoteltypemodel.HotelTypeUpdate) error {

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(hoteltypemodel.EntityName, err)
	}

	return nil
}
