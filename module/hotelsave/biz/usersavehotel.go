package htsavebiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
)

type SaveHotelStore interface {
	Create(ctx context.Context, data *htsavemodel.HotelSaveCreate) error
}

type SaveHotelBiz struct {
	store SaveHotelStore
}

func NewSaveHotelBiz(store SaveHotelStore) *SaveHotelBiz {
	return &SaveHotelBiz{store: store}
}

func (biz *SaveHotelBiz) SaveHotel(ctx context.Context, data *htsavemodel.HotelSaveCreate) error {
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
