package dealbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	dealmodel "h5travelotobackend/module/deals/model"
)

type DeleteDealStore interface {
	DeleteDeal(ctx context.Context, id int) error
	FindWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*dealmodel.Deal, error)
}

type deleteDealBiz struct {
	store DeleteDealStore
}

func NewDeleteDealBiz(store DeleteDealStore) *deleteDealBiz {
	return &deleteDealBiz{store: store}
}

func (biz *deleteDealBiz) DeleteDeal(ctx context.Context, id int) error {
	oldData, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err == common.RecordNotFound {
			return common.ErrEntityNotFound(dealmodel.EntityName, err)
		}
		return common.ErrInternal(err)
	}

	if oldData.Status == common.StatusDeleted {
		return common.ErrEntityDeleted(dealmodel.EntityName, nil)
	}

	if err := biz.store.DeleteDeal(ctx, id); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
