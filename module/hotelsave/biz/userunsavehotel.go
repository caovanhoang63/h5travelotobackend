package htsavebiz

import (
	"golang.org/x/net/context"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
)

type UnsaveHotelStore interface {
	Delete(ctx context.Context, del *htsavemodel.HotelSaveDelete) error
}

type UnsaveHotelBiz struct {
	store UnsaveHotelStore
}

func NewUnsaveHotelBiz(store UnsaveHotelStore) *UnsaveHotelBiz {
	return &UnsaveHotelBiz{store: store}
}

func (biz *UnsaveHotelBiz) UnsaveHotel(ctx context.Context, del *htsavemodel.HotelSaveDelete) error {
	if err := biz.store.Delete(ctx, del); err != nil {
		return err
	}
	return nil
}
