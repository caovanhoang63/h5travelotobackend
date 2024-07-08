package hoteltypebiz

import (
	"context"
	"h5travelotobackend/common"
	hoteltypemodel "h5travelotobackend/module/hoteltypes/model"
)

type ListHotelTypeStore interface {
	ListAllHotelTypes(ctx context.Context) ([]hoteltypemodel.HotelType, error)
}

type listHotelTypeBiz struct {
	store ListHotelTypeStore
}

func NewListHotelTypeBiz(store ListHotelTypeStore) *listHotelTypeBiz {
	return &listHotelTypeBiz{store: store}
}

func (b *listHotelTypeBiz) ListAllHotelTypes(ctx context.Context,
) ([]hoteltypemodel.HotelType, error) {
	data, err := b.store.ListAllHotelTypes(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(hoteltypemodel.EntityName, err)
	}

	return data, nil
}
