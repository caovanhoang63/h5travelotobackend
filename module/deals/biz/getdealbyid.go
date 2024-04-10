package dealbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	dealmodel "h5travelotobackend/module/deals/model"
)

type FindDealStore interface {
	FindWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*dealmodel.Deal, error)
}

type findDealBiz struct {
	store FindDealStore
}

func NewFindDealBiz(store FindDealStore) *findDealBiz {
	return &findDealBiz{store: store}
}

func (biz *findDealBiz) FindDealById(ctx context.Context, id int) (*dealmodel.Deal, error) {
	deal, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return nil, common.ErrEntityNotFound(dealmodel.EntityName, err)
		}
		return nil, common.ErrInternal(err)
	}
	return deal, nil
}
