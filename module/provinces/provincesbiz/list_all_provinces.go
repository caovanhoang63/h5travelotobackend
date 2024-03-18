package provincesbiz

import (
	"context"
	"h5travelotobackend/common"
	provincemodel "h5travelotobackend/module/provinces/model"
)

type ListAllProvincesStore interface {
	ListAllProvinces(ctx context.Context) ([]provincemodel.Province, error)
}

type listAllProvincesBiz struct {
	store ListAllProvincesStore
}

func NewListAllProvincesBiz(store ListAllProvincesStore) *listAllProvincesBiz {
	return &listAllProvincesBiz{store: store}
}

func (biz *listAllProvincesBiz) ListAllProvinces(ctx context.Context) ([]provincemodel.Province, error) {
	provinces, err := biz.store.ListAllProvinces(ctx)
	if err != nil {
		return nil, common.ErrCannotListEntity(provincemodel.EntityName, err)
	}

	return provinces, nil
}
