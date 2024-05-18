package htcollectionbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htcollection "h5travelotobackend/module/htcollection/model"
)

type AddRemoveStore interface {
	RemoveHotelFromCollection(ctx context.Context, hotelId, collectionId int) error
	FindCollection(ctx context.Context,
		conditions map[string]interface{},
	) (*htcollection.HotelCollection, error)
}

type removeHotelBiz struct {
	store AddRemoveStore
}

func NewRemoveHotelBiz(store AddRemoveStore) *removeHotelBiz {
	return &removeHotelBiz{store: store}
}

func (biz *removeHotelBiz) RemoveHotelFromCollection(ctx context.Context,
	hotelId int, collectionId int,
	requester common.Requester) error {
	cond := map[string]interface{}{"id": collectionId}

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

	if err = biz.store.RemoveHotelFromCollection(ctx, hotelId, collectionId); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
