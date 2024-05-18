package htcollectionbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htcollection "h5travelotobackend/module/htcollection/model"
)

type DeleteCollectionStore interface {
	Delete(ctx context.Context, id int) error
	FindCollection(ctx context.Context,
		conditions map[string]interface{},
	) (*htcollection.HotelCollection, error)
}

type deleteCollectionBiz struct {
	store DeleteCollectionStore
}

func NewDeleteCollectionBiz(store DeleteCollectionStore) *deleteCollectionBiz {
	return &deleteCollectionBiz{store: store}
}

func (biz *deleteCollectionBiz) DeleteCollection(
	ctx context.Context,
	id int,
	requester common.Requester,
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

	if err = biz.store.Delete(ctx, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
