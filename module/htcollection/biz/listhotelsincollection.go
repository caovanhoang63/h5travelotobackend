package htcollectionbiz

import (
	"errors"
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htsavemodel "h5travelotobackend/module/hotelsave/model"
	htcollection "h5travelotobackend/module/htcollection/model"
)

type ListHotelInCollectionStore interface {
	ListHotelsInCollection(ctx context.Context,
		conditions map[string]interface{},
		filter *htsavemodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]common.Hotel, error)
	FindCollection(ctx context.Context,
		conditions map[string]interface{},
	) (*htcollection.HotelCollection, error)
}

type listHotelInCollectionBiz struct {
	store ListHotelInCollectionStore
}

func NewListHotelsInCollectionBiz(store ListHotelInCollectionStore) *listHotelInCollectionBiz {
	return &listHotelInCollectionBiz{
		store: store,
	}
}

func (biz *listHotelInCollectionBiz) ListHotelsInCollection(ctx context.Context,
	filter *htsavemodel.Filter,
	id int,
	paging *common.Paging,
	requester common.Requester) ([]common.Hotel, error) {
	cond := map[string]interface{}{"collection_id": id}

	collection, err := biz.store.FindCollection(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if errors.Is(err, common.RecordNotFound) {
			return nil, common.NewResourceNotFound(htcollection.EntityName)
		}
	}

	if collection.IsPrivate && requester.GetUserId() != collection.UserId && requester.GetRole() != common.RoleAdmin {
		return nil, common.ErrNoPermission(nil)
	}

	if collection.Status != common.StatusActive {
		return nil, common.ErrEntityDeleted(htcollection.EntityName, nil)
	}

	result, err := biz.store.ListHotelsInCollection(ctx, cond,
		nil, paging,
		"Hotel", "Hotel.Province",
		"Hotel.District", "Hotel.Ward")

	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return result, nil
}
