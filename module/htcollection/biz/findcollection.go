package htcollectionbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htcollection "h5travelotobackend/module/htcollection/model"
)

type FindCollectionStore interface {
	FindCollection(ctx context.Context,
		conditions map[string]interface{},
	) (*htcollection.HotelCollection, error)
}

type FindSavedHotelBiz struct {
	store FindCollectionStore
}

func NewFincCollectionBiz(store FindCollectionStore) *FindSavedHotelBiz {
	return &FindSavedHotelBiz{store: store}
}

func (biz *FindSavedHotelBiz) FindCollectionById(ctx context.Context,
	id int, requester common.Requester) (
	*htcollection.HotelCollection,
	error) {

	cond := map[string]interface{}{"id": id}

	result, err := biz.store.FindCollection(ctx, cond)
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return nil, common.NewResourceNotFound(htcollection.EntityName)
		}
		return nil, common.ErrInternal(err)
	}

	if result.Status != common.StatusActive {
		return nil, common.ErrEntityDeleted(htcollection.EntityName, nil)
	}

	if requester.GetUserId() != result.UserId && !result.IsPrivate {
		return nil, common.ErrNoPermission(nil)
	}

	return result, nil
}
