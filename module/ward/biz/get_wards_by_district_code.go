package wardbiz

import (
	"context"
	"h5travelotobackend/common"
	wardmodel "h5travelotobackend/module/ward/model"
)

type ListWardStore interface {
	ListWardWithCondition(ctx context.Context, filter *wardmodel.Filter) ([]wardmodel.Ward, error)
}

type listWardBiz struct {
	store ListWardStore
}

func NewListWardBiz(store ListWardStore) *listWardBiz {
	return &listWardBiz{store: store}

}

func (biz *listWardBiz) ListWardsByDistrictCode(ctx context.Context, districtCode int) ([]wardmodel.Ward, error) {
	filter := wardmodel.Filter{
		DistrictCode: districtCode,
	}
	wards, err := biz.store.ListWardWithCondition(ctx, &filter)
	if err != nil {
		return nil, common.ErrCannotListEntity(wardmodel.EntityName, err)
	}

	return wards, nil
}
