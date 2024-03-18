package wardstorage

import (
	"context"
	"h5travelotobackend/common"
	wardmodel "h5travelotobackend/module/ward/model"
)

func (s *sqlStore) ListWardWithCondition(ctx context.Context, filter *wardmodel.Filter) ([]wardmodel.Ward, error) {
	var wards []wardmodel.Ward
	db := s.db
	db = db.Table(wardmodel.Ward{}.TableName())

	if filter != nil {
		if filter.DistrictCode > 0 {
			db.Where("district_code = ? ", filter.DistrictCode)
		}
	}

	if err := db.Find(&wards).Error; err != nil {
		return nil, common.ErrDb(err)
	}

	return wards, nil

}
