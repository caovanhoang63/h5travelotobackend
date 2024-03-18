package districtbiz

import (
	"context"
	"h5travelotobackend/common"
	districtmodel "h5travelotobackend/module/districts/model"
)

type ListDistrictStore interface {
	ListDistrictWithCondition(ctx context.Context, filter *districtmodel.Filter) ([]districtmodel.District, error)
}

type listDistrictBiz struct {
	store ListDistrictStore
}

func NewListDistrictBiz(store ListDistrictStore) *listDistrictBiz {
	return &listDistrictBiz{store: store}

}

func (biz *listDistrictBiz) ListDistrictByProvinceCode(ctx context.Context, provinceCode int) ([]districtmodel.District, error) {
	filter := districtmodel.Filter{
		ProvinceCode: provinceCode,
	}
	districts, err := biz.store.ListDistrictWithCondition(ctx, &filter)
	if err != nil {
		return nil, common.ErrCannotListEntity(districtmodel.EntityName, err)
	}

	return districts, nil
}
