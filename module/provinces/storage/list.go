package provincestorage

import (
	"context"
	"h5travelotobackend/common"
	provincemodel "h5travelotobackend/module/provinces/model"
)

func (s *sqlStore) ListAllProvinces(ctx context.Context) ([]provincemodel.Province, error) {
	var provinces []provincemodel.Province
	if err := s.db.Find(&provinces).Error; err != nil {
		return nil, common.ErrDb(err)
	}
	return provinces, nil
}
