package hoteltypebiz

import (
	"context"
	"h5travelotobackend/common"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
)

type FindHotelTypeStore interface {
	FindById(ctx context.Context, id int) (*hoteltypemodel.HotelType, error)
}

type findHotelTypeBiz struct {
	store FindHotelTypeStore
}

func NewFindHotelTypeBiz(store FindHotelTypeStore) *findHotelTypeBiz {
	return &findHotelTypeBiz{store: store}
}

func (b *findHotelTypeBiz) FindHotelTypeById(ctx context.Context, id int) (*hoteltypemodel.HotelType, error) {
	data, err := b.store.FindById(ctx, id)
	if err != nil {
		return nil, common.ErrEntityNotFound(hoteltypemodel.EntityName, err)
	}

	if data.Id == 0 {
		return nil, common.ErrCannotDeleteEntity(hoteltypemodel.EntityName, nil)
	}

	return data, nil
}
