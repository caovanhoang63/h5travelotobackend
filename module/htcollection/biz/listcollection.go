package htcollectionbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	htcollection "h5travelotobackend/module/htcollection/model"
)

type ListCollectionStore interface {
	ListCollectionWithCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *htcollection.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]htcollection.HotelCollection, error)
}

type listCollectionBiz struct {
	store ListCollectionStore
}

func NewListCollectionBiz(store ListCollectionStore) *listCollectionBiz {
	return &listCollectionBiz{
		store: store,
	}
}

func (biz *listCollectionBiz) ListCollectionsWithCondition(ctx context.Context,
	filter *htcollection.Filter,
	paging *common.Paging,
	requester common.Requester) ([]htcollection.HotelCollection, error) {

	cond := map[string]interface{}{"user_id": requester.GetUserId()}

	if filter.UserId != requester.GetUserId() && requester.GetRole() != common.RoleAdmin {
		cond["is_private"] = false
	}
	result, err := biz.store.ListCollectionWithCondition(ctx, cond,
		filter, paging)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return result, nil
}
