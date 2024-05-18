package htcollectionbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htcollection "h5travelotobackend/module/htcollection/model"
)

type CreateCollectionStore interface {
	Create(ctx context.Context, data *htcollection.HotelCollectionCreate) error
}

type createCollectionBiz struct {
	store CreateCollectionStore
}

func NewCreateCollectionBiz(store CreateCollectionStore) *createCollectionBiz {
	return &createCollectionBiz{store: store}

}

func (biz *createCollectionBiz) CreateCollection(ctx context.Context,
	data *htcollection.HotelCollectionCreate,
) error {
	if data.IsPrivate == nil {
		data.IsPrivate = new(bool)
		*data.IsPrivate = true
	}
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil

}
