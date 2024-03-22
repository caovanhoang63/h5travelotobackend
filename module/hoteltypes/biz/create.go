package hoteltypebiz

import (
	"context"
	"h5travelotobackend/common"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
)

type CreateHotelTypeStore interface {
	Create(ctx context.Context, data *hoteltypemodel.HotelTypeCreate) error
}

type createHotelTypeBiz struct {
	store CreateHotelTypeStore
}

func NewCreateHotelTypeBiz(store CreateHotelTypeStore) *createHotelTypeBiz {
	return &createHotelTypeBiz{store: store}
}

func (biz *createHotelTypeBiz) CreateHotelType(ctx context.Context, data *hoteltypemodel.HotelTypeCreate) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(hoteltypemodel.EntityName, err)
	}

	return nil
}
