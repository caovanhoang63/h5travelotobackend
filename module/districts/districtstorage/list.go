package districtstorage

import (
	"context"
	"h5travelotobackend/common"
	districtmodel "h5travelotobackend/module/districts/model"
)

func (s *sqlStore) ListDistrictWithCondition(ctx context.Context, filter *districtmodel.Filter) ([]districtmodel.District, error) {
	var districts []districtmodel.District
	db := s.db
	db = db.Table(districtmodel.District{}.TableName())

	if filter != nil {
		if filter.ProvinceCode > 0 {
			db.Where("province_code = ? ", filter.ProvinceCode)
		}
	}

	if err := db.Find(&districts).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return districts, nil

}
