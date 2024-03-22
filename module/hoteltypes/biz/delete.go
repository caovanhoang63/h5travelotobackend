package hoteltypebiz

import (
	"context"
	"h5travelotobackend/common"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
)

type DeleteHotelTypeStore interface {
	Delete(ctx context.Context, id int) error
}

type deleteHotelTypeBiz struct {
	store DeleteHotelTypeStore
}

func NewDeleteHotelTypeBiz(store DeleteHotelTypeStore) *deleteHotelTypeBiz {
	return &deleteHotelTypeBiz{store: store}
}

func (biz *deleteHotelTypeBiz) DeleteHotelType(ctx context.Context, id int) error {
	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotCreateEntity(hoteltypemodel.EntityName, err)
	}

	return nil
}
