package dealbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	dealmodel "h5travelotobackend/module/deals/model"
)

type CreateDealStore interface {
	Create(ctx context.Context, deal *dealmodel.DealCreate) error
}

type createDealBiz struct {
	store CreateDealStore
}

func NewCreateDealBiz(store CreateDealStore) *createDealBiz {
	return &createDealBiz{store: store}
}

func (biz *createDealBiz) CreateDeal(ctx context.Context, deal *dealmodel.DealCreate) error {

	if err := biz.store.Create(ctx, deal); err != nil {
		return common.ErrCannotCreateEntity(dealmodel.EntityName, err)
	}
	return nil
}
