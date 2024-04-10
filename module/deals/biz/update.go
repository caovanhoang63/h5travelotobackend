package dealbiz

import (
	"golang.org/x/net/context"
	"h5travelotobackend/common"
	dealmodel "h5travelotobackend/module/deals/model"
)

type UpdateDealStore interface {
	Update(ctx context.Context, id int, update *dealmodel.DealUpdate) error
}

type updateDealBiz struct {
	store UpdateDealStore
}

func NewUpdateDealBiz(store UpdateDealStore) *updateDealBiz {
	return &updateDealBiz{store: store}
}

func (biz *updateDealBiz) UpdateDeal(ctx context.Context, id int, update *dealmodel.DealUpdate) error {
	if err := biz.store.Update(ctx, id, update); err != nil {
		return common.ErrInternal(err)
	}
	return nil
}
