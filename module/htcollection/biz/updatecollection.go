package htcollectionbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htcollection "h5travelotobackend/module/htcollection/model"
)

type UpdateCollectionStore interface {
	Update(ctx context.Context, id int,
		update *htcollection.HotelCollectionUpdate,
	) error
	FindCollection(ctx context.Context,
		conditions map[string]interface{},
	) (*htcollection.HotelCollection, error)
}

type updateCollectionBiz struct {
	store UpdateCollectionStore
}

func NewUpdateCollectionBiz(store UpdateCollectionStore) *updateCollectionBiz {
	return &updateCollectionBiz{store: store}

}

func (biz *updateCollectionBiz) UpdateCollection(ctx context.Context,
	id int, data *htcollection.HotelCollectionUpdate, requester common.Requester,
) error {

	cond := map[string]interface{}{"id": id}

	collection, err := biz.store.FindCollection(ctx, cond)
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return common.NewResourceNotFound(htcollection.EntityName)
		}
	}

	if requester.GetUserId() != collection.UserId && requester.GetRole() != common.RoleAdmin {
		return common.ErrNoPermission(nil)
	}

	if collection.Status != common.StatusActive {
		return common.ErrEntityDeleted(htcollection.EntityName, nil)
	}

	if err = biz.store.Update(ctx, id, data); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
