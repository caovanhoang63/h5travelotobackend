package htcollectionbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htcollection "h5travelotobackend/module/htcollection/model"
)

type AddHotelStore interface {
	AddHotelToCollection(ctx context.Context, data *htcollection.HotelCollectionDetailCreate) error
	FindCollection(ctx context.Context,
		conditions map[string]interface{},
	) (*htcollection.HotelCollection, error)
}

type addHotelBiz struct {
	store AddHotelStore
}

func NewAddHotelBiz(store AddHotelStore) *addHotelBiz {
	return &addHotelBiz{store: store}
}

func (biz *addHotelBiz) AddHotelToCollection(ctx context.Context,
	data *htcollection.HotelCollectionDetailCreate,
	requester common.Requester) error {
	cond := map[string]interface{}{"id": data.CollectionId}

	collection, err := biz.store.FindCollection(ctx, cond)
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.NewResourceNotFound(htcollection.EntityName)
		}
	}

	if requester.GetUserId() != collection.UserId {
		return common.ErrNoPermission(nil)
	}

	if collection.Status != common.StatusActive {
		return common.ErrEntityDeleted(htcollection.EntityName, nil)
	}

	if err = biz.store.AddHotelToCollection(ctx, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
