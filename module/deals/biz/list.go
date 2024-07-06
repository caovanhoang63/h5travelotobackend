package dealbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	dealmodel "h5travelotobackend/module/deals/model"
)

type ListDealStore interface {
	ListWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		filter *dealmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]dealmodel.Deal, error)
}

type listDealBiz struct {
	store ListDealStore
}

func NewListDealBiz(store ListDealStore) *listDealBiz {
	return &listDealBiz{store: store}
}

func (biz *listDealBiz) ListDeal(ctx context.Context, filter *dealmodel.Filter, paging *common.Paging) ([]dealmodel.Deal, error) {
	deals, err := biz.store.ListWithCondition(ctx, nil, filter, paging, "RoomType")
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return deals, nil
}
